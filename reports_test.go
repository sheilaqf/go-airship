package airship_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	airship "github.com/sheilaqf/go-airship"
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
		var res *airship.GetDevicesResponse
		var params *airship.GetDevicesParams
		var resp *airship.GetDevicesResponse
		var err error

		BeforeEach(func() {
			resp = &airship.GetDevicesResponse{
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

			params = &airship.GetDevicesParams{}

			res, err = client.Reports.GetDevices(params)
		})

		It("should not return an error", func() {
			Ω(err).Should(BeNil())
		})

		It("should contain all the information", func() {
			Ω(res).Should(Equal(resp))
		})
	})

	Describe("Return an error with the api", func() {
		var params *airship.GetDevicesParams
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

			params = &airship.GetDevicesParams{}

			_, err = client.Reports.GetDevices(params)
		})

		It("should not return an error", func() {
			Ω(err).ShouldNot(BeNil())
		})

		It("should contain all the information", func() {
			Ω(err).Should(Equal(airshipErr))
		})
	})

	Describe("Get events report with only required params", func() {
		var res *airship.GetEventsResponse
		var params *airship.GetEventsParams
		var resp *airship.GetEventsResponse
		var err error

		BeforeEach(func() {
			resp = &airship.GetEventsResponse{
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

			params = &airship.GetEventsParams{
				Start:     "2018-08-28 00:00:00",
				End:       "2018-08-29 13:30:45",
				Precision: "HOURLY",
			}

			res, err = client.Reports.GetEvents(params)
		})

		It("should not return an error", func() {
			Ω(err).Should(BeNil())
		})

		It("should contain all the information", func() {
			Ω(res).Should(Equal(resp))
		})
	})

	Describe("Get app opens report with only required params", func() {
		var res *airship.ListAppOpensResponse
		var params *airship.ListAppOpensParams
		var resp *airship.ListAppOpensResponse
		var err error

		BeforeEach(func() {
			resp = &airship.ListAppOpensResponse{
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

			params = &airship.ListAppOpensParams{
				Start:     "2018-08-28 00:00:00",
				End:       "2018-08-29 13:30:45",
				Precision: "HOURLY",
			}

			res, err = client.Reports.ListOpens(params)
		})

		It("should not return an error", func() {
			Ω(err).Should(BeNil())
		})

		It("should contain all the information", func() {
			Ω(res).Should(Equal(resp))
		})
	})

	Describe("Get OptIns report with only required params", func() {
		var res *airship.ListOptInsResponse
		var params *airship.ListOptInsParams
		var resp *airship.ListOptInsResponse
		var err error

		BeforeEach(func() {
			resp = &airship.ListOptInsResponse{
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

			params = &airship.ListOptInsParams{
				Start:     "2018-08-28 00:00:00",
				End:       "2018-08-29 13:30:45",
				Precision: "HOURLY",
			}

			res, err = client.Reports.ListOptIns(params)
		})

		It("should not return an error", func() {
			Ω(err).Should(BeNil())
		})

		It("should contain all the information", func() {
			Ω(res).Should(Equal(resp))
		})
	})

	Describe("Get OptOuts report with only required params", func() {
		var res *airship.ListOptOutsResponse
		var params *airship.ListOptOutsParams
		var resp *airship.ListOptOutsResponse
		var err error

		BeforeEach(func() {
			resp = &airship.ListOptOutsResponse{
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

			params = &airship.ListOptOutsParams{
				Start:     "2018-08-28 00:00:00",
				End:       "2018-08-29 13:30:45",
				Precision: "HOURLY",
			}

			res, err = client.Reports.ListOptOuts(params)
		})

		It("should not return an error", func() {
			Ω(err).Should(BeNil())
		})

		It("should contain all the information", func() {
			Ω(res).Should(Equal(resp))
		})
	})

	Describe("Get responses report with only required params", func() {
		var res *airship.ListResponsesResponse
		var params *airship.ListResponsesParams
		var resp *airship.ListResponsesResponse
		var err error

		BeforeEach(func() {
			resp = &airship.ListResponsesResponse{
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

			params = &airship.ListResponsesParams{
				Start:     "2018-08-28 00:00:00",
				End:       "2018-08-29 13:30:45",
				Precision: "HOURLY",
			}

			res, err = client.Reports.ListResponses(params)
		})

		It("should not return an error", func() {
			Ω(err).Should(BeNil())
		})

		It("should contain all the information", func() {
			Ω(res).Should(Equal(resp))
		})
	})

	Describe("Get responses list report with only required params", func() {
		var res *airship.ListResponsePushesResponse
		var params *airship.ListResponsePushesParams
		var resp *airship.ListResponsePushesResponse
		var err error

		BeforeEach(func() {
			resp = &airship.ListResponsePushesResponse{
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

			params = &airship.ListResponsePushesParams{
				Start:     "2018-08-28 00:00:00",
				End:       "2018-08-29 13:30:45",
				Precision: "HOURLY",
			}

			res, err = client.Reports.ListReponsePushes(params)
		})

		It("should not return an error", func() {
			Ω(err).Should(BeNil())
		})

		It("should contain all the information", func() {
			Ω(res).Should(Equal(resp))
		})
	})

	Describe("Get pushes send report with only required params", func() {
		var res *airship.ListSendsResponse
		var params *airship.ListSendsParams
		var resp *airship.ListSendsResponse
		var err error

		BeforeEach(func() {
			resp = &airship.ListSendsResponse{
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

			params = &airship.ListSendsParams{
				Start:     "2018-08-28 00:00:00",
				End:       "2018-08-29 13:30:45",
				Precision: "HOURLY",
			}

			res, err = client.Reports.ListSends(params)
		})

		It("should not return an error", func() {
			Ω(err).Should(BeNil())
		})

		It("should contain all the information", func() {
			Ω(res).Should(Equal(resp))
		})
	})

	Describe("Get time in app report with only required params", func() {
		var res *airship.ListTimeInAppResponse
		var params *airship.ListTimeInAppParams
		var resp *airship.ListTimeInAppResponse
		var err error

		BeforeEach(func() {
			resp = &airship.ListTimeInAppResponse{
				Sends: []*airship.ReportsSends{
					&airship.ReportsSends{
						Android: 1,
						IOS:     1,
						Date:    "2018-05-01 00:00:00",
					},
				},
			}

			mux.HandleFunc("/api/reports/timeinapp", func(w http.ResponseWriter, r *http.Request) {
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

			params = &airship.ListTimeInAppParams{
				Start:     "2018-08-28 00:00:00",
				End:       "2018-08-29 13:30:45",
				Precision: "HOURLY",
			}

			res, err = client.Reports.ListTimeInApp(params)
		})

		It("should not return an error", func() {
			Ω(err).Should(BeNil())
		})

		It("should contain all the information", func() {
			Ω(res).Should(Equal(resp))
		})
	})
})
