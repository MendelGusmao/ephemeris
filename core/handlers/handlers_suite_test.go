package handlers

import (
	"bytes"
	"encoding/json"
	"ephemeris/core/config"
	"ephemeris/core/representers"
	"ephemeris/core/server"
	"ephemeris/testing/fake"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-martini/martini"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	response   *httptest.ResponseRecorder
	testConfig = config.EphemerisConfig{
		APIRoot:  "/api",
		Database: config.DatabaseConfig{"testdb", "", 0},
		Session:  config.SessionConfig{"test", "session"},
	}
	m      *martini.ClassicMartini
	cookie string

	sessionURI = "/api/session"
)

func TestMain(t *testing.T) {
	m = Setup()
	RegisterFailHandler(Fail)
	RunSpecs(t, "Main Suite")
}

func Setup() *martini.ClassicMartini {
	m := martini.Classic()
	server.Configure(testConfig, m)
	return m
}

func Request(method, route string, useCookie bool) {
	response = httptest.NewRecorder()
	request, _ := http.NewRequest(method, route, nil)

	if useCookie && len(cookie) > 0 {
		request.Header.Add("Cookie", cookie)
	}

	m.ServeHTTP(response, request)

	if setCookie := response.Header().Get("Set-Cookie"); len(setCookie) > 0 {
		cookie = setCookie
	}
}

func PostRequest(method, route string, body io.Reader, useCookie bool) {
	response = httptest.NewRecorder()
	request, _ := http.NewRequest(method, route, body)
	request.Header.Add("Content-Type", "application/json")

	if useCookie {
		request.Header.Add("Cookie", cookie)
	}

	m.ServeHTTP(response, request)

	if setCookie := response.Header().Get("Set-Cookie"); len(setCookie) > 0 {
		cookie = setCookie
	}
}

func Login(useCookie bool) {
	body, err := json.Marshal(representers.UserRequest{
		Username: "ephemeris",
		Password: fake.String("ephemeris"),
	})

	if err != nil {
		log.Println("Unable to marshal user")
	}

	PostRequest("POST", sessionURI, bytes.NewReader(body), useCookie)
}
