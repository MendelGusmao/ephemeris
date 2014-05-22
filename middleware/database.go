package middleware

import (
	"events/lib/martini"
	"labix.org/v2/mgo"
)

type DatabaseOptions struct {
	URL       string
	Name      string
	Monotonic bool
}

func Database(databaseOptions DatabaseOptions) martini.Handler {
	session, err := mgo.Dial(databaseOptions.URL)

	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, databaseOptions.Monotonic)

	return func(c martini.Context) {
		s := session.Clone()
		c.Map(s.DB(databaseOptions.Name))
		defer s.Close()
		c.Next()
	}
}
