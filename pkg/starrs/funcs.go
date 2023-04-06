package starrs

import (
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
