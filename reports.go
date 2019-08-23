package airship

import (
	"net/http"
	"path"

	"github.com/dghubble/sling"
)

var (
	// ReportsDevicesPath ...
	ReportsDevicesPath = path.Join(ReportsPath, "devices")
	// ReportsEventsPath ...
	ReportsEventsPath = path.Join(ReportsPath, "events")
	// ReportsAppOpensPath ...
	ReportsAppOpensPath = path.Join(ReportsPath, "opens")
	// ReportsOptInsPath ...
	ReportsOptInsPath = path.Join(ReportsPath, "optins")
	// ReportsOptInsPath ...
	ReportsOptOutsPath = path.Join(ReportsPath, "optouts")
	// ReportsResponsesPath ...
	ReportsResponsesPath = path.Join(ReportsPath, "responses")
	// ReportsResponseListPath ...
	ReportsResponsesListPath = path.Join(ReportsResponsesPath, "list")
	// ReportsSendsPaths ...
	ReportsSendsPath = path.Join(ReportsPath, "sends")
	// ReportsTimeInAppPath ...
	ReportsTimeInAppPath = path.Join(ReportsPath, "timeinapp")
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

// GetDevicesResponse ...
type GetDevicesResponse struct {
	// TotalUniqueDevices ...
	TotalUniqueDevices int `json:"total_unique_devices"`
	// DateClosed ...
	DateClosed string `json:"date_closed"`
	// DateComputed ...
	DateComputed string `json:"date_computed"`
	// Counts ...
	Counts map[string]map[string]int `json:"counts"`
}

// GetDevicesParams ...
type GetDevicesParams struct {
	// Date ...
	Date string `url:"date,omitempty"`
}

// Devices ...
func (r *ReportsService) GetDevices(params *GetDevicesParams) (*GetDevicesResponse, error) {
	success := new(GetDevicesResponse)
	failure := new(AirshipError)

	res, err := r.sling.New().Get(ReportsDevicesPath).QueryStruct(params).Receive(success, failure)

	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, failure
	}

	return success, nil
}

