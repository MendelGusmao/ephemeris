package handlers

import (
	"ephemeris/testing/stubs"
	"net/http"

	"github.com/erikstmartin/go-testdb"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Session", func() {
	var (
		uri = sessionURI
	)

	BeforeEach(func() {
		testdb.Reset()
	})

	Context("GET", func() {
		BeforeEach(func() {
			stubs.SelectNonExistentUser()
			stubs.SelectUserWithPassword()
		})

		It("returns HTTP Forbidden", func() {
			Request("GET", uri, false)
			Expect(response.Code).To(Equal(http.StatusForbidden))
		})

		It("returns HTTP No Content", func() {
			Login(false)
			Request("GET", uri, true)
			Expect(response.Code).To(Equal(http.StatusNoContent))
		})
	})

	Context("POST", func() {
		It("returns HTTP Not Found", func() {
			stubs.SelectNonExistentUser()
			stubs.SelectUserWithWrongPassword()

			Login(false)
			Expect(response.Code).To(Equal(http.StatusNotFound))
		})

		It("returns HTTP Created", func() {
			stubs.SelectNonExistentUser()
			stubs.SelectUserWithPassword()

			Login(false)
			Expect(response.Code).To(Equal(http.StatusCreated))
		})

		It("returns HTTP No Content", func() {
			stubs.SelectNonExistentUser()
			stubs.SelectUserWithPassword()

			Login(false)
			Login(true)
			Expect(response.Code).To(Equal(http.StatusNoContent))
		})

		It("returns HTTP Internal Server Error", func() {
			stubs.SelectUserWithError()

			Login(false)
			Expect(response.Code).To(Equal(http.StatusInternalServerError))
		})
	})

	Context("DELETE", func() {
		It("returns HTTP No Content", func() {
			stubs.SelectUser()
			stubs.SelectUserWithPassword()

			Login(false)
			Request("DELETE", uri, true)
			Expect(response.Code).To(Equal(http.StatusNoContent))
		})
	})
})
