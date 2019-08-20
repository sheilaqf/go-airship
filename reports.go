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
	OptIns []*ReportsOptIns `url:"optins,omitempty"`
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

// OptIns ...
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

// ReportOptOutsResponse ...
type ReportOptOutsResponse struct {
	// Opens ...
	OptOuts []*ReportsOptOuts `url:"optouts,omitempty"`
	// NextPage ...
	NextPage string `url:"next_page,omitempty"`
}

// ReportsOptOuts ...
type ReportsOptOuts struct {
	// Android ...
	Android int `json:"android"`
	// IOS ...
	IOS int `json:"ios"`
	// Date ...
	Date string `json:"date"`
}

// ReportsOptOutsParams ...
type ReportsOptOutsParams struct {
	// Start ...
	Start string `url:"start"`
	// End ...
	End string `url:"end"`
	// Precision ...
	Precision string `url:"precision"`
}

// OptOuts ...
func (r *ReportsService) OptOuts(params *ReportsOptOutsParams) (*ReportOptOutsResponse, error) {
	success := new(ReportOptOutsResponse)
	failure := new(AirshipError)

	res, err := r.sling.New().Get(path.Join(ReportsPath, ReportsOptOutsPath)).QueryStruct(params).Receive(success, failure)

	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, failure
	}

	return success, nil
}

// ReportsResponsesResponse ...
type ReportsResponsesResponse struct {
	// Opens ...
	Responses []*ReportsResponses `url:"responses,omitempty"`
	// NextPage ...
	NextPage string `url:"next_page,omitempty"`
}

// ReportsResponses ...
type ReportsResponses struct {
	// Android ...
	Android *ReportsResponsesStats `json:"android"`
	// IOS ...
	IOS *ReportsResponsesStats `json:"ios"`
	// Date ...
	Date string `json:"date"`
}

// ReportsResponsesStats ...
type ReportsResponsesStats struct {
	// Direct ...
	Direct int `json:"direct"`
	// Influenced ...
	Influenced int `json:"influenced"`
}

// ReportsResponsesParams ...
type ReportsResponsesParams struct {
	// Start ...
	Start string `url:"start"`
	// End ...
	End string `url:"end"`
	// Precision ...
	Precision string `url:"precision"`
}

// Responses ...
func (r *ReportsService) Responses(params *ReportsResponsesParams) (*ReportsResponsesResponse, error) {
	success := new(ReportsResponsesResponse)
	failure := new(AirshipError)

	res, err := r.sling.New().Get(path.Join(ReportsPath, ReportsOptOutsPath)).QueryStruct(params).Receive(success, failure)

	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, failure
	}

	return success, nil
}

// ReportsResponsesListResponse ...
type ReportsResponsesListResponse struct {
	// Opens ...
	Pushes []*ReportsResponsesList `url:"pushes,omitempty"`
	// NextPage ...
	NextPage string `url:"next_page,omitempty"`
}

// ReportsResponsesList ...
type ReportsResponsesList struct {
	// PushUUID ...
	PushUUID string `json:"push_uuid"`
	// PushTime ...
	PushTime string `json:"push_time"`
	// PushType ...
	PushType string `json:"push_type"`
	// GroupID ...
	GroupID string `json:"group_id"`
	// DirectResponses ...
	DirectResponses int `json:"direct_responses"`
	// Sends ...
	Sends int `json:"sends"`
	// OpenChannelsSends ...
	OpenChannelSends *ReportsOpenChannelSends `json:"open_channel_sends"`
}

// ReportsOpenChannelSends ...
type ReportsOpenChannelSends struct {
	// Plattforms ...
	Platforms []*ReportsOpenChannelSendsPlattform `json:"platforms"`
}

// ReportsOpenChannelSendsPlattform ...
type ReportsOpenChannelSendsPlattform struct {
	// ID ...
	ID string `json:"id"`
	// ID ...
	Sends int `json:"sends"`
}

// ReportsResponsesListParams ...
type ReportsResponsesListParams struct {
	// Start ...
	Start string `url:"start"`
	// End ...
	End string `url:"end"`
	// Precision ...
	Precision string `url:"precision"`
}

// Responses ...
func (r *ReportsService) ResponsesList(params *ReportsResponsesListParams) (*ReportsResponsesListResponse, error) {
	success := new(ReportsResponsesListResponse)
	failure := new(AirshipError)

	res, err := r.sling.New().Get(path.Join(ReportsPath, ReportsResponsesListPath)).QueryStruct(params).Receive(success, failure)

	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, failure
	}

	return success, nil
}
