// Package update checks for an available update on GitHub.
// It has baked in assumptions, but is mostly portable.
// Makes it easy for the notifiarr client application to send a
// notification when a new version is available.
package update

import (
	"archive/zip"
	"bytes"
	"compress/bzip2"
	"compress/gzip"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

// SleepTime is how long we wait after updating before exiting.
const SleepTime = 5 * time.Second

const (
	downloadTimeout = 5 * time.Minute
)

// Command is the input data to perform an in-place update.
type Command struct {
	URL         string   // file to download.
	Path        string   // file to be updated.
	Args        []string // optional, but non-nil will crash.
	*log.Logger          // debug logs.
}

// Errors. Don't trigger these.
var (
	ErrInvalidURL = fmt.Errorf("invalid URL provided")
	ErrNoPath     = fmt.Errorf("a path to the file being replaced must be provided")
)

// Now downloads the new file to a temp name in the same folder as the running file.
// Moves the running file to a backup name in the same folder.
// Moves the new file to the same location that the running file was at.
// Triggers another invocation of the app that sleeps 5 seconds then restarts.
// The running app must exit after this returns!
// The restart command can trigger the above Restart() procedure.
// And that prcoedure relaunches the app; this allows "in-place" upgrades.
// This also makes sure the new file works before this app exits.
// This is not required though, and you can totally upgrade "a different app".
func Now(ctx context.Context, u *Command) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, downloadTimeout)
	defer cancel()

	return NowWithContext(ctx, u)
}

// NowWithContext is the same as Now() except you can pass in your own context.
func NowWithContext(ctx context.Context, update *Command) (string, error) {
	if update.Path == "" {
		return "", ErrNoPath
	} else if _, err := os.Stat(update.Path); err != nil {
		return "", fmt.Errorf("checking path: %w", err)
	} else if !strings.HasPrefix(update.URL, "http") {
		return "", ErrInvalidURL
	}

	backupFile, err := update.replaceFile(ctx)
	if err != nil {
		return backupFile, err
	}

	update.Printf("[Update] Triggering Restart: %s %s", update.Path, strings.Join(update.Args, " "))

	if err := exec.Command(update.Path, update.Args...).Start(); err != nil { //nolint:gosec
		return backupFile, fmt.Errorf("executing restart command %w", err)
	}

	return backupFile, nil
}

func (u *Command) replaceFile(ctx context.Context) (string, error) {
	tempFolderPath, err := filepath.Abs(filepath.Dir(u.Path))
	if err != nil {
		return "", fmt.Errorf("getting application folder: %w", err)
	}

	tempFile, err := u.writeFile(ctx, tempFolderPath)
	if err != nil {
		return "", err
	}

	if runtime.GOOS != "windows" {
		_ = os.Chmod(tempFile, 0o755) //nolint:gomnd
	}

	suff := ""
	if strings.HasSuffix(u.Path, ".exe") {
		suff = ".exe"
	}

	backupFile := strings.TrimSuffix(u.Path, ".exe")
	backupFile += ".backup." + time.Now().Format("060102T150405") + suff
	u.Printf("[Update] Renaming %s => %s", u.Path, backupFile)

	if err := os.Rename(u.Path, backupFile); err != nil {
		return backupFile, fmt.Errorf("renaming original file: %w", err)
	}

	u.Printf("[Update] Renaming %s => %s", tempFile, u.Path)

	if err := os.Rename(tempFile, u.Path); err != nil {
		return backupFile, fmt.Errorf("renaming downloaded file: %w", err)
	}
	/* // Hack used for testing.
	u.Printf("[Update] Renaming [HACK] %s => %s", backupFile, u.Path)

	if err := os.Rename(backupFile, u.Path); err != nil {
		return backupFile, fmt.Errorf("renaming downloaded file %w", err)
	}
	/**/
	return backupFile, nil
}

