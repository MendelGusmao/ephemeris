package handlers

import (
	"bytes"
	"encoding/json"
	"ephemeris/core/representers"
	"ephemeris/testing/fake"
	"ephemeris/testing/stubs"
	"log"
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rafaeljusto/go-testdb"
)

var _ = Describe("Users", func() {
	user := representers.UserRequest{
		Username:      "Test User",
		Password:      fake.String("Test Password"),
		Administrator: false,
	}

	body, err := json.Marshal(user)
	if err != nil {
		log.Println("Unable to marshal user")
	}

	BeforeEach(func() {
		testdb.Reset()
	})

	Context("Getting users", func() {
		It("returns HTTP OK", func() {
			stubs.SelectAllUsers(stubs.ResultSuccess)

			Request("GET", "/api/users", false)
			Expect(response.Code).To(Equal(http.StatusOK))
		})

		It("returns HTTP No Content", func() {
			stubs.SelectAllUsers(stubs.ResultNoRows)

			Request("GET", "/api/users", false)
			Expect(response.Code).To(Equal(http.StatusNoContent))
		})

		It("returns HTTP Internal Server Error (error fetching user)", func() {
			stubs.SelectAllUsers(stubs.ResultError)

			Request("GET", "/api/users", false)
			Expect(response.Code).To(Equal(http.StatusInternalServerError))
		})
	})

	Context("Creating users", func() {
		BeforeEach(func() {
			stubs.SelectUserWithPassword()
		})

		It("returns a HTTP Created", func() {
			stubs.InsertUser(stubs.ResultSuccess)

			Login(false)
			PostRequest("POST", "/api/users", bytes.NewReader(body), true)
			Expect(response.Code).To(Equal(http.StatusCreated))
		})

		It("returns a HTTP Internal Server Error", func() {
			stubs.InsertUser(stubs.ResultError)

			Login(false)
			PostRequest("POST", "/api/users", bytes.NewReader(body), true)
			Expect(response.Code).To(Equal(http.StatusInternalServerError))
		})
	})

	Context("Getting an user", func() {
		It("returns HTTP OK", func() {
			stubs.SelectUser(stubs.ResultSuccess)

			Request("GET", "/api/users/1", false)
			Expect(response.Code).To(Equal(http.StatusOK))
		})

		It("returns HTTP Not Found", func() {
			stubs.SelectUser(stubs.ResultNoRows)

			Request("GET", "/api/users/1", false)
			Expect(response.Code).To(Equal(http.StatusNotFound))
		})

		It("returns HTTP Internal Server Error (error fetching user)", func() {
			stubs.SelectUser(stubs.ResultError)

			Request("GET", "/api/users/1", false)
			Expect(response.Code).To(Equal(http.StatusInternalServerError))
		})
	})

	Context("Updating an user", func() {
		BeforeEach(func() {
			stubs.SelectUser(stubs.ResultSuccess)
			stubs.SelectUserWithPassword()
		})

		It("returns HTTP OK", func() {
			stubs.UpdateUser(stubs.ResultSuccess)

			Login(false)
			PostRequest("PUT", "/api/users/1", bytes.NewReader(body), true)
			Expect(response.Code).To(Equal(http.StatusOK))
		})

		It("returns HTTP Not Found", func() {
			stubs.SelectUser(stubs.ResultNoRows)

			Login(false)
			PostRequest("PUT", "/api/users/1", bytes.NewReader(body), true)
			Expect(response.Code).To(Equal(http.StatusNotFound))
		})

		It("returns HTTP Internal Server Error (checking if object exists)", func() {
			stubs.SelectUser(stubs.ResultError)

			Login(false)
			PostRequest("PUT", "/api/users/1", bytes.NewReader(body), true)
			Expect(response.Code).To(Equal(http.StatusInternalServerError))
		})

		It("returns HTTP Internal Server Error (updating object)", func() {
			stubs.SelectUser(stubs.ResultSuccess)
			stubs.UpdateUser(stubs.ResultError)

			Login(false)
			PostRequest("PUT", "/api/users/1", bytes.NewReader(body), true)
			Expect(response.Code).To(Equal(http.StatusInternalServerError))
		})
	})

	Context("Deleting an user", func() {
		BeforeEach(func() {
			stubs.SelectUserWithPassword()
		})

		It("returns HTTP No Content", func() {
			stubs.SelectUser(stubs.ResultSuccess)
			stubs.DeleteUser(stubs.ResultSuccess)

			Login(false)
			Request("DELETE", "/api/users/1", true)
			Expect(response.Code).To(Equal(http.StatusNoContent))
		})

		It("returns HTTP Not Found", func() {
			stubs.SelectUser(stubs.ResultNoRows)

			Login(false)
			Request("DELETE", "/api/users/1", true)
			Expect(response.Code).To(Equal(http.StatusNotFound))
		})

		It("returns HTTP Internal Server Error (checking if object exists)", func() {
			stubs.SelectUser(stubs.ResultError)

			Login(false)
			Request("DELETE", "/api/users/1", true)
			Expect(response.Code).To(Equal(http.StatusInternalServerError))
		})

		It("returns HTTP Internal Server Error (deleting object)", func() {
			stubs.SelectUser(stubs.ResultSuccess)
			stubs.DeleteUser(stubs.ResultError)

			Login(false)
			Request("DELETE", "/api/users/1", true)
			Expect(response.Code).To(Equal(http.StatusInternalServerError))
		})
	})
})
