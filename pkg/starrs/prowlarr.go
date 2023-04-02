package starrs

import "fmt"

func (s *Starrs) AppProfiles(config *AppConfig) (any, error) {
	s.log.Tracef("Call:AppProfiles(%s, %s)", config.App, config.Name)

	_, err := s.newAPIinstance(config)
	if err != nil {
		return nil, err
	}

	return nil, fmt.Errorf("prowlarr app profiles do not work yet")
}

func (s *Starrs) CustomFilters(config *AppConfig) (any, error) {
	s.log.Tracef("Call:CustomFilters(%s, %s)", config.App, config.Name)

	_, err := s.newAPIinstance(config)
	if err != nil {
		return nil, err
	}

	return nil, fmt.Errorf("prowlarr custom filters do not work yet")
}
