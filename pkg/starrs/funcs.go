package starrs

import (
	"bufio"
	"io"
	"net/http"
	"strings"

	"golift.io/starr"
)

func (s *Starrs) newInstance(config *AppConfig) *instance {
	return &instance{
		Starrs: s,
		config: config,
		Config: &starr.Config{
			APIKey: config.Key,
			URL:    strings.TrimSuffix(config.URL, "/") + "/",
			// HTTPPass: instance.Pass,
			// HTTPUser: instance.User,
			Username: config.User,
			Password: config.Pass,
			Client:   starr.Client(timeout, config.SSL),
		},
	}
}

func (s *Starrs) newAPIinstance(config *AppConfig) (*instance, error) {
	instance := s.newInstance(config)
	if instance.APIKey == "" {
		data, err := instance.testWithoutKey()
		if err != nil {
			return nil, err
		}

		instance.Config.APIKey = data.Key
	}

	return instance, nil
}

func (i *instance) getInitializeJS() (*instanceTest, error) {
	req, err := http.NewRequestWithContext(i.ctx, http.MethodGet, i.Config.URL+"initialize.js", nil)
	if err != nil {
		return nil, err
	}

	resp, err := i.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		_, _ = io.Copy(io.Discard, resp.Body)
		return nil, &starr.ReqError{Code: resp.StatusCode}
	}

	return readInitializeJS(resp.Body)
}

func readInitializeJS(input io.Reader) (*instanceTest, error) {
	output := &instanceTest{}
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

	return output, scanner.Err()
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
// type InitializeJS struct {
// 	App     string
// 	APIKey  string
// 	Release string
// 	Name    string
// }