func DownloadURL(ctx context.Context, uri string, fileName string, logger *log.Logger) (string, error) {
	if logger == nil {
		logger = log.New(io.Discard, "", 0)
	}

	command := &Command{
		URL:    uri,
		Path:   fileName,
		Logger: logger,
	}

	tempFile, err := command.writeFile(ctx, os.TempDir())
	if err != nil {
		return "", err
	}

	newPath := filepath.Join(os.TempDir(), fileName)

	switch user, err := user.Current(); {
	default:
	case err != nil:
	case runtime.GOOS == "darwin":
		fallthrough
	case runtime.GOOS == "windows":
		newPath = filepath.Join(user.HomeDir, "Downloads", fileName)
	}

	if _, err := os.Stat(newPath); err == nil {
		// file already exists.
		ext := filepath.Ext(fileName)
		base := strings.TrimSuffix(fileName, ext) + "."
		// find a new name.
		for i := 1; i <= 99; i++ {
			tryPath := filepath.Join(filepath.Dir(newPath), fmt.Sprint(base, i, ext))
			if _, err := os.Stat(tryPath); err != nil {
				newPath = tryPath
				break
			}

		}
	}
	if err = os.Rename(tempFile, newPath); err != nil {
		return "", fmt.Errorf("renaming temp file: %w", err)
	}

	return newPath, nil
}

func (u *Command) writeFile(ctx context.Context, folderPath string) (string, error) {
	tempFile, err := os.CreateTemp(folderPath, filepath.Base(u.Path))
	if err != nil {
		return "", fmt.Errorf("creating temporary file: %w", err)
	}
	defer tempFile.Close()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.URL, nil)
	if err != nil {
		return tempFile.Name(), fmt.Errorf("creating request: %w", err)
	}

	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return tempFile.Name(), fmt.Errorf("downloading release: %w", err)
	}
	defer resp.Body.Close()

	return tempFile.Name(), u.decompressFile(tempFile, resp.Body, resp.ContentLength)
}

func (u *Command) decompressFile(tempFile *os.File, resp io.Reader, size int64) error {
	switch {
	case strings.HasSuffix(u.URL, ".zip"):
		if body, err := io.ReadAll(resp); err != nil {
			return fmt.Errorf("reading file from URL: %w", err)
		} else if err := u.writeZipFile(tempFile, body, size); err != nil {
			return err
		}
	case strings.HasSuffix(u.URL, ".gz"):
		if err := u.writeGZipFile(tempFile, resp); err != nil {
			return err
		}
	case strings.HasSuffix(u.URL, ".bz2"):
		if _, err := io.Copy(tempFile, bzip2.NewReader(resp)); err != nil {
			return fmt.Errorf("bzunzipping temporary file: %w", err)
		}
	default:
		if _, err := io.Copy(tempFile, resp); err != nil {
			return fmt.Errorf("writing temporary file: %w", err)
		}
	}

	return nil
}

func (u *Command) writeGZipFile(tempFile *os.File, resp io.Reader) error {
	gzFile, err := gzip.NewReader(resp)
	if err != nil {
		return fmt.Errorf("reading gzip file: %w", err)
	}
	defer gzFile.Close()

	if _, err := io.Copy(tempFile, gzFile); err != nil { //nolint:gosec
		return fmt.Errorf("gunzipping temporary file: %w", err)
	}

	return nil
}

func (u *Command) writeZipFile(tempFile *os.File, body []byte, size int64) error {
	if size < 1 {
		size = int64(len(body))
	}

	// Open a zip archive for reading.
	zipReader, err := zip.NewReader(bytes.NewReader(body), size)
	if err != nil {
		return fmt.Errorf("unzipping downloaded file: %w", err)
	}

	// Find the exe file and write that.
	for _, zipFile := range zipReader.File {
		if runtime.GOOS == "windows" && strings.HasSuffix(zipFile.Name, ".exe") {
			zipOpen, err := zipFile.Open()
			if err != nil {
				return fmt.Errorf("reading zipped exe file: %w", err)
			}
			defer zipOpen.Close()

			if _, err := io.Copy(tempFile, zipOpen); err != nil { //nolint:gosec
				return fmt.Errorf("unzipping temporary file: %w", err)
			}

			return nil
		}
	}

	u.Println("[Update] exe file not found in zip file")

	return nil
}
