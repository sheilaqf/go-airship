package airship

import (
	"github.com/dghubble/sling"
)

// CreateAndSendService ...
type CreateAndSendService struct {
	sling *sling.Sling
}

func newCreateAndSendService(sling *sling.Sling) *CreateAndSendService {
	return &CreateAndSendService{
		sling: sling,
	}
}
