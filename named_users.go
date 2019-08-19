package airship

import (
	"github.com/dghubble/sling"
)

// NamedUsersService ...
type NamedUsersService struct {
	sling *sling.Sling
}

func newNamedUsersService(sling *sling.Sling) *NamedUsersService {
	return &NamedUsersService{
		sling: sling,
	}
}
