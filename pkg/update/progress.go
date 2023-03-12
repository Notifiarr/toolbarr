package update

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/user"
	"path/filepath"
	"strings"
	"sync"

	"github.com/Notifiarr/toolbarr/pkg/mnd"
	wailsRuntime "github.com/wailsapp/wails/v2/pkg/runtime"
	"golift.io/datacounter"
)

type Progress struct {
	ctx     context.Context
	cmd     *Command
	counter *datacounter.ReaderCounter
	close   func() error
	error   error
	mu      sync.Mutex
	size    int64
	done    bool
}

// Progress returns the size of the downloaded data, and the total size expected.
// Provided as floats so you may create a percentage.
// Also returns true/false if the download is finished, and any error reported.
func (p *Progress) Progress() (float64, float64, bool, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	return float64(p.counter.Count()), float64(p.size), p.done, p.error
}

func (p *Progress) Read(buf []byte) (int, error) {
	defer wailsRuntime.EventsEmit(p.ctx, "downloadProgress", float64(p.counter.Count())/float64(p.size))
	return p.counter.Read(buf) //nolint:wrapcheck
}

func (p *Progress) Close() error {
	defer wailsRuntime.EventsEmit(p.ctx, "downloadFinished")
	return p.close()
}

func (p *Progress) Path() string {
	return p.cmd.Path
}

func (p *Progress) Size() int64 {
	return p.size
}

func DownloadURL(ctx context.Context, uri string, fileName string, logger *log.Logger) (*Progress, error) {
	if logger == nil {
		logger = log.New(io.Discard, "", 0)
	}

	command := &Command{
		URL:    uri,
		Path:   filepath.Join(os.TempDir(), fileName),
		Logger: logger,
	}

	switch user, err := user.Current(); {
	case err != nil:
	case mnd.IsMac:
		fallthrough
	case mnd.IsWindows:
		command.Path = filepath.Join(user.HomeDir, "Downloads", fileName)
	}

	if _, err := os.Stat(command.Path); err == nil {
		// file already exists.
		ext := filepath.Ext(fileName)
		base := strings.TrimSuffix(fileName, ext) + "."
		// find a new name.
		for i := 1; i <= 99; i++ {
			tryPath := filepath.Join(filepath.Dir(command.Path), fmt.Sprint(base, i, ext))
			if _, err := os.Stat(tryPath); err != nil {
				command.Path = tryPath
				break
			}
		}
	}

	progress, err := command.downloadFile(ctx)
	if err != nil {
		return nil, err
	}

	return progress, nil
}

func (u *Command) downloadFile(ctx context.Context) (*Progress, error) {
	openFile, err := os.Create(u.Path)
	if err != nil {
		return nil, fmt.Errorf("creating temporary file: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.URL, nil)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	resp, err := (&http.Client{}).Do(req) //nolint:bodyclose // closed later.
	if err != nil {
		return nil, fmt.Errorf("downloading release: %w", err)
	}

	progress := &Progress{
		ctx:     ctx,
		cmd:     u,
		close:   resp.Body.Close,
		counter: datacounter.NewReaderCounter(resp.Body),
		size:    resp.ContentLength,
	}
	go progress.download(openFile)

	return progress, nil
}

func (p *Progress) download(openFile io.WriteCloser) {
	defer func() {
		openFile.Close() // Destination
		p.Close()        // Source
	}()

	err := p.cmd.decompressFile(openFile, p, p.size)

	p.mu.Lock()
	defer p.mu.Unlock()

	p.done = true

	if err != nil {
		p.error = err
	}
}
