package middleware

import (
	"github.com/MendelGusmao/gorm"
	"github.com/go-martini/martini"
	_ "github.com/lib/pq"
)

type DBOptions struct {
	Driver             string
	URL                string
	MaxIdleConnections int
}

func Database(dbOptions DBOptions) martini.Handler {
	return func(c martini.Context) {
		db, err := gorm.Open(dbOptions.Driver, dbOptions.URL)

		if err != nil {
			panic(err)
		}

		db.DB().SetMaxIdleConns(dbOptions.MaxIdleConnections)

		c.Map(&db)
		defer db.Close()
		c.Next()
	}
}
