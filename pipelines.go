package airship

import (
	"github.com/dghubble/sling"
)

// PipelinesService ...
type PipelinesService struct {
	sling *sling.Sling
}

func newPipelinesService(sling *sling.Sling) *PipelinesService {
	return &PipelinesService{
		sling: sling.Path(PipelinesPath),
	}
}
