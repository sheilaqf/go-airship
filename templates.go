package airship

import (
	"github.com/dghubble/sling"
)

// TemplatesService ...
type TemplatesService struct {
	sling *sling.Sling
}

func newTemplatesService(sling *sling.Sling) *TemplatesService {
	return &TemplatesService{
		sling: sling.Path(TemplatesPath),
	}
}
