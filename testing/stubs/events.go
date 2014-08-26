package stubs

import (
	"github.com/erikstmartin/go-testdb"
)

const (
	sqlEventInsert = `
		INSERT INTO events ()
		VALUES ()
	`
)

func EventCreationOK() (string, *testdb.Result) {
	return "", testdb.NewResult(1, nil, 1, nil)
}
