package api_test

import (
	"fmt"
	"io"
	"log"

	"net/http"

	"ginkgo-sample/tests/api"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Wetter API common check", func() {
	var resp *http.Response
	var err error
	var url string
	AfterEach(func() {
		fmt.Println("Running AfterEach hook")
		// clean up reading Body
		DeferCleanup(func() {
			defer func(Body io.ReadCloser) {
				err := Body.Close()
				if err != nil {
				}
			}(resp.Body)
		})

	})

	Context("GET by default", func() {
		It("API returns 200 by default", func() {
			resp, err = http.Get(api.OpenMeteoURL)
			if err != nil {
				panic(err)
			}
			Expect(resp.StatusCode).To(Equal(200))
		})
	})

	Context("When default url is wrong", func() {
		// when you need specifically to do something right before the test
		JustBeforeEach(func() {
			url = api.OpenMeteoURL + ",wrongUrl"
		})
		It("GET returns 400 ", func() {
			resp, err = http.Get(url)
			if err != nil {
				panic(err)
			}
			Expect(resp.StatusCode).To(Equal(400))
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
			}
			GinkgoWriter.Println(string(body))
		})
	})

	Context("Wne you set fail explicitly", func() {
		It("Test will fail always ", func() {
			Fail("This test fails")
		})

		It("This test won't be executed, because previous one failed ", func() {
			Expect(true).To(BeTrue())
		})
	})

})
