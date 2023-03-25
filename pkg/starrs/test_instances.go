//nolint:wrapcheck,goerr113
package starrs

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/Notifiarr/toolbarr/pkg/logs"
	"golift.io/starr"
	"golift.io/starr/lidarr"
	"golift.io/starr/prowlarr"
	"golift.io/starr/radarr"
	"golift.io/starr/readarr"
	"golift.io/starr/sonarr"
	_ "modernc.org/sqlite" // database driver for sqlite3.
)

/*
 * This file is for testing starr instances.
 * Testing, as in, making a request to see if it's up and the API key works.
 */

const timeout = 10 * time.Second

type InstanceTest struct {
	App     string
	Key     string
	Version string
	Name    string
}

func TestDBPath(ctx context.Context, logger *logs.Logger, instance *Instance) (*InstanceTest, error) {
	_, err := os.Stat(instance.DBPath)
	if err != nil {
		return nil, fmt.Errorf(logger.Translate("Locating DB file failed: %v", err.Error()))
	}

	conn, err := sql.Open("sqlite", instance.DBPath)
	if err != nil {
		return nil, fmt.Errorf("opening sqlite DB: %w", err)
	}
	defer conn.Close()

	// backup := &Info{
	// 	Name:   filePath,
	// 	Size:   fileInfo.Size(),
	// 	Tables: c.getSQLLiteRowInt64(ctx, conn, "SELECT count(*) FROM sqlite_master WHERE type = 'table'"),
	// }
	// backup.Ver, _ = c.getSQLLiteRowString(ctx, conn, "select sqlite_version()")
	// backup.Integ, backup.Rows = c.getSQLLiteRowString(ctx, conn, "PRAGMA integrity_check")
	// backup.Quick, _ = c.getSQLLiteRowString(ctx, conn, "PRAGMA quick_check")

	return &InstanceTest{
		Version: "not working yet",
		App:     "Unknown",
		Name:    "Unknown",
	}, nil
}

func TestInstance(ctx context.Context, logger *logs.Logger, instance *Instance) (*InstanceTest, error) {
	starrConfig := starrConfig(logger, instance)

	if starrConfig.APIKey == "" {
		return starrConfig.testWithoutKey(ctx)
	}

	switch starr.App(instance.App) {
	case starr.Lidarr:
		return starrConfig.testLidarr(ctx)
	case starr.Prowlarr:
		return starrConfig.testProwlarr(ctx)
	case starr.Radarr:
		return starrConfig.testRadarr(ctx)
	case starr.Readarr:
		return starrConfig.testReadarr(ctx)
	case starr.Sonarr:
		return starrConfig.testSonarr(ctx)
	case "Whisparr":
		return starrConfig.testWhisparr(ctx)
	default:
		return nil, fmt.Errorf("%w: missing app", starr.ErrRequestError)
	}
}

func (s *StarrConfig) testWithoutKey(ctx context.Context) (*InstanceTest, error) {
	if s.Username != "" {
		if err := s.Login(ctx); err != nil {
			return nil, fmt.Errorf("login (username/password) failed: %w", err)
		}
	}

	return s.getInitializeJS(ctx)
}

func (s *StarrConfig) testLidarr(ctx context.Context) (*InstanceTest, error) {
	status, err := lidarr.New(s.Config).GetSystemStatusContext(ctx)
	if err != nil {
		return nil, err
	}

	return &InstanceTest{
		App:     status.AppName,
		Key:     s.APIKey,
		Version: status.Version,
		Name:    status.InstanceName,
	}, err
}

func (s *StarrConfig) testProwlarr(ctx context.Context) (*InstanceTest, error) {
	status, err := prowlarr.New(s.Config).GetSystemStatusContext(ctx)
	if err != nil {
		return nil, err
	}

	return &InstanceTest{
		App:     status.AppName,
		Key:     s.APIKey,
		Version: status.Version,
		Name:    status.InstanceName,
	}, err
}

func (s *StarrConfig) testRadarr(ctx context.Context) (*InstanceTest, error) {
	status, err := radarr.New(s.Config).GetSystemStatusContext(ctx)
	if err != nil {
		return nil, err
	}

	return &InstanceTest{
		App:     status.AppName,
		Key:     s.APIKey,
		Version: status.Version,
		Name:    status.InstanceName,
	}, err
}

func (s *StarrConfig) testReadarr(ctx context.Context) (*InstanceTest, error) {
	status, err := readarr.New(s.Config).GetSystemStatusContext(ctx)
	if err != nil {
		return nil, err
	}

	return &InstanceTest{
		App:     status.AppName,
		Key:     s.APIKey,
		Version: status.Version,
		Name:    status.InstanceName,
	}, err
}

func (s *StarrConfig) testSonarr(ctx context.Context) (*InstanceTest, error) {
	status, err := sonarr.New(s.Config).GetSystemStatusContext(ctx)
	if err != nil {
		return nil, err
	}

	return &InstanceTest{
		App:     status.AppName,
		Key:     s.APIKey,
		Version: status.Version,
		Name:    status.InstanceName,
	}, err
}

func (s *StarrConfig) testWhisparr(ctx context.Context) (*InstanceTest, error) {
	status, err := radarr.New(s.Config).GetSystemStatusContext(ctx)
	if err != nil {
		return nil, err
	}

	return &InstanceTest{
		App:     status.AppName,
		Key:     s.APIKey,
		Version: status.Version,
		Name:    status.InstanceName,
	}, err
}
