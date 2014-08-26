package fake

import (
	"ephemeris/core/middleware"
	"github.com/jinzhu/gorm"
	"github.com/martini-contrib/render"
	"log"
	"os"
)

func DB() *gorm.DB {
	db, _ := gorm.Open("testdb", "")
	return &db
}

func Logger() *middleware.ApplicationLogger {
	return &middleware.ApplicationLogger{
		Logger: log.New(os.Stdout, "[ephemeris-test] ", 0),
	}
}

func Renderer() render.Render {
	return renderer{}
}
