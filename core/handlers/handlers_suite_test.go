package handlers

import (
	"bytes"
	"encoding/json"
	"ephemeris/core/config"
	"ephemeris/core/models"
	"ephemeris/core/server"
	"ephemeris/testing/fake"
	"fmt"
	"io"
	"log"
	"log/syslog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/MendelGusmao/go-testdb"
	"github.com/go-martini/martini"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	response   *httptest.ResponseRecorder
	testConfig = config.EphemerisConfig{
		APIRoot:  "/api",
		Log:      config.LogConfig{syslog.LOG_DEBUG, config.SyslogConfig{}},
		Database: config.DatabaseConfig{"testdb", "", 0},
		Session:  config.SessionConfig{"test", [][]byte{{65, 66, 67, 1, 2, 3}}, config.RedisConfig{}},
	}
	m      *martini.ClassicMartini
	cookie string

	sessionURI         = "/api/session"
	postgresDateFormat = "2006-01-02 15:04:05.000000-07"
)

func TestMain(t *testing.T) {
	var err error
	m, err = server.Setup(testConfig)

	if err != nil {
		fmt.Println("Error configuring server:", err)
		t.Fatal(err)
	}

	testdb.EnableTimeParsingWithFormat(postgresDateFormat)
	RegisterFailHandler(Fail)
	RunSpecs(t, "Handlers Suite")
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
	body, err := json.Marshal(models.UserRequest{
		Username: "ephemeris",
		Password: fake.String("ephemeris"),
	})

	if err != nil {
		log.Println("Unable to marshal user")
	}

	PostRequest("POST", sessionURI, bytes.NewReader(body), useCookie)
}
