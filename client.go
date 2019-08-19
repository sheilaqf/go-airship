package airship

import (
	"net/http"

	"github.com/dghubble/sling"
)

const (
	// AirshipNorthAmericaURL ...
	AirshipNorthAmericaURL = "https://go.urbanairship.com/api/"
	// AirshipEuropeURL ...
	AirshipEuropeURL = "https://go.airship.eu/api/"
)

const (
	// AirshipAcceptHeader this is basically doing the versioning of the API requests
	AirshipAcceptHeader = "application/vnd.urbanairship+json; version=3;"
)

const (
	// PushPath ...
	PushPath = "push"
	// ReportsPath ...
	ReportsPath = "reports"
	// ReportsDevicesPath ...
	ReportsDevicesPath = "devices"
	// FeedsPath ...
	FeedsPath = "feeds"
	// SchedulesPath ...
	SchedulesPath = "schedules"
	// PipelinesPath ...
	PipelinesPath = "pipelines"
	// ExperimentsPath ...
	ExperimentsPath = "experiments"
	// TemplatesPath ...
	TemplatesPath = "templates"
	// RegionsPath ...
	RegionsPath = "regions"
	// ChannelsPath ...
	ChannelsPath = "channels"
	// NamedUsersPath ...
	NamedUsersPath = "named_users"
	// SegmentsPath ...
	SegmentsPath = "segements"
	// LocationPath ...
	LocationPath = "location"
	// ListsPath ...
	ListsPath = "lists"
	// CreateAndSendPath ...
	CreateAndSendPath = "create-and-send"
)

// Client holds the client for localytics
type Client struct {
	sling *sling.Sling
	opts  *Opts

	Push          *PushService
	Reports       *ReportsService
	Feeds         *FeedsService
	Schedules     *SchedulesService
	Pipelines     *PipelinesService
	Experiments   *ExperimentsService
	Templates     *TemplatesService
	Regions       *RegionsService
	Channels      *ChannelsService
	NamedUsers    *NamedUsersService
	Location      *LocationService
	Lists         *ListsService
	CreateAndSend *CreateAndSendService
}

// Opt is an option for the client.
type Opt func(*Opts)

// Opts holds alls the options for the client.
type Opts struct {
	ApiKey      string
	ApiSecret   string
	EndpointURL string
}

// Response ...
type Response struct {
	// OK ...
	OK bool `json:"ok"`
	// OperationID ...
	OperationID string `json:"operation_id,omitempty"`
}

// Auth configures the API Access Key and API Master Secret.
func Auth(key string, secret string) func(o *Opts) {
	return func(o *Opts) {
		o.ApiKey = key
		o.ApiSecret = secret
	}
}

// EndpointURL configures the API specific endpoint.
// There are shared constants for the already available ones.
func EndpointURL(endpointURL string) func(o *Opts) {
	return func(o *Opts) {
		o.EndpointURL = endpointURL
	}
}

// New returns a new instance of the client. See /examples for the use
//
//	httpClient := &http.Client{}
//
//	client := airship.New(httpClient, airship.Auth(apiKey, apiMasterKey))
//
//	res, err := client.Apps()
//
func New(httpClient *http.Client, opts ...Opt) *Client {
	return mustNew(httpClient, opts...)
}

// this creates a new client for the API
func mustNew(httpClient *http.Client, opts ...Opt) *Client {
	options := new(Opts)

	c := new(Client)
	c.opts = options

	configure(c, opts...)
	configureSling(c, httpClient)

	// attaching the services ...
	c.Push = newPushService(c.sling.New())
	c.Reports = newReportsService(c.sling.New())
	c.Feeds = newFeedsService(c.sling.New())
	c.Schedules = newSchedulesService(c.sling.New())
	c.Pipelines = newPipelinesService(c.sling.New())
	c.Experiments = newExperimentsService(c.sling.New())
	c.Templates = newTemplatesService(c.sling.New())
	c.Regions = newRegionsService(c.sling.New())
	c.Channels = newChannelsService(c.sling.New())
	c.NamedUsers = newNamedUsersService(c.sling.New())
	c.Location = newLocationService(c.sling.New())
	c.Lists = newListsService(c.sling.New())
	c.CreateAndSend = newCreateAndSendService(c.sling.New())

	return c
}

func configureSling(c *Client, httpClient *http.Client) error {
	c.sling = sling.New()
	c.sling.Client(httpClient).SetBasicAuth(c.opts.ApiKey, c.opts.ApiSecret)
	c.sling.Set("Accept", "application/vnd.localytics.v1+hal+json")
	c.sling.Base(c.opts.EndpointURL)

	return nil
}

func configure(c *Client, opts ...Opt) error {
	for _, o := range opts {
		o(c.opts)
	}

	return nil
}
