package models

import (
	"github.com/martini-contrib/binding"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Event Model", func() {
	BeforeEach(func() {
	})

	Context("Validating an event", func() {
		It("returns an Invalid URL error", func() {
			event := EventRequest{}
			expectedError := binding.Errors{{
				FieldNames:     []string{"url"},
				Classification: "URLError",
				Message:        "Invalid URL",
			}}

			event.validateURL("%gh&%ij", "url")

			Expect(event.errors).To(Equal(expectedError))
		})

		It("returns an Invalid URL Scheme error", func() {
			event := EventRequest{
				URL:     "foo://bar",
				LogoURL: "http://bar",
			}
			expectedError := binding.Errors{{
				FieldNames:     []string{"URL"},
				Classification: "URLError",
				Message:        "Invalid URL Scheme",
			}}

			event.validateURLs()

			Expect(event.errors).To(Equal(expectedError))
		})

		It("returns an Invalid Visibility error", func() {
			event := EventRequest{
				Status: "open",
			}
			expectedError := binding.Errors{{
				FieldNames:     []string{"visibility"},
				Classification: "EnumError",
				Message:        "Invalid visibility ''",
			}}

			event.validateEnums()

			Expect(event.errors).To(Equal(expectedError))
		})

		It("returns an Invalid Status error", func() {
			event := EventRequest{
				Visibility: "public",
			}
			expectedError := binding.Errors{{
				FieldNames:     []string{"status"},
				Classification: "EnumError",
				Message:        "Invalid status ''",
			}}

			event.validateEnums()

			Expect(event.errors).To(Equal(expectedError))
		})
	})
})
