package api_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Data-driven test suite", func() {
	It("Data driven test sample ", func() {
		By("1 step")
		Expect(true).To(BeTrue())
		By("Continue to do something")
		Expect(false).To(BeTrue())
		By("End")
	})

})
