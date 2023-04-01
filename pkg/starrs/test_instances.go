//nolint:wrapcheck,goerr113
package starrs

import (
	"fmt"
	"time"

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

// timeout is how long to wait for an instance to respond.
const timeout = 10 * time.Second

// InstanceTest is the data returned the front end after an instances is tested.
type InstanceTest struct {
	App     string
	Key     string
	Version string
	Name    string
	Count   int // for whatever, but table count for now.
}

func (s *Starrs) testDBPath(config *StarrConfig) (*InstanceTest, error) {
	sql, err := s.newSQL(config)
	if err != nil {
		return nil, fmt.Errorf(s.log.Translate("Connection test failed! %v", err.Error()))
	}
	defer sql.Close()

	tables, err := sql.RowsStringSlice(s.ctx, "SELECT name FROM sqlite_schema WHERE type='table'")
	if err != nil {
		return nil, fmt.Errorf(s.log.Translate("Connection test failed! Querying Sqlite3 DB: %v", err.Error()))
	}

	version, _ := sql.RowString(s.ctx, "select sqlite_version()")

	return &InstanceTest{
		Count:   len(tables),
		Version: version,
	}, nil
}

func (s *Starrs) TestInstance(config *StarrConfig) (string, error) {
	s.log.Tracef("Call:TestInstance(%s, %s)", config.App, config.Name)

	test, err := s.testInstance(config)
	if err != nil {
		return "", err
	}

	msg := s.log.Translate("<li>Connection test successful! Found %s (%s) with version %s.</li>",
		test.App, test.Name, test.Version)

	if test.App != config.App {
		msg = s.log.Translate("Connection test failed! Wrong app found. Expected %s but found %s. App Version: %s",
			config.App, test.App, test.Version)
		return "", fmt.Errorf(msg)
	}

	if config.DBPath == "" {
		return msg, nil
	}

	if test, err = s.testDBPath(config); err != nil {
		return "", err
	}

	msg += " " + s.log.Translate("<li>Database file test successful! SQLite3 version: %s, tables: %d</li>",
		test.Version, test.Count)

	return msg, nil
}

func (s *Starrs) testInstance(config *StarrConfig) (*InstanceTest, error) {
	instance := s.newInstance(config)

	if instance.APIKey == "" {
		return instance.testWithoutKey()
	}

	switch starr.App(config.App) {
	case starr.Lidarr:
		return instance.testLidarr()
	case starr.Prowlarr:
		return instance.testProwlarr()
	case starr.Radarr:
		return instance.testRadarr()
	case starr.Readarr:
		return instance.testReadarr()
	case starr.Sonarr:
		return instance.testSonarr()
	case "Whisparr":
		return instance.testWhisparr()
	default:
		return nil, fmt.Errorf("%w: missing app", starr.ErrRequestError)
	}
}

func (i *instance) testWithoutKey() (*InstanceTest, error) {
	if i.Username != "" {
		if err := i.Login(i.ctx); err != nil {
			return nil, fmt.Errorf(i.log.Translate("Login (username/password) failed: %v", err.Error()))
		}
	}

	return i.getInitializeJS()
}

func (i *instance) testLidarr() (*InstanceTest, error) {
	status, err := lidarr.New(i.Config).GetSystemStatusContext(i.ctx)
	if err != nil {
		return nil, err
	}

	return &InstanceTest{
		App:     status.AppName,
		Key:     i.APIKey,
		Version: status.Version,
		Name:    status.InstanceName,
	}, err
}

func (i *instance) testProwlarr() (*InstanceTest, error) {
	status, err := prowlarr.New(i.Config).GetSystemStatusContext(i.ctx)
	if err != nil {
		return nil, err
	}

	return &InstanceTest{
		App:     status.AppName,
		Key:     i.APIKey,
		Version: status.Version,
		Name:    status.InstanceName,
	}, err
}

func (i *instance) testRadarr() (*InstanceTest, error) {
	status, err := radarr.New(i.Config).GetSystemStatusContext(i.ctx)
	if err != nil {
		return nil, err
	}

	return &InstanceTest{
		App:     status.AppName,
		Key:     i.APIKey,
		Version: status.Version,
		Name:    status.InstanceName,
	}, err
}

func (i *instance) testReadarr() (*InstanceTest, error) {
	status, err := readarr.New(i.Config).GetSystemStatusContext(i.ctx)
	if err != nil {
		return nil, err
	}

	return &InstanceTest{
		App:     status.AppName,
		Key:     i.APIKey,
		Version: status.Version,
		Name:    status.InstanceName,
	}, err
}

func (i *instance) testSonarr() (*InstanceTest, error) {
	status, err := sonarr.New(i.Config).GetSystemStatusContext(i.ctx)
	if err != nil {
		return nil, err
	}

	return &InstanceTest{
		App:     status.AppName,
		Key:     i.APIKey,
		Version: status.Version,
		Name:    status.InstanceName,
	}, err
}

func (i *instance) testWhisparr() (*InstanceTest, error) {
	status, err := radarr.New(i.Config).GetSystemStatusContext(i.ctx)
	if err != nil {
		return nil, err
	}

	return &InstanceTest{
		App:     status.AppName,
		Key:     i.APIKey,
		Version: status.Version,
		Name:    status.InstanceName,
	}, err
}
