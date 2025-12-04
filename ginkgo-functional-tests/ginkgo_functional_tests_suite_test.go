package ginkgo_functional_tests_test

import (
	ginkgo_functional_tests "examples.com/go-miscellaneous/ginkgo-functional-tests"
	"github.com/labstack/gommon/log"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestGinkgoFunctionalTests(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GinkgoFunctionalTests Suite")
}

var _ = Describe("Person is child or not", Ordered, func() {
	BeforeAll(func() {
		log.Infof("before all executed")
	})

	Context("when the person is a child", func() {

		BeforeEach(func() {
			log.Infof("person is a child")
		})

		It("returns true", func() {
			person := ginkgo_functional_tests.Person{Age: 10}

			response := person.IsChild()

			Expect(response).To(BeTrue())
		})
	})

	Context("when the person is not a child", func() {
		BeforeEach(func() {
			log.Infof("person is not a child")
		})
		It("returns false", func() {
			person := ginkgo_functional_tests.Person{Age: 18}

			response := person.IsChild()

			Expect(response).To(BeFalse())
		})

		It("returns false", func() {
			person := ginkgo_functional_tests.Person{Age: 7}

			response := person.IsChild()

			Expect(response).To(BeTrue())
		})

		// if a test case is prefixed with `x` then this functional test case will be skipped
		It("returns false. this will be skipped", func() {
			person := ginkgo_functional_tests.Person{Age: 17}

			response := person.IsChild()

			Expect(response).To(BeTrue())
		})

		// if a test case is prefixed with `F` then only this functional test will run
		It("returns false. this will be skipped", func() {
			person := ginkgo_functional_tests.Person{Age: 37}

			response := person.IsChild()

			Expect(response).To(BeFalse())
		})
	})

	DescribeTable("isChild table test", func(age int, expectedResponse bool) {
		p := ginkgo_functional_tests.Person{Age: age}

		Expect(p.IsChild()).To(Equal(expectedResponse))
	},
		Entry("When person is a child", 10, true),
		Entry("When person is not a child", 22, false))
})
