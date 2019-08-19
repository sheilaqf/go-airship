package airship

import (
	"github.com/dghubble/sling"
)

// FeedsService ...
type FeedsService struct {
	sling *sling.Sling
}

func newFeedsService(sling *sling.Sling) *FeedsService {
	return &FeedsService{
		sling: sling,
	}
}
