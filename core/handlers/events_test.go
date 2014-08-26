package handlers

import (
	"ephemeris/core/middleware"
	"ephemeris/core/models"
	"ephemeris/core/representers"
	"ephemeris/testing/fake"
	"ephemeris/testing/stubs"
	// "fmt"
	"github.com/erikstmartin/go-testdb"
	"github.com/jinzhu/gorm"
	"github.com/martini-contrib/render"
	"testing"
	"time"
)

var (
	day, _ = time.ParseDuration("24h")
)

func TestCreateEvent(t *testing.T) {
	db, _ := gorm.Open("testdb", "")

	tests := []struct {
		Database *gorm.DB
		Request  representers.EventRequest
		Logger   *middleware.ApplicationLogger
		Renderer render.Render
		User     *models.User
	}{
		{
			Database: &db,
			Request: representers.EventRequest{
				Name:        "ephemeris",
				Place:       "ephemeris",
				Description: "ephemeris",
				URL:         "ephemeris",
				LogoURL:     "ephemeris",
				Beginning:   time.Now().Add(day),
				End:         time.Now().Add(3 * day),
				RegistrationBeginning: time.Now().Add(day),
				RegistrationEnd:       time.Now().Add(2 * day),
				Visibility:            string(models.EventVisibilityPublic),
				Status:                string(models.EventStatusOpen),
			},
			Logger: fake.Logger(),
			User: &models.User{
				Id:       1,
				Username: "ephemeris",
			},
			Renderer: fake.Renderer(),
		},
	}

	for _, test := range tests {
		query, result := stubs.EventCreationOK()
		testdb.StubExec(query, result)

		createEvent(test.Database, test.Request, test.Logger, test.Renderer, test.User)
	}
}
