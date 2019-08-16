package airship

import (
	"github.com/dghubble/sling"
)

// ChannelsService ...
type ChannelsService struct {
	sling *sling.Sling
}

func newChannelsService(sling *sling.Sling) *ChannelsService {
	return &ChannelsService{
		sling: sling.Path(ChannelsPath),
	}
}
