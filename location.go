package airship

import (
	"github.com/dghubble/sling"
)

// LocationService ...
type LocationService struct {
	sling *sling.Sling
}

func newLocationService(sling *sling.Sling) *LocationService {
	return &LocationService{
		sling: sling.Path(LocationPath),
	}
}
