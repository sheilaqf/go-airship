package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"

	airship "github.com/Onefootball/go-airship"
)

var (
	apiKey    string
	apiSecret string
)

func main() {
	flag.StringVar(&apiKey, "api-key", "", "api key for the API")
	flag.StringVar(&apiSecret, "api-secret", "", "api secret for the api")

	flag.Parse()

	httpClient := &http.Client{}

	// create client
	client := airship.New(httpClient, airship.EndpointURL(airship.AirshipNorthAmericaURL), airship.Auth(apiKey, apiSecret))

	// query the app
	res, err := client.Reports.Devices(&airship.ReportsDevicesParams{})
	if err != nil {
		panic(err)
	}

	b, err := json.Marshal(res)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}

	fmt.Printf(string(b))
}
