package handlers

import (
	"ephemeris/core/models"
	"ephemeris/testing/stubs"
	"net/http"

	"github.com/MendelGusmao/go-testdb"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Session Handlers", func() {
	var (
		uri = sessionURI
	)

	BeforeEach(func() {
		testdb.Reset()
		cookie = ""
	})

	Context("Checking", func() {
		BeforeEach(func() {
			stubs.SelectUser(stubs.ResultNoRows, nil)
			stubs.SelectUserWithPassword(stubs.ResultSuccess, stubs.UserData(models.UserRoleNone))
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
			stubs.SelectUser(stubs.ResultNoRows, nil)
			stubs.SelectUserWithPassword(stubs.ResultNoRows, nil)

			Login(false)
			Expect(response.Code).To(Equal(http.StatusNotFound))
		})

		It("returns HTTP Created", func() {
			stubs.SelectUser(stubs.ResultNoRows, nil)
			stubs.SelectUserWithPassword(stubs.ResultSuccess, stubs.UserData(models.UserRoleNone))

			Login(false)
			Expect(response.Code).To(Equal(http.StatusCreated))
		})

		It("returns HTTP No Content", func() {
			stubs.SelectUser(stubs.ResultNoRows, nil)
			stubs.SelectUserWithPassword(stubs.ResultSuccess, stubs.UserData(models.UserRoleNone))

			Login(false)
			Login(true)
			Expect(response.Code).To(Equal(http.StatusNoContent))
		})

		It("returns HTTP Internal Server Error", func() {
			stubs.SelectUserWithPassword(stubs.ResultError, nil)

			Login(false)
			Expect(response.Code).To(Equal(http.StatusInternalServerError))
		})
	})

	Context("Logging out", func() {
		It("returns HTTP No Content", func() {
			stubs.SelectUser(stubs.ResultSuccess, stubs.UserData(models.UserRoleNone))
			stubs.SelectUserWithPassword(stubs.ResultSuccess, stubs.UserData(models.UserRoleNone))

			Login(false)
			Request("DELETE", uri, true)
			Expect(response.Code).To(Equal(http.StatusNoContent))
		})
	})
})
