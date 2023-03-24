package starrs

import (
	"bufio"
	"context"
	"io"
	"net/http"
	"strings"

	"github.com/Notifiarr/toolbarr/pkg/logs"
	"golift.io/starr"
)

func starrConfig(logger *logs.Logger, instance *Instance) *StarrConfig {
	return &StarrConfig{
		Logger:   logger,
		instance: instance,
		Config: &starr.Config{
			APIKey: instance.Key,
			URL:    strings.TrimSuffix(instance.URL, "/") + "/",
			// HTTPPass: instance.Pass,
			// HTTPUser: instance.User,
			Username: instance.User,
			Password: instance.Pass,
			Client:   starr.Client(timeout, instance.SSL),
		},
	}
}

// window.Readarr = {
// 	apiRoot: '/readarr/api/v1',
// 	apiKey: 'c328973936884f8fa31c840d7bee52ca',
// 	release: '0.1.4.1596-develop',
// 	version: '0.1.4.1596',
// 	instanceName: 'Readarr',
// 	theme: 'auto',
// 	branch: 'nightly',
// 	analytics: false,
// 	userHash: '8caf6f13',
// 	urlBase: '/readarr',
// 	isProduction: true
//   };

// InitializeJS is the data we care about in the initialize.js file.
type InitializeJS struct {
	App     string
	APIKey  string
	Release string
	Name    string
}

func (s *StarrConfig) getInitializeJS(ctx context.Context) (*InstanceTest, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, s.Config.URL+"initialize.js", nil)
	if err != nil {
		return nil, err //nolint:wrapcheck
	}

	resp, err := s.Client.Do(req)
	if err != nil {
		return nil, err //nolint:wrapcheck
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		_, _ = io.Copy(io.Discard, resp.Body)
		return nil, &starr.ReqError{Code: resp.StatusCode}
	}

	return readInitializeJS(resp.Body)
}

func readInitializeJS(input io.Reader) (*InstanceTest, error) {
	output := &InstanceTest{}
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		switch split := strings.Fields(scanner.Text()); {
		case len(split) < 2: //nolint:gomnd
			continue
		case split[0] == "apiKey:":
			output.Key = strings.Trim(split[1], `"',`)
		case split[0] == "version:":
			output.Version = strings.Trim(split[1], `"',`)
		case split[0] == "instanceName:":
			output.Name = strings.Trim(split[1], `"',`)
		case strings.HasPrefix(split[0], "window."):
			output.App = strings.TrimPrefix(split[0], "window.")
		}
	}

	return output, scanner.Err() //nolint:wrapcheck
}
