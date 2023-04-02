package starrs

import (
	"fmt"
)

func (s *Starrs) BlockList(config *AppConfig) (any, error) {
	s.log.Tracef("Call:BlockList(%s, %s)", config.App, config.Name)

	return nil, fmt.Errorf("block lists do not work yet")
}
