//nolint:tagliatelle
package update

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"
	"strings"
	"time"

	"golang.org/x/mod/semver"
)

// OSsuffixMap is the OS to file suffix map for downloads.
var OSsuffixMap = map[string]string{ //nolint:gochecknoglobals
	"darwin":  ".dmg",
	"windows": runtime.GOARCH + ".installer.exe",
	"linux":   runtime.GOARCH + ".gz",
}

var ErrNoFile = fmt.Errorf("no downloadable file found in release")

// Latest is where we find the latest release.
const LatestGH = "https://api.github.com/repos/%s/releases/latest"

// GitHub API and JSON unmarshal timeout.
const timeout = 10 * time.Second

// Update contains running Version, Current version and Download URL for Current version.
// Outdate is true if the running version is older than the current version.
type Update struct {
	Outdate bool      // True if we're outdated, update available.
	Version string    // Version passed in externally.
	Current string    // Current release available on GH or US.
	CurrURL string    // URL of current release on GH or US.
	RelDate time.Time // Current version release date.
	RelSize int64     // Current release file size.
}

// CheckGitHub checks if the provided repo has an updated version on GitHub.
func CheckGitHub(ctx context.Context, userRepo string, version string) (*Update, error) {
	release, err := GetRelease(ctx, fmt.Sprintf(LatestGH, userRepo))
	if err != nil {
		return nil, err
	}

	return FillUpdate(release, version)
}

// GetRelease returns a GitHub release. See Check for an example on how to use it.
func GetRelease(ctx context.Context, uri string) (*GitHubReleasesLatest, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, fmt.Errorf("requesting github: %w", err)
	}

	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return nil, fmt.Errorf("querying github: %w", err)
	}
	defer resp.Body.Close()

	var release GitHubReleasesLatest
	if err = json.NewDecoder(resp.Body).Decode(&release); err != nil {
		return nil, fmt.Errorf("decoding github response: %w", err)
	}

	return &release, nil
}

// FillUpdate compares a current version with the latest GitHub release.
func FillUpdate(release *GitHubReleasesLatest, version string) (*Update, error) {
	update := &Update{
		RelDate: release.PublishedAt,
		CurrURL: release.HTMLURL,
		Current: strings.TrimPrefix(release.TagName, "v"),
		Version: strings.TrimPrefix(version, "v"),
	}
	update.Outdate = semver.Compare("v"+update.Current, "v"+update.Version) < 0

	for _, file := range release.Assets {
		if strings.HasSuffix(file.BrowserDownloadURL, OSsuffixMap[runtime.GOOS]) {
			update.CurrURL = file.BrowserDownloadURL
			update.RelDate = file.UpdatedAt
			update.RelSize = file.Size

			break
		}
	}

	if release.HTMLURL == update.CurrURL {
		return update, fmt.Errorf("%w: %s", ErrNoFile, update.CurrURL)
	}

	return update, nil
}

// GitHubReleasesLatest is the output from the releases/latest API on GitHub.
type GitHubReleasesLatest struct {
	URL             string    `json:"url"`
	AssetsURL       string    `json:"assets_url"`
	UploadURL       string    `json:"upload_url"`
	HTMLURL         string    `json:"html_url"`
	ID              int64     `json:"id"`
	Author          GHuser    `json:"author"`
	NodeID          string    `json:"node_id"`
	TagName         string    `json:"tag_name"`
	TargetCommitish string    `json:"target_commitish"`
	Name            string    `json:"name"`
	Draft           bool      `json:"draft"`
	Prerelease      bool      `json:"prerelease"`
	CreatedAt       time.Time `json:"created_at"`
	PublishedAt     time.Time `json:"published_at"`
	Assets          []GHasset `json:"assets"`
	TarballURL      string    `json:"tarball_url"`
	ZipballURL      string    `json:"zipball_url"`
	Body            string    `json:"body"`
}

// GHasset is part of GitHubReleasesLatest.
type GHasset struct {
	URL                string    `json:"url"`
	ID                 int64     `json:"id"`
	NodeID             string    `json:"node_id"`
	Name               string    `json:"name"`
	Label              string    `json:"label"`
	Uploader           GHuser    `json:"uploader"`
	ContentType        string    `json:"content_type"`
	State              string    `json:"state"`
	Size               int64     `json:"size"`
	DownloadCount      int       `json:"download_count"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
	BrowserDownloadURL string    `json:"browser_download_url"`
}

// GHuser is part of GitHubReleasesLatest.
type GHuser struct {
	Login             string `json:"login"`
	ID                int64  `json:"id"`
	NodeID            string `json:"node_id"`
	AvatarURL         string `json:"avatar_url"`
	GravatarID        string `json:"gravatar_id"`
	URL               string `json:"url"`
	HTMLURL           string `json:"html_url"`
	FollowersURL      string `json:"followers_url"`
	FollowingURL      string `json:"following_url"`
	GistsURL          string `json:"gists_url"`
	StarredURL        string `json:"starred_url"`
	SubscriptionsURL  string `json:"subscriptions_url"`
	OrganizationsURL  string `json:"organizations_url"`
	ReposURL          string `json:"repos_url"`
	EventsURL         string `json:"events_url"`
	ReceivedEventsURL string `json:"received_events_url"`
	Type              string `json:"type"`
	SiteAdmin         bool   `json:"site_admin"`
}
