package airship_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"

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

	Describe("List schedules with only required params", func() {
		var res *airship.ListSchedulesResponse
		var params *airship.ListSchedulesParams
		var resp *airship.ListSchedulesResponse
		var err error

		BeforeEach(func() {
			resp = &airship.ListSchedulesResponse{
				Count: 1,
				OK:    true,
				Schedules: []*airship.Schedule{
					&airship.Schedule{
						Name: "foobar",
					},
				},
			}

			mux.HandleFunc("/api/schedules", func(w http.ResponseWriter, r *http.Request) {
				Ω(r.Method).Should(Equal("GET"))

				p := r.URL.Query()
				Ω(p.Get("start")).Should(Equal(params.Start))
				Ω(p.Get("limit")).Should(Equal(strconv.Itoa(params.Limit)))

				w.Header().Set("Content-Type", "application/vnd.urbanairship+json; version=3;")

				b, err := json.Marshal(resp)
				Ω(err).Should(BeNil())

				w.Write(b)
			})

			params = &airship.ListSchedulesParams{
				Start: "1",
				Limit: 1,
			}

			res, err = client.Schedules.ListSchedules(params)
		})

		It("should not return an error", func() {
			Ω(err).Should(BeNil())
		})

		It("should contain all the information", func() {
			Ω(res).Should(Equal(resp))
		})
	})
})
