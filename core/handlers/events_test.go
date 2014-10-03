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

var _ = Describe("/api/events", func() {
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

	Context("GET", func() {
		BeforeEach(func() {
			stubs.SelectUser()
		})

		It("returns HTTP OK", func() {
			stubs.SelectAllEvents(true)

			Request("GET", "/api/events", false)
			Expect(response.Code).To(Equal(http.StatusOK))
		})

		It("returns HTTP No Content", func() {
			stubs.SelectAllEvents(false)

			Request("GET", "/api/events", false)
			Expect(response.Code).To(Equal(http.StatusNoContent))
		})

		It("returns HTTP Internal Server Error", func() {
			stubs.SelectAllEventsWithError()

			Request("GET", "/api/events", false)
			Expect(response.Code).To(Equal(http.StatusInternalServerError))
		})
	})

	Context("POST", func() {
		BeforeEach(func() {
			stubs.UpdateUser()
			stubs.SelectUser()
			stubs.SelectUserWithPassword()
		})

		It("returns a HTTP Forbidden", func() {
			testdb.Reset()

			PostRequest("POST", "/api/events", bytes.NewReader(body), false)
			Expect(response.Code).To(Equal(http.StatusForbidden))
		})

		It("returns a HTTP Created", func() {
			stubs.InsertEvent()

			Login(false)
			PostRequest("POST", "/api/events", bytes.NewReader(body), true)
			Expect(response.Code).To(Equal(http.StatusCreated))
		})

		It("returns a HTTP Internal Server Error", func() {
			stubs.InsertEventWithError()

			Login(false)
			PostRequest("POST", "/api/events", bytes.NewReader(body), true)
			Expect(response.Code).To(Equal(http.StatusInternalServerError))
		})
	})
})
