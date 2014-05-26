package middleware

import (
	"ephemeris/lib/gorm"
	"ephemeris/lib/martini"
	_ "ephemeris/lib/pq"
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

		c.Map(&db)
		defer db.Close()
		c.Next()
	}
}
