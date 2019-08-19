package airship

import (
	"github.com/dghubble/sling"
)

// SegmentsService ...
type SegmentsService struct {
	sling *sling.Sling
}

func newSegmentsService(sling *sling.Sling) *SegmentsService {
	return &SegmentsService{
		sling: sling,
	}
}
