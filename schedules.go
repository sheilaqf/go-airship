package airship

import (
	"github.com/dghubble/sling"
)

// SchedulesServices ...
type SchedulesService struct {
	sling *sling.Sling
}

func newSchedulesService(sling *sling.Sling) *SchedulesService {
	return &SchedulesService{
		sling: sling.Path(SchedulesPath),
	}
}
