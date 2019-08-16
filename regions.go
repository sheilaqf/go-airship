package airship

import (
	"github.com/dghubble/sling"
)

// RegionsService ...
type RegionsService struct {
	sling *sling.Sling
}

func newRegionsService(sling *sling.Sling) *RegionsService {
	return &RegionsService{
		sling: sling.Path(RegionsPath),
	}
}
