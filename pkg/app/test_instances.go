//nolint:wrapcheck
package app

import (
	"context"
	"crypto/tls"
	"net/http"
	"time"

	"github.com/Notifiarr/toolbarr/pkg/config"
	"golift.io/starr"
	"golift.io/starr/lidarr"
	"golift.io/starr/prowlarr"
	"golift.io/starr/radarr"
	"golift.io/starr/readarr"
	"golift.io/starr/sonarr"
)

const timeout = 10 * time.Second

func starrConfig(instance *config.Instance) *starr.Config {
	return &starr.Config{
		APIKey:   instance.Key,
		URL:      instance.URL,
		HTTPPass: instance.Pass,
		HTTPUser: instance.User,
		Username: instance.User,
		Password: instance.Pass,
		Client: &http.Client{
			Timeout: timeout,
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: !instance.SSL}, //nolint:gosec
			},
		},
	}
}

func testLidarr(ctx context.Context, config *starr.Config) (string, error) {
	status, err := lidarr.New(config).GetSystemStatusContext(ctx)
	if err != nil {
		return "", err
	}

	return "Connection Successful! Version: " + status.Version, nil
}

func testProwlarr(ctx context.Context, config *starr.Config) (string, error) {
	status, err := prowlarr.New(config).GetSystemStatusContext(ctx)
	if err != nil {
		return "", err
	}

	return "Connection Successful! Version: " + status.Version, nil
}

func testRadarr(ctx context.Context, config *starr.Config) (string, error) {
	status, err := radarr.New(config).GetSystemStatusContext(ctx)
	if err != nil {
		return "", err
	}

	return "Connection Successful! Version: " + status.Version, nil
}

func testReadarr(ctx context.Context, config *starr.Config) (string, error) {
	status, err := readarr.New(config).GetSystemStatusContext(ctx)
	if err != nil {
		return "", err
	}

	return "Connection Successful! Version: " + status.Version, nil
}

func testSonarr(ctx context.Context, config *starr.Config) (string, error) {
	status, err := sonarr.New(config).GetSystemStatusContext(ctx)
	if err != nil {
		return "", err
	}

	return "Connection Successful! Version: " + status.Version, nil
}

func testWhisparr(ctx context.Context, config *starr.Config) (string, error) {
	status, err := radarr.New(config).GetSystemStatusContext(ctx)
	if err != nil {
		return "", err
	}

	return "Connection Successful! Version: " + status.Version, nil
}
