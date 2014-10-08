package handlers

import (
	"bytes"
	"encoding/json"
	"ephemeris/core/representers"
	"ephemeris/testing/stubs"
	"log"
	"net/http"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rafaeljusto/go-testdb"
)

var _ = Describe("Events", func() {
	var (
		day, _ = time.ParseDuration("24h")
	)

	event := representers.EventRequest{
		Name:        "Test Name",
		Place:       "Test Place",
		Description: "Test Description",
		URL:         "http://test-url",
		LogoURL:     "http://test-logo-url",
		Beginning:   time.Now().Add(day),
		End:         time.Now().Add(2 * day),
		RegistrationBeginning: time.Now().Add(day),
		RegistrationEnd:       time.Now().Add(2 * day),
		Visibility:            "public",
		Status:                "open",
	}

	body, err := json.Marshal(event)
	if err != nil {
		log.Println("Unable to marshal event")
	}

	BeforeEach(func() {
		testdb.Reset()
	})

	Context("Getting events", func() {
		It("returns HTTP OK", func() {
			stubs.SelectUser()
			stubs.SelectAllEvents(stubs.ResultSuccess)

			Request("GET", "/api/events", false)
			Expect(response.Code).To(Equal(http.StatusOK))
		})

		It("returns HTTP No Content", func() {
			stubs.SelectAllEvents(stubs.ResultNoRows)

			Request("GET", "/api/events", false)
			Expect(response.Code).To(Equal(http.StatusNoContent))
		})

		It("returns HTTP Internal Server Error (error fetching event)", func() {
			stubs.SelectAllEvents(stubs.ResultError)

			Request("GET", "/api/events", false)
			Expect(response.Code).To(Equal(http.StatusInternalServerError))
		})

		It("returns HTTP Internal Server Error (error fetching event user)", func() {
			stubs.SelectAllEvents(stubs.ResultSuccess)
			stubs.SelectUserWithError()

			Request("GET", "/api/events", false)
			Expect(response.Code).To(Equal(http.StatusInternalServerError))
		})
	})

	Context("Creating events", func() {
		BeforeEach(func() {
			stubs.UpdateUser()
			stubs.SelectUser()
			stubs.SelectUserWithPassword()
		})

		It("returns a HTTP Created", func() {
			stubs.InsertEvent(stubs.ResultSuccess)

			Login(false)
			PostRequest("POST", "/api/events", bytes.NewReader(body), true)
			Expect(response.Code).To(Equal(http.StatusCreated))
		})

		It("returns a HTTP Internal Server Error", func() {
			stubs.InsertEvent(stubs.ResultError)

			Login(false)
			PostRequest("POST", "/api/events", bytes.NewReader(body), true)
			Expect(response.Code).To(Equal(http.StatusInternalServerError))
		})
	})

	Context("Getting an event", func() {
		It("returns HTTP OK", func() {
			stubs.SelectUser()
			stubs.SelectEvent(stubs.ResultSuccess)

			Request("GET", "/api/events/1", false)
			Expect(response.Code).To(Equal(http.StatusOK))
		})

		It("returns HTTP Not Found", func() {
			stubs.SelectEvent(stubs.ResultNoRows)

			Request("GET", "/api/events/1", false)
			Expect(response.Code).To(Equal(http.StatusNotFound))
		})

		It("returns HTTP Internal Server Error (error fetching event)", func() {
			stubs.SelectEvent(stubs.ResultError)

			Request("GET", "/api/events/1", false)
			Expect(response.Code).To(Equal(http.StatusInternalServerError))
		})

		It("returns HTTP Internal Server Error (error fetching event user)", func() {
			stubs.SelectEvent(stubs.ResultSuccess)
			stubs.SelectUserWithError()

			Request("GET", "/api/events/1", false)
			Expect(response.Code).To(Equal(http.StatusInternalServerError))
		})
	})

	Context("Updating an event", func() {
		BeforeEach(func() {
			stubs.SelectUser()
			stubs.SelectUserWithPassword()
		})

		It("returns HTTP OK", func() {
			stubs.SelectEvent(stubs.ResultSuccess)
			stubs.UpdateEvent(stubs.ResultSuccess)

			Login(false)
			PostRequest("PUT", "/api/events/1", bytes.NewReader(body), true)
			Expect(response.Code).To(Equal(http.StatusOK))
		})

		It("returns HTTP Not Found", func() {
			stubs.SelectEvent(stubs.ResultNoRows)

			Login(false)
			PostRequest("PUT", "/api/events/1", bytes.NewReader(body), true)
			Expect(response.Code).To(Equal(http.StatusNotFound))
		})

		It("returns HTTP Internal Server Error (checking if object exists)", func() {
			stubs.SelectEvent(stubs.ResultError)

			Login(false)
			PostRequest("PUT", "/api/events/1", bytes.NewReader(body), true)
			Expect(response.Code).To(Equal(http.StatusInternalServerError))
		})

		It("returns HTTP Internal Server Error (updating object)", func() {
			stubs.SelectEvent(stubs.ResultSuccess)
			stubs.UpdateEvent(stubs.ResultError)

			Login(false)
			PostRequest("PUT", "/api/events/1", bytes.NewReader(body), true)
			Expect(response.Code).To(Equal(http.StatusInternalServerError))
		})
	})

	Context("Deleting an event", func() {
		BeforeEach(func() {
			stubs.SelectUser()
			stubs.SelectUserWithPassword()
		})

		It("returns HTTP No Content", func() {
			stubs.SelectEvent(stubs.ResultSuccess)
			stubs.DeleteEvent(stubs.ResultSuccess)

			Login(false)
			Request("DELETE", "/api/events/1", true)
			Expect(response.Code).To(Equal(http.StatusNoContent))
		})

		It("returns HTTP Not Found", func() {
			stubs.SelectEvent(stubs.ResultNoRows)

			Login(false)
			Request("DELETE", "/api/events/1", true)
			Expect(response.Code).To(Equal(http.StatusNotFound))
		})

		It("returns HTTP Internal Server Error (checking if object exists)", func() {
			stubs.SelectEvent(stubs.ResultError)

			Login(false)
			Request("DELETE", "/api/events/1", true)
			Expect(response.Code).To(Equal(http.StatusInternalServerError))
		})

		It("returns HTTP Internal Server Error (deleting object)", func() {
			stubs.SelectEvent(stubs.ResultSuccess)
			stubs.DeleteEvent(stubs.ResultError)

			Login(false)
			Request("DELETE", "/api/events/1", true)
			Expect(response.Code).To(Equal(http.StatusInternalServerError))
		})
	})
})
