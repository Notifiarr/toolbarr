package update

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type UnstableFile struct {
	Time time.Time
	File string
	Ver  string
}

// LatestUS is where we find the latest unstable.
const unstableURL = "https://unstable.golift.io"

// CheckUnstable checks if the provided app has an updated version on GitHub.
// Pass in version in format version-revision, e.g. v0.1.2-346.
func CheckUnstable(ctx context.Context, app string, version string) (*Update, error) {
	uri := fmt.Sprintf("%s/%s/%s.%s.installer.zip", unstableURL, strings.ToLower(app), app, runtime.GOARCH)
	if runtime.GOOS == "linux" {
		uri = fmt.Sprintf("%s/%s/%s.%s.gz", unstableURL, strings.ToLower(app), app, runtime.GOARCH)
	} else if runtime.GOOS == "darwin" {
		uri = fmt.Sprintf("%s/%s/%s.dmg", unstableURL, strings.ToLower(app), app)
	}

	release, err := GetUnstable(ctx, uri)
	if err != nil {
		return nil, err
	}

	return fillUnstable(release, version), nil
}

// GetUnstable returns an unstable release. See CheckUnstable for an example on how to use it.
func GetUnstable(ctx context.Context, uri string) (*UnstableFile, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri+".txt", nil)
	if err != nil {
		return nil, fmt.Errorf("requesting unstable: %w", err)
	}

	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return nil, fmt.Errorf("querying unstable: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading http response: %w", err)
	}

	modTime, _ := time.Parse(time.RFC1123, resp.Header.Get("last-modified"))

	return &UnstableFile{
		Ver:  string(body),
		File: uri,
		Time: modTime,
	}, nil
}

// fillUnstable compares a current version with the latest GitHub release.
// Pass in version in format version-revision, e.g. v0.1.2-346.
func fillUnstable(release *UnstableFile, version string) *Update {
	update := &Update{
		RelDate: release.Time,
		CurrURL: release.File,
		Current: release.Ver,
		Version: "v" + strings.TrimPrefix(version, "v"),
	}

	splitN := strings.Split(release.Ver, "-") // new (from website)
	splitO := strings.Split(version, "-")     // old (passed in)

	if len(splitN) != 2 || len(splitO) != 2 {
		return update
	}

	intN, err := strconv.Atoi(splitN[1])
	if err != nil {
		return update
	}

	intO, err := strconv.Atoi(splitO[1])
	if err != nil {
		return update
	}

	if intN > intO {
		update.Outdate = true
	}

	return update
}
