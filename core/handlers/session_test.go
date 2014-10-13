package handlers

import (
	"ephemeris/testing/stubs"
	"net/http"

	"github.com/MendelGusmao/go-testdb"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Session", func() {
	var (
		uri = sessionURI
	)

	BeforeEach(func() {
		testdb.Reset()
		cookie = ""
	})

	Context("Checking", func() {
		BeforeEach(func() {
			stubs.SelectUser(stubs.ResultNoRows)
			stubs.SelectUserWithPassword(stubs.ResultSuccess)
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

	Context("Logging in", func() {
		It("returns HTTP Not Found", func() {
			stubs.SelectUser(stubs.ResultNoRows)
			stubs.SelectUserWithPassword(stubs.ResultNoRows)

			Login(false)
			Expect(response.Code).To(Equal(http.StatusNotFound))
		})

		It("returns HTTP Created", func() {
			stubs.SelectUser(stubs.ResultNoRows)
			stubs.SelectUserWithPassword(stubs.ResultSuccess)

			Login(false)
			Expect(response.Code).To(Equal(http.StatusCreated))
		})

		It("returns HTTP No Content", func() {
			stubs.SelectUser(stubs.ResultNoRows)
			stubs.SelectUserWithPassword(stubs.ResultSuccess)

			Login(false)
			Login(true)
			Expect(response.Code).To(Equal(http.StatusNoContent))
		})

		It("returns HTTP Internal Server Error", func() {
			stubs.SelectUserWithPassword(stubs.ResultError)

			Login(false)
			Expect(response.Code).To(Equal(http.StatusInternalServerError))
		})
	})

	Context("Logging out", func() {
		It("returns HTTP No Content", func() {
			stubs.SelectUser(stubs.ResultSuccess)
			stubs.SelectUserWithPassword(stubs.ResultSuccess)

			Login(false)
			Request("DELETE", uri, true)
			Expect(response.Code).To(Equal(http.StatusNoContent))
		})
	})
})
