package main_test

import (
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
