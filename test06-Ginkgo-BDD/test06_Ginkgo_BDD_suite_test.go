package main

import (
	"log"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)



//ginkgo docs https://onsi.github.io/ginkgo/#your-first-ginkgo-suite
//ginkgo bootstrap (Cmder) generates *_suite_test.go file
func TestTest06GinkgoBDD(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Test06GinkgoBDD Suite")
}

var _ = Describe("Person.IsChild", func() {
	Context("when the person is a child", func() {
		It("returns true", func() {
			person := Person{Age: 10}
			response := person.IsChild()
			Expect(response).To(BeTrue())
			//Expect(response).To(BeFalse())
		})
	})

	Context("when the person is not a child", func() {
		BeforeEach(func() {
			log.Print("not a child")
		})
		AfterEach(func() {
			log.Print("not a child")
		})

		//XIt means skip the test block
		//FIt means skip the before test block
		XIt("returns true", func() {
			person := Person{Age: 100}
			response := person.IsChild()
			Expect(response).To(BeFalse())
			//Expect(response).To(BeFalse())
		})
	})

	



})


