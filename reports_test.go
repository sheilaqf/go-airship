package airship_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"

	airship "github.com/Onefootball/go-airship"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("App", func() {

	var client *airship.Client
	var httpClient *http.Client
	var mux *http.ServeMux
	var server *httptest.Server
	var resp *airship.ReportsDevicesResponse

	BeforeEach(func() {
		httpClient, mux, server = fakeTestServer()

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

		client = airship.New(httpClient, airship.EndpointURL(airship.AirshipNorthAmericaURL))
	})

	AfterEach(func() {
		defer server.Close()
	})

	Describe("Get Device Reports with empty query", func() {
		var res *airship.ReportsDevicesResponse
		var params *airship.ReportsDevicesParams
		var err error

		BeforeEach(func() {
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
})
