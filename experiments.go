package airship

import (
	"github.com/dghubble/sling"
)

// ExperimentsService ...
type ExperimentsService struct {
	sling *sling.Sling
}

func newExperimentsService(sling *sling.Sling) *ExperimentsService {
	return &ExperimentsService{
		sling: sling.Path(ExperimentsPath),
	}
}
