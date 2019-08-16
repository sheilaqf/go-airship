package airship

import (
	"net/http"

	"github.com/dghubble/sling"
)

// PushService ...
type PushService struct {
	sling *sling.Sling
}

// Push ...
type Push struct {
	// Audience ...
	Audience map[string][]AudienceSelector `json:"audience"`
	// Campaigns ...
	Campaigns []*Campaign `json:"campaigns,omitempty"`
	// DeviceTypes ...
	DeviceTypes []string `json:"device_types"`
}

// InAppMessage ...
type InAppMessage struct {
	// Actions ...
	Actions interface{} `json:"actions,omitempty"`
	// Alert ...
	Alert string `json:"actions"`
	// Display ...
	Display InAppMessageDisplay `json:"display,omitempty"`
}

// InAppMessageDisplay ...
type InAppMessageDisplay struct {
	// Duration ...
	Duration int `json:"duration,omitempty"`
	// Position ...
	Position string `json:"position,omitempty"`
	// PrimaryColor ...
	PrimaryColor string `json:"primary_color,omitempty"`
	// SecondaryColor ...
	SecondaryColor string `json:"secondary_color,omitempty"`
	// DisplayType ...
	DisplayType string `json:"display_type"`
	// Expire ...
	Expire interface{} `json:"expire,omitempty"`
	// Extra ...
	Extra map[string]string `json:"extra,omitempty"`
}

// Actions ...
type Actions struct {
	// AddTag ...
	AddTag []string `json:"add_tag,omitempty"`
	// Open ...
	Open interface{} `json:"open"`
	// RemoveTag ...
	RemoveTag []string `json:"remove_tag,omitempty"`
	// Share ...
	Share string `json:"share,omitempty"`
}

// ActionsOpenURL ...
type ActionsOpenURL struct {
	// Content ...
	Content string `json:"add_tag"`
	// ActionType ...
	ActionType string `json:"type"`
}

// ActionsOpenDeepLink ...
type ActionsOpenDeepLink struct {
	// Content ...
	Content string `json:"add_tag"`
	// ActionType ...
	ActionType string `json:"type"`
	// FallbackURL ...
	FallbackURL string `json:"fallback_url,omitempty"`
}

// ActionsOpenLandingPage ...
type ActionsOpenLandingPage struct {
	// Content ...
	Content string `json:"add_tag"`
	// ActionType ...
	ActionType string `json:"type"`
	// FallbackURL ...
	FallbackURL string `json:"fallback_url,omitempty"`
}

// ActionsOpenLandingPageContent ...
type ActionsOpenLandingPageContent struct {
	// Body ...
	Body string `json:"body"`
	// ContentEncoding ...
	ContentEncoding string `json:"content_encoding,omitempty"`
	// ContentEncoding ...
	ContentType string `json:"content_type,omitempty"`
	// ActionType ...
	ActionType string `json:"type"`
	// FallbackURL ...
	FallbackURL string `json:"fallback_url,omitempty"`
}

// AudienceSelector ...
type AudienceSelector struct {
	// AmazonChannel ...
	AmazonChannel []string `json:"amazon_channel,omitempty"`
	// AndroidChannel ...
	AndroidChannel []string `json:"android_channel,omitempty"`
	// Channel ...
	Channel []string `json:"channel,omitempty"`
	// IOSChannel ...
	IOSChannel []string `json:"ios_channel,omitempty"`
	// Location ...
	Location *Location `json:"location,omitempty"`
	// NamedUser ...
	NamedUser []string `json:"named_user,omitempty"`
	// OpenChannel ...
	OpenChannel []string `json:"open_channel,omitempty"`
	// Segment ...
	Segment []string `json:"segment,omitempty"`
	// SmsID ...
	SmsID *SmsID `json:"sms_id,omitempty"`
	// StaticList ...
	StaticList []string `json:"static_list,omitempty"`
	// Tag ...
	Tag []string `json:"tag,omitempty"`
	// WNS ...
	WNS []string `json:"wns,omitempty"`
	// Date ...
	Date *Date `json:"date,omitempty"`
}

// Date ...
type Date struct {
	// ID ...
	ID string `json:"id,omitempty"`
	// Hours ...
	Hours DateWindow `json:"hours,omitempty"`
	// Days ...
	Days DateWindow `json:"days,omitempty"`
	// Weeks ...
	Weeks DateWindow `json:"weeks,omitempty"`
	// Hours ...
	Months DateWindow `json:"monts,omitempty"`
}

// DateWindow ...
type DateWindow struct {
	// End ...
	End string `json:"end,omitempty"`
	// Start ...
	Start string `json:"start,omitempty"`
	// Recent ...
	Recent int `json:"recent,omitempty"`
}

// Location ...
type Location struct {
}

// SmsID ...
type SmsID struct {
	// MsisDN ...
	MsisDN string `json:"msisdn"`
	// Sender ...
	Sender string `json:"sender"`
}

// Campaign ...
type Campaign struct {
	// Categories ...
	Categories []string `json:"categories"`
}

func newPushService(sling *sling.Sling) *PushService {
	return &PushService{
		sling: sling.Path(PushPath),
	}
}

// Push ...
func (p *PushService) Push(push []*Push) (*Response, error) {
	success := new(Response)
	failure := new(AirshipError)

	res, err := p.sling.New().BodyJSON(push).Receive(success, failure)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusAccepted {
		return nil, failure
	}

	return success, nil
}