package airship

import (
	"path"

	"github.com/dghubble/sling"
)

// ReportsService ...
type ReportsService struct {
	sling *sling.Sling
}

// ReportsDevicesResponse ...
type ReportsDevicesResponse struct {
	// TotalUniqueDevices ...
	TotalUniqueDevices int `json:"total_unique_devices"`
	// DateClosed ...
	DateClosed string `json:"date_closed"`
	// DateComputed ...
	DateComputed string `json:"date_computed"`
	// Counts ...
	Counts map[string]map[string]int `json:"counts"`
}

// ReportDevicesParams ...
type ReportsDevicesParams struct {
	// Date ...
	Date string `url:"date,omitempty"`
}

func newReportsService(sling *sling.Sling) *ReportsService {
	return &ReportsService{
		sling: sling,
	}
}

// Devices ...
func (r *ReportsService) Devices(params *ReportsDevicesParams) (*ReportsDevicesResponse, error) {
	success := new(ReportsDevicesResponse)
	failure := new(AirshipError)

	res, err := r.sling.New().Get(path.Join(ReportsPath, ReportsDevicesPath)).QueryStruct(params).Receive(success, failure)

	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, failure
	}

	return success, nil
}
