package airship

import (
	"github.com/dghubble/sling"
)

// SchedulesServices ...
type SchedulesService struct {
	sling *sling.Sling
}

func newSchedulesService(sling *sling.Sling) *SchedulesService {
	return &SchedulesService{
		sling: sling,
	}
}

// ListSchedulesParams ...
type ListSchedulesParams struct {
	// Start ...
	Start string `url:"start,omitempty"`
	// Start ...
	Limit int `url:"limit,omitempty"`
}

// ListSchedulesResponse ...
type ListSchedulesResponse struct {
	// Count ...
	Count int `json:"count,omitempty"`
	// Count ...
	OK bool `json:"ok,omitempty"`
	// NextPage ...
	NextPage string `json:"next_page,omitempty"`
	// Schedules ...
	Schedules []*Schedule `json:"schedules"`
	// TotalCount ...
	TotalCount int `json:"total_count,omitempty"`
}

// Schedule ...
type Schedule struct {
	// Name ...
	Name string `json:"count,omitempty"`
	// Push ...
	Push *Push `json:"push,omitempty"`
	// PushIDs ...
	PushIDs []string `json:"push_ids,omitempty"`
	// URL ...
	URL string `json:"url,omitempty"`
	// Spec ...
	Spec *ScheduleSpec `json:"schedule"`
}

// ScheduleSpec ...
type ScheduleSpec struct {
	// ScheduleTime ...
	ScheduleTime string `json:"scheduled_time,omitempty"`
	// LocalScheduleTime ...
	LocalScheduleTime string `json:"local_scheduled_time,omitempty"`
	// BestTime ...
	BestTime string `json:"best_time,omitempty"`
}

// ListSchedules ...
func (s *SchedulesService) ListSchedules(params *ListSchedulesParams) (*ListSchedulesResponse, error) {
	success := new(ListSchedulesResponse)
	failure := new(AirshipError)

	res, err := s.sling.New().Get(SchedulesPath).QueryStruct(params).Receive(success, failure)

	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, failure
	}

	return success, nil
}
