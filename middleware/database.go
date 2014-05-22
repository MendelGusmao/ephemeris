package middleware

import (
	"events/lib/gorm"
	"events/lib/martini"
	_ "events/lib/pq"
)

type DatabaseOptions struct {
	URL string
}

func Database(databaseOptions DatabaseOptions) martini.Handler {
	return func(c martini.Context) {
		db, err := gorm.Open("postgres", databaseOptions.URL)

		if err != nil {
			panic(err)
		}

		c.Map(db)
		defer db.Close()
		c.Next()
	}
}
