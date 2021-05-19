package airship_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	airship "github.com/sheilaqf/go-airship"
)

var _ = Describe("Push", func() {

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

	Describe("Send push to devices and send success", func() {
		var res *airship.Response
		var pushs []*airship.Push
		var err error

		BeforeEach(func() {
			res = &airship.Response{
				OK:          true,
				OperationID: "1234",
			}

			pushs = []*airship.Push{
				&airship.Push{
					Audience: map[string][]*airship.AudienceSelector{
						"OR": []*airship.AudienceSelector{
							&airship.AudienceSelector{
								Tag: []string{
									"sports",
									"entertainment",
								},
							},
							&airship.AudienceSelector{
								IOSChannel: []string{
									"871922F4F7C6DF9D51AC7ABAE9AA5FCD7188D7BFA19A2FA99E1D2EC5F2D76506",
								},
							},
						},
					},
				},
			}

			mux.HandleFunc("/api/push", func(w http.ResponseWriter, r *http.Request) {
				Ω(r.Method).Should(Equal("POST"))

				w.Header().Set("Content-Type", "application/vnd.urbanairship+json; version=3;")
				w.WriteHeader(http.StatusAccepted)

				// Read the content
				var b []byte
				Ω(r.Body).ShouldNot(BeNil())

				b, err = ioutil.ReadAll(r.Body)
				Ω(err).Should(BeNil())

				// Restore the io.ReadCloser to its original state
				r.Body = ioutil.NopCloser(bytes.NewBuffer(b))

				var p []*airship.Push
				err = json.Unmarshal(b, &p)
				Ω(err).Should(BeNil())
				Ω(len(pushs)).Should(Equal(len(p)))

				b, err = json.Marshal(res)
				Ω(err).Should(BeNil())

				w.Write(b)
			})

			res, err = client.Push.PostPush(pushs)
		})

		It("should not return an error", func() {
			Ω(err).Should(BeNil())
		})

		It("should contain vaild response", func() {
			Ω(res).Should(Equal(res))
		})
	})
})
