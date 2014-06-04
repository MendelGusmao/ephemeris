package middleware

import (
	"github.com/go-martini/martini"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type DatabaseOptions struct {
	URL                string
	MaxIdleConnections int
}

func Database(databaseOptions DatabaseOptions) martini.Handler {
	return func(c martini.Context) {
		db, err := gorm.Open("postgres", databaseOptions.URL)

		if err != nil {
			panic(err)
		}

		db.DB().SetMaxIdleConns(databaseOptions.MaxIdleConnections)

		c.Map(&db)
		defer db.Close()
		c.Next()
	}
}
