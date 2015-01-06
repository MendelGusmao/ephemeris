package handlers

import (
	"bytes"
	"encoding/json"
	"ephemeris/core/models"
	"ephemeris/testing/fake"
	"ephemeris/testing/stubs"
	"log"
	"net/http"

	"github.com/MendelGusmao/go-testdb"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Users", func() {
	var (
		usersURI = "/api/users"
		userURI  = "/api/users/1"
	)

	user := models.UserRequest{
		Username: "Test User",
		Password: fake.String("Test Password"),
		Role:     int(models.UserRoleRegular),
	}

	body, err := json.Marshal(user)
	if err != nil {
		log.Println("Unable to marshal user")
	}

	BeforeEach(func() {
		testdb.Reset()
		cookie = ""
		stubs.SelectUserWithPassword(stubs.ResultSuccess, stubs.UserData(models.UserRoleAdministrator))
		Login(false)
	})

	Context("Getting users", func() {
		BeforeEach(func() {
			stubs.SelectUser(stubs.ResultSuccess, stubs.UserData(models.UserRoleAdministrator))
		})

		It("returns HTTP OK", func() {
			stubs.SelectAllUsers(stubs.ResultSuccess, stubs.UserData(models.UserRoleRegular))

			Request("GET", usersURI, true)
			Expect(response.Code).To(Equal(http.StatusOK))
		})

		It("returns HTTP No Content", func() {
			stubs.SelectAllUsers(stubs.ResultNoRows, stubs.UserData(models.UserRoleRegular))

			Request("GET", usersURI, true)
			Expect(response.Code).To(Equal(http.StatusNoContent))
		})

		It("returns HTTP Internal Server Error (error fetching user)", func() {
			stubs.SelectAllUsers(stubs.ResultError, stubs.UserData(models.UserRoleRegular))

			Request("GET", usersURI, true)
			Expect(response.Code).To(Equal(http.StatusInternalServerError))
		})
	})

	Context("Creating users", func() {
		BeforeEach(func() {
			stubs.SelectUser(stubs.ResultSuccess, stubs.UserData(models.UserRoleAdministrator))
		})

		It("returns a HTTP Created", func() {
			stubs.InsertUser(stubs.ResultSuccess)

			PostRequest("POST", usersURI, bytes.NewReader(body), true)
			Expect(response.Code).To(Equal(http.StatusCreated))
		})

		It("returns a HTTP Internal Server Error", func() {
			stubs.InsertUser(stubs.ResultError)

			PostRequest("POST", usersURI, bytes.NewReader(body), true)
			Expect(response.Code).To(Equal(http.StatusInternalServerError))
		})
	})

	Context("Getting an user", func() {
		BeforeEach(func() {
			stubs.SelectUser(stubs.ResultSuccess, stubs.UserData(models.UserRoleAdministrator))
		})

		It("returns HTTP OK", func() {
			stubs.SelectUser(stubs.ResultSuccess, stubs.UserData(models.UserRoleRegular))

			Request("GET", userURI, true)
			Expect(response.Code).To(Equal(http.StatusOK))
		})

		It("returns HTTP Not Found", func() {
			stubs.SelectUser(stubs.ResultNoRows, nil)

			Request("GET", userURI, true)
			Expect(response.Code).To(Equal(http.StatusNotFound))
		})

		It("returns HTTP Internal Server Error (error fetching user)", func() {
			stubs.SelectUser(stubs.ResultError, nil)

			Request("GET", userURI, true)
			Expect(response.Code).To(Equal(http.StatusInternalServerError))
		})
	})

	Context("Updating an user", func() {
		BeforeEach(func() {
			stubs.SelectUser(stubs.ResultSuccess, stubs.UserData(models.UserRoleAdministrator))
		})

		It("returns HTTP No Content", func() {
			stubs.SelectUser(stubs.ResultSuccess, stubs.UserData(models.UserRoleRegular))
			stubs.UpdateUser(stubs.ResultSuccess)

			PostRequest("PUT", userURI, bytes.NewReader(body), true)
			Expect(response.Code).To(Equal(http.StatusNoContent))
		})

		It("returns HTTP Not Found", func() {
			stubs.SelectUser(stubs.ResultNoRows, nil)

			PostRequest("PUT", userURI, bytes.NewReader(body), true)
			Expect(response.Code).To(Equal(http.StatusNotFound))
		})

		It("returns HTTP Internal Server Error (checking if object exists)", func() {
			stubs.SelectUser(stubs.ResultError, nil)

			PostRequest("PUT", userURI, bytes.NewReader(body), true)
			Expect(response.Code).To(Equal(http.StatusInternalServerError))
		})

		It("returns HTTP Internal Server Error (updating object)", func() {
			stubs.SelectUser(stubs.ResultSuccess, stubs.UserData(models.UserRoleRegular))
			stubs.UpdateUser(stubs.ResultError)

			PostRequest("PUT", userURI, bytes.NewReader(body), true)
			Expect(response.Code).To(Equal(http.StatusInternalServerError))
		})
	})

	Context("Deleting an user", func() {
		BeforeEach(func() {
			stubs.SelectUser(stubs.ResultSuccess, stubs.UserData(models.UserRoleAdministrator))
		})

		It("returns HTTP No Content", func() {
			stubs.SelectUser(stubs.ResultSuccess, stubs.UserData(models.UserRoleRegular))
			stubs.DeleteUser(stubs.ResultSuccess)

			Request("DELETE", userURI, true)
			Expect(response.Code).To(Equal(http.StatusNoContent))
		})

		It("returns HTTP Not Found", func() {
			stubs.SelectUser(stubs.ResultNoRows, nil)

			Request("DELETE", userURI, true)
			Expect(response.Code).To(Equal(http.StatusNotFound))
		})

		It("returns HTTP Internal Server Error (checking if object exists)", func() {
			stubs.SelectUser(stubs.ResultError, nil)

			Request("DELETE", userURI, true)
			Expect(response.Code).To(Equal(http.StatusInternalServerError))
		})

		It("returns HTTP Internal Server Error (deleting object)", func() {
			stubs.SelectUser(stubs.ResultSuccess, stubs.UserData(models.UserRoleRegular))
			stubs.DeleteUser(stubs.ResultError)

			Request("DELETE", userURI, true)
			Expect(response.Code).To(Equal(http.StatusInternalServerError))
		})
	})
})
