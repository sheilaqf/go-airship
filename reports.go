package airship

import (
	"path"

	"github.com/dghubble/sling"
)

// ReportsService ...
type ReportsService struct {
	sling *sling.Sling
}

func newReportsService(sling *sling.Sling) *ReportsService {
	return &ReportsService{
		sling: sling,
	}
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

// ReportsEventsResponse ...
type ReportsEventsResponse struct {
	// OK ...
	OK bool `json:",omitempty"`
	// TotalValue ...
	TotalValue int `json:"total_value,omitempty"`
	// TotalCount ...
	TotalCount int `json:"total_count,omitempty"`
	// NextPage ...
	NextPage string `json:"next_page,omitempty"`
	// Events ...
	Events []*ReportsEvent `json:"events,omitempty"`
}

// ReportsEvent ...
type ReportsEvent struct {
	// Name ...
	Name string `json:"name,omitempty"`
	// Conversion ...
	Conversion string `json:"conversion,omitempty"`
	// Location ...
	Location string `json:"location,omitempty"`
	// Count ...
	Count int `json:"count,omitempty"`
	// Count ...
	Value int `json:"value,omitempty"`
}

// ReportsEventsParams ...
type ReportsEventsParams struct {
	// Start ...
	Start string `url:"start"`
	// End ...
	End string `url:"end"`
	// Precision ...
	Precision string `url:"precision"`
	// Page ...
	Page int `url:"page,omitempty"`
	// PageSize ...
	PageSize int `url:"page_size,omitempty"`
}

// Events ...
func (r *ReportsService) Events(params *ReportsEventsParams) (*ReportsEventsResponse, error) {
	success := new(ReportsEventsResponse)
	failure := new(AirshipError)

	res, err := r.sling.New().Get(path.Join(ReportsPath, ReportsEventsPath)).QueryStruct(params).Receive(success, failure)

	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, failure
	}

	return success, nil
}

// ReportsAppOpens ...
type ReportsAppOpensResponse struct {
	// Opens ...
	Opens []*ReportsAppOpens `url:"opens,omitempty"`
	// NextPage ...
	NextPage string `url:"next_page,omitempty"`
}

// ReportsAppOpens ...
type ReportsAppOpens struct {
	// Android ...
	Android int `json:"android"`
	// IOS ...
	IOS int `json:"ios"`
	// Date ...
	Date string `json:"date"`
}

// ReportsAppOpensParams ...
type ReportsAppOpensParams struct {
	// Start ...
	Start string `url:"start"`
	// End ...
	End string `url:"end"`
	// Precision ...
	Precision string `url:"precision"`
}

// Opens ...
func (r *ReportsService) Opens(params *ReportsAppOpensParams) (*ReportsAppOpensResponse, error) {
	success := new(ReportsAppOpensResponse)
	failure := new(AirshipError)

	res, err := r.sling.New().Get(path.Join(ReportsPath, ReportsAppOpensPath)).QueryStruct(params).Receive(success, failure)

	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, failure
	}

	return success, nil
}

// ReportOptInsResponse ...
type ReportOptInsResponse struct {
	// Opens ...
	Opens []*ReportsOptIns `url:"optins,omitempty"`
	// NextPage ...
	NextPage string `url:"next_page,omitempty"`
}

// ReportsOptIns ...
type ReportsOptIns struct {
	// Android ...
	Android int `json:"android"`
	// IOS ...
	IOS int `json:"ios"`
	// Date ...
	Date string `json:"date"`
}

// ReportsOptInsParams ...
type ReportsOptInsParams struct {
	// Start ...
	Start string `url:"start"`
	// End ...
	End string `url:"end"`
	// Precision ...
	Precision string `url:"precision"`
}

// Opens ...
func (r *ReportsService) OptIns(params *ReportsOptInsParams) (*ReportOptInsResponse, error) {
	success := new(ReportOptInsResponse)
	failure := new(AirshipError)

	res, err := r.sling.New().Get(path.Join(ReportsPath, ReportsOptInsPath)).QueryStruct(params).Receive(success, failure)

	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, failure
	}

	return success, nil
}
