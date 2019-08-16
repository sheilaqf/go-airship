package airship

import (
	"github.com/dghubble/sling"
)

// ListsService ...
type ListsService struct {
	sling *sling.Sling
}

func newListsService(sling *sling.Sling) *ListsService {
	return &ListsService{
		sling: sling.Path(ListsPath),
	}
}
