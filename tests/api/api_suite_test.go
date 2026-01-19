package api_test

import (
	"fmt"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestApi(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Api Suite")
}

var _ = BeforeSuite(func() {
	fmt.Println("Running BeforeSuite hook")
})

var _ = AfterSuite(func() {
	fmt.Println("Running AfterSuite hook")
})
