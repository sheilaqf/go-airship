package airship

import (
	"github.com/dghubble/sling"
)

// ReportsService ...
type ReportsService struct {
	sling *sling.Sling
}

func newReportsService(sling *sling.Sling) *ReportsService {
	return &ReportsService{
		sling: sling.Path(ReportsPath),
	}
}

// Devices ...
func (p *PushService) Devices(push []*Push) (*Response, error) {
	return nil, nil
}