// GetEventsResponse ...
type GetEventsResponse struct {
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

// GetEventsParams ...
type GetEventsParams struct {
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
func (r *ReportsService) GetEvents(params *GetEventsParams) (*GetEventsResponse, error) {
	success := new(GetEventsResponse)
	failure := new(AirshipError)

	res, err := r.sling.New().Get(ReportsEventsPath).QueryStruct(params).Receive(success, failure)

	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, failure
	}

	return success, nil
}

// ListAppOpens ...
type ListAppOpensResponse struct {
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

// ListAppOpensParams ...
type ListAppOpensParams struct {
	// Start ...
	Start string `url:"start"`
	// End ...
	End string `url:"end"`
	// Precision ...
	Precision string `url:"precision"`
}

// Opens ...
func (r *ReportsService) ListOpens(params *ListAppOpensParams) (*ListAppOpensResponse, error) {
	success := new(ListAppOpensResponse)
	failure := new(AirshipError)

	res, err := r.sling.New().Get(ReportsAppOpensPath).QueryStruct(params).Receive(success, failure)

	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, failure
	}

	return success, nil
}

// ListOptInsResponse ...
type ListOptInsResponse struct {
	// OptIns ...
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

// ListOptInsParams ...
type ListOptInsParams struct {
	// Start ...
	Start string `url:"start"`
	// End ...
	End string `url:"end"`
	// Precision ...
	Precision string `url:"precision"`
}

// OptIns ...
func (r *ReportsService) ListOptIns(params *ListOptInsParams) (*ListOptInsResponse, error) {
	success := new(ListOptInsResponse)
	failure := new(AirshipError)

	res, err := r.sling.New().Get(ReportsOptInsPath).QueryStruct(params).Receive(success, failure)

	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, failure
	}

	return success, nil
}

// ListOptOutsResponse ...
type ListOptOutsResponse struct {
	// OptOuts ...
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

// ListOptOutsParams ...
type ListOptOutsParams struct {
	// Start ...
	Start string `url:"start"`
	// End ...
	End string `url:"end"`
	// Precision ...
	Precision string `url:"precision"`
}

// OptOuts ...
func (r *ReportsService) ListOptOuts(params *ListOptOutsParams) (*ListOptOutsResponse, error) {
	success := new(ListOptOutsResponse)
	failure := new(AirshipError)

	res, err := r.sling.New().Get(ReportsOptOutsPath).QueryStruct(params).Receive(success, failure)

	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, failure
	}

	return success, nil
}

// ListResponsesResponse ...
type ListResponsesResponse struct {
	// Responses ...
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

// ListResponsesParams ...
type ListResponsesParams struct {
	// Start ...
	Start string `url:"start"`
	// End ...
	End string `url:"end"`
	// Precision ...
	Precision string `url:"precision"`
}

// Responses ...
func (r *ReportsService) ListResponses(params *ListResponsesParams) (*ListResponsesResponse, error) {
	success := new(ListResponsesResponse)
	failure := new(AirshipError)

	res, err := r.sling.New().Get(ReportsResponsesPath).QueryStruct(params).Receive(success, failure)

	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, failure
	}

	return success, nil
}

// ListResponsePushesResponse ...
type ListResponsePushesResponse struct {
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

// ListResponsePushesParams ...
type ListResponsePushesParams struct {
	// Start ...
	Start string `url:"start"`
	// End ...
	End string `url:"end"`
	// Precision ...
	Precision string `url:"precision"`
}

// Responses ...
func (r *ReportsService) ListReponsePushes(params *ListResponsePushesParams) (*ListResponsePushesResponse, error) {
	success := new(ListResponsePushesResponse)
	failure := new(AirshipError)

	res, err := r.sling.New().Get(ReportsResponsesListPath).QueryStruct(params).Receive(success, failure)

	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, failure
	}

	return success, nil
}

// ListSendsResponse ...
type ListSendsResponse struct {
	// Sends ...
	Sends []*ReportsSends `url:"sends,omitempty"`
	// NextPage ...
	NextPage string `url:"next_page,omitempty"`
}

// ReportsSends ...
type ReportsSends struct {
	// Android ...
	Android int `json:"android"`
	// IOS ...
	IOS int `json:"iods"`
	// Date ...
	Date string `json:"date"`
}

// ListSendsParams ...
type ListSendsParams struct {
	// Start ...
	Start string `url:"start"`
	// End ...
	End string `url:"end"`
	// Precision ...
	Precision string `url:"precision"`
}

// Sends ...
func (r *ReportsService) ListSends(params *ListSendsParams) (*ListSendsResponse, error) {
	success := new(ListSendsResponse)
	failure := new(AirshipError)

	res, err := r.sling.New().Get(ReportsSendsPath).QueryStruct(params).Receive(success, failure)

	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, failure
	}

	return success, nil
}

// ListTimeInAppResponse ...
type ListTimeInAppResponse struct {
	// Sends ...
	Sends []*ReportsSends `url:"sends,omitempty"`
	// NextPage ...
	NextPage string `url:"next_page,omitempty"`
}

// ListTimeInAppParams ...
type ListTimeInAppParams struct {
	// Start ...
	Start string `url:"start"`
	// End ...
	End string `url:"end"`
	// Precision ...
	Precision string `url:"precision"`
}

// ListTimeInApp ...
func (r *ReportsService) ListTimeInApp(params *ListTimeInAppParams) (*ListTimeInAppResponse, error) {
	success := new(ListTimeInAppResponse)
	failure := new(AirshipError)

	res, err := r.sling.New().Get(ReportsTimeInAppPath).QueryStruct(params).Receive(success, failure)

	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, failure
	}

	return success, nil
}
