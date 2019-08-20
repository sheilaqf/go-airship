package airship_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"

	airship "github.com/Onefootball/go-airship"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Reports", func() {

	var client *airship.Client
	var httpClient *http.Client
	var mux *http.ServeMux
	var server *httptest.Server

	BeforeEach(func() {
		httpClient, mux, server = fakeTestServer()

		client = airship.New(httpClient, airship.EndpointURL(airship.AirshipNorthAmericaURL))
	})

	AfterEach(func() {
		defer server.Close()
	})

	Describe("Get devices reports with empty query", func() {
		var res *airship.ReportsDevicesResponse
		var params *airship.ReportsDevicesParams
		var resp *airship.ReportsDevicesResponse
		var err error

		BeforeEach(func() {
			resp = &airship.ReportsDevicesResponse{
				TotalUniqueDevices: 1,
				DateClosed:         "2019-08-18 00:00:00",
				DateComputed:       "2019-08-19 06:08:04",
				Counts: map[string]map[string]int{
					"ios": map[string]int{
						"unique_devices": 1,
						"opted_in":       1,
						"opted_out":      1,
						"uninstalled":    1,
					},
				},
			}

			mux.HandleFunc("/api/reports/devices", func(w http.ResponseWriter, r *http.Request) {
				Ω(r.Method).Should(Equal("GET"))

				w.Header().Set("Content-Type", "application/vnd.urbanairship+json; version=3;")

				b, err := json.Marshal(resp)
				Ω(err).Should(BeNil())

				w.Write(b)
			})

			params = &airship.ReportsDevicesParams{}

			res, err = client.Reports.Devices(params)
		})

		It("should not return an error", func() {
			Ω(err).Should(BeNil())
		})

		It("should contain all the information", func() {
			Ω(res).Should(Equal(resp))
		})
	})

	Describe("Return an error with the api", func() {
		var params *airship.ReportsDevicesParams
		var airshipErr *airship.AirshipError
		var err error

		BeforeEach(func() {
			airshipErr = &airship.AirshipError{
				OK:      false,
				Message: "airship: fooBar",
			}

			mux.HandleFunc("/api/reports/devices", func(w http.ResponseWriter, r *http.Request) {
				Ω(r.Method).Should(Equal("GET"))

				w.Header().Set("Content-Type", "application/vnd.urbanairship+json; version=3;")
				w.WriteHeader(http.StatusForbidden)

				b, err := json.Marshal(airshipErr)
				Ω(err).Should(BeNil())

				w.Write(b)
			})

			params = &airship.ReportsDevicesParams{}

			_, err = client.Reports.Devices(params)
		})

		It("should not return an error", func() {
			Ω(err).ShouldNot(BeNil())
		})

		It("should contain all the information", func() {
			Ω(err).Should(Equal(airshipErr))
		})
	})

	Describe("Get events report with only required params", func() {
		var res *airship.ReportsEventsResponse
		var params *airship.ReportsEventsParams
		var resp *airship.ReportsEventsResponse
		var err error

		BeforeEach(func() {
			resp = &airship.ReportsEventsResponse{
				OK:         true,
				TotalValue: 1,
				TotalCount: 1,
				NextPage:   "https://go.urbanairship.com/api/reports/events?start=2018-08-01T10:00:00.000Z&end=2018-08-15T20:00:00.000Z&precision=MONTHLY&page_size=20&page=2",
				Events: []*airship.ReportsEvent{
					&airship.ReportsEvent{
						Name:       "banner_image",
						Conversion: "indirect",
						Count:      1,
						Value:      1,
					},
				},
			}

			mux.HandleFunc("/api/reports/events", func(w http.ResponseWriter, r *http.Request) {
				Ω(r.Method).Should(Equal("GET"))

				p := r.URL.Query()
				Ω(p.Get("start")).Should(Equal(params.Start))
				Ω(p.Get("end")).Should(Equal(params.End))
				Ω(p.Get("precision")).Should(Equal(params.Precision))

				w.Header().Set("Content-Type", "application/vnd.urbanairship+json; version=3;")

				b, err := json.Marshal(resp)
				Ω(err).Should(BeNil())

				w.Write(b)
			})

			params = &airship.ReportsEventsParams{
				Start:     "2018-08-28 00:00:00",
				End:       "2018-08-29 13:30:45",
				Precision: "HOURLY",
			}

			res, err = client.Reports.Events(params)
		})

		It("should not return an error", func() {
			Ω(err).Should(BeNil())
		})

		It("should contain all the information", func() {
			Ω(res).Should(Equal(resp))
		})
	})

	Describe("Get app opens report with only required params", func() {
		var res *airship.ReportsAppOpensResponse
		var params *airship.ReportsAppOpensParams
		var resp *airship.ReportsAppOpensResponse
		var err error

		BeforeEach(func() {
			resp = &airship.ReportsAppOpensResponse{
				Opens: []*airship.ReportsAppOpens{
					&airship.ReportsAppOpens{
						Android: 1,
						IOS:     1,
						Date:    "2018-08-28 13:30:45",
					},
				},
			}

			mux.HandleFunc("/api/reports/opens", func(w http.ResponseWriter, r *http.Request) {
				Ω(r.Method).Should(Equal("GET"))

				p := r.URL.Query()
				Ω(p.Get("start")).Should(Equal(params.Start))
				Ω(p.Get("end")).Should(Equal(params.End))
				Ω(p.Get("precision")).Should(Equal(params.Precision))

				w.Header().Set("Content-Type", "application/vnd.urbanairship+json; version=3;")

				b, err := json.Marshal(resp)
				Ω(err).Should(BeNil())

				w.Write(b)
			})

			params = &airship.ReportsAppOpensParams{
				Start:     "2018-08-28 00:00:00",
				End:       "2018-08-29 13:30:45",
				Precision: "HOURLY",
			}

			res, err = client.Reports.Opens(params)
		})

		It("should not return an error", func() {
			Ω(err).Should(BeNil())
		})

		It("should contain all the information", func() {
			Ω(res).Should(Equal(resp))
		})
	})

	Describe("Get OptIns report with only required params", func() {
		var res *airship.ReportOptInsResponse
		var params *airship.ReportsOptInsParams
		var resp *airship.ReportOptInsResponse
		var err error

		BeforeEach(func() {
			resp = &airship.ReportOptInsResponse{
				OptIns: []*airship.ReportsOptIns{
					&airship.ReportsOptIns{
						Android: 1,
						IOS:     1,
						Date:    "2018-08-28 13:30:45",
					},
				},
			}

			mux.HandleFunc("/api/reports/optins", func(w http.ResponseWriter, r *http.Request) {
				Ω(r.Method).Should(Equal("GET"))

				p := r.URL.Query()
				Ω(p.Get("start")).Should(Equal(params.Start))
				Ω(p.Get("end")).Should(Equal(params.End))
				Ω(p.Get("precision")).Should(Equal(params.Precision))

				w.Header().Set("Content-Type", "application/vnd.urbanairship+json; version=3;")

				b, err := json.Marshal(resp)
				Ω(err).Should(BeNil())

				w.Write(b)
			})

			params = &airship.ReportsOptInsParams{
				Start:     "2018-08-28 00:00:00",
				End:       "2018-08-29 13:30:45",
				Precision: "HOURLY",
			}

			res, err = client.Reports.OptIns(params)
		})

		It("should not return an error", func() {
			Ω(err).Should(BeNil())
		})

		It("should contain all the information", func() {
			Ω(res).Should(Equal(resp))
		})
	})

	Describe("Get OptOuts report with only required params", func() {
		var res *airship.ReportOptOutsResponse
		var params *airship.ReportsOptOutsParams
		var resp *airship.ReportOptOutsResponse
		var err error

		BeforeEach(func() {
			resp = &airship.ReportOptOutsResponse{
				OptOuts: []*airship.ReportsOptOuts{
					&airship.ReportsOptOuts{
						Android: 1,
						IOS:     1,
						Date:    "2018-08-28 13:30:45",
					},
				},
			}

			mux.HandleFunc("/api/reports/optouts", func(w http.ResponseWriter, r *http.Request) {
				Ω(r.Method).Should(Equal("GET"))

				p := r.URL.Query()
				Ω(p.Get("start")).Should(Equal(params.Start))
				Ω(p.Get("end")).Should(Equal(params.End))
				Ω(p.Get("precision")).Should(Equal(params.Precision))

				w.Header().Set("Content-Type", "application/vnd.urbanairship+json; version=3;")

				b, err := json.Marshal(resp)
				Ω(err).Should(BeNil())

				w.Write(b)
			})

			params = &airship.ReportsOptOutsParams{
				Start:     "2018-08-28 00:00:00",
				End:       "2018-08-29 13:30:45",
				Precision: "HOURLY",
			}

			res, err = client.Reports.OptOuts(params)
		})

		It("should not return an error", func() {
			Ω(err).Should(BeNil())
		})

		It("should contain all the information", func() {
			Ω(res).Should(Equal(resp))
		})
	})

	Describe("Get responses report with only required params", func() {
		var res *airship.ReportsResponsesResponse
		var params *airship.ReportsResponsesParams
		var resp *airship.ReportsResponsesResponse
		var err error

		BeforeEach(func() {
			resp = &airship.ReportsResponsesResponse{
				Responses: []*airship.ReportsResponses{
					&airship.ReportsResponses{
						Android: &airship.ReportsResponsesStats{
							Influenced: 1,
							Direct:     1,
						},
						IOS: &airship.ReportsResponsesStats{
							Influenced: 1,
							Direct:     1,
						},
						Date: "2018-08-28 13:30:45",
					},
				},
			}

			mux.HandleFunc("/api/reports/responses", func(w http.ResponseWriter, r *http.Request) {
				Ω(r.Method).Should(Equal("GET"))

				p := r.URL.Query()
				Ω(p.Get("start")).Should(Equal(params.Start))
				Ω(p.Get("end")).Should(Equal(params.End))
				Ω(p.Get("precision")).Should(Equal(params.Precision))

				w.Header().Set("Content-Type", "application/vnd.urbanairship+json; version=3;")

				b, err := json.Marshal(resp)
				Ω(err).Should(BeNil())

				w.Write(b)
			})

			params = &airship.ReportsResponsesParams{
				Start:     "2018-08-28 00:00:00",
				End:       "2018-08-29 13:30:45",
				Precision: "HOURLY",
			}

			res, err = client.Reports.Responses(params)
		})

		It("should not return an error", func() {
			Ω(err).Should(BeNil())
		})

		It("should contain all the information", func() {
			Ω(res).Should(Equal(resp))
		})
	})

	Describe("Get responses list report with only required params", func() {
		var res *airship.ReportsResponsesListResponse
		var params *airship.ReportsResponsesListParams
		var resp *airship.ReportsResponsesListResponse
		var err error

		BeforeEach(func() {
			resp = &airship.ReportsResponsesListResponse{
				Pushes: []*airship.ReportsResponsesList{
					&airship.ReportsResponsesList{
						PushUUID:        "f4db3752-a982-4a2b-994e-7b5fd1c7f02f",
						PushTime:        "2018-08-15 02:12:22",
						PushType:        "UNICAST_PUSH",
						GroupID:         "4e768dc7-4ebc-4206-890a-60b5627763a7",
						DirectResponses: 0,
						Sends:           1,
						OpenChannelSends: &airship.ReportsOpenChannelSends{
							Platforms: []*airship.ReportsOpenChannelSendsPlattform{},
						},
					},
				},
			}

			mux.HandleFunc("/api/reports/responses/list", func(w http.ResponseWriter, r *http.Request) {
				Ω(r.Method).Should(Equal("GET"))

				p := r.URL.Query()
				Ω(p.Get("start")).Should(Equal(params.Start))
				Ω(p.Get("end")).Should(Equal(params.End))
				Ω(p.Get("precision")).Should(Equal(params.Precision))

				w.Header().Set("Content-Type", "application/vnd.urbanairship+json; version=3;")

				b, err := json.Marshal(resp)
				Ω(err).Should(BeNil())

				w.Write(b)
			})

			params = &airship.ReportsResponsesListParams{
				Start:     "2018-08-28 00:00:00",
				End:       "2018-08-29 13:30:45",
				Precision: "HOURLY",
			}

			res, err = client.Reports.ResponsesList(params)
		})

		It("should not return an error", func() {
			Ω(err).Should(BeNil())
		})

		It("should contain all the information", func() {
			Ω(res).Should(Equal(resp))
		})
	})

	Describe("Get pushes send report with only required params", func() {
		var res *airship.ReportsSendsResponse
		var params *airship.ReportsSendsParams
		var resp *airship.ReportsSendsResponse
		var err error

		BeforeEach(func() {
			resp = &airship.ReportsSendsResponse{
				Sends: []*airship.ReportsSends{
					&airship.ReportsSends{
						Android: 1,
						IOS:     1,
						Date:    "2018-05-01 00:00:00",
					},
				},
			}

			mux.HandleFunc("/api/reports/sends", func(w http.ResponseWriter, r *http.Request) {
				Ω(r.Method).Should(Equal("GET"))

				p := r.URL.Query()
				Ω(p.Get("start")).Should(Equal(params.Start))
				Ω(p.Get("end")).Should(Equal(params.End))
				Ω(p.Get("precision")).Should(Equal(params.Precision))

				w.Header().Set("Content-Type", "application/vnd.urbanairship+json; version=3;")

				b, err := json.Marshal(resp)
				Ω(err).Should(BeNil())

				w.Write(b)
			})

			params = &airship.ReportsSendsParams{
				Start:     "2018-08-28 00:00:00",
				End:       "2018-08-29 13:30:45",
				Precision: "HOURLY",
			}

			res, err = client.Reports.Sends(params)
		})

		It("should not return an error", func() {
			Ω(err).Should(BeNil())
		})

		It("should contain all the information", func() {
			Ω(res).Should(Equal(resp))
		})
	})
})
