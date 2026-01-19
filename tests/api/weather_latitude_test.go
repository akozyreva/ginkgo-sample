package api_test

import (
	"encoding/json"
	"fmt"
	"ginkgo-sample/tests/api"
	"io"
	"log"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("WeatherLatitude", func() {
	Context("GET by default", func() {
		It("API returns in response def latitude", func() {
			resp, err := api.SendApiWeatherRequest()
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
			}
			var data map[string]interface{}
			if err := json.Unmarshal(body, &data); err != nil {
				log.Fatal(err)
			}
			fmt.Println(data)
			expLatitude, ok := data["latitude"].(float64)
			Expect(ok).To(Equal(true), "Expected that in resp body there's param latitude")
			Expect(expLatitude).To(Equal(52.52))
		})
	})
})
