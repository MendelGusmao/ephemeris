package stubs

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/rafaeljusto/go-testdb"
)

var (
	sqlSelectUser                = "SELECT * FROM `users` WHERE (`id` = ?) LIMIT 1"
	sqlSelectUserWithPassword    = "SELECT * FROM `users` WHERE (`password` = ?) AND (`username` = ?) ORDER BY `users`.id ASC LIMIT 1"
	sqlSelectUserWithPasswordAlt = "SELECT * FROM `users` WHERE (`username` = ?) AND (`password` = ?) ORDER BY `users`.id ASC LIMIT 1"
	sqlUpdateUser                = "UPDATE `users` SET `administrator` = ?, `created_at` = ?, `password` = ?, `updated_at` = ?, `username` = ?  WHERE (`id` = ?)"

	userFields = []string{
		"id",
		"username",
		"password",
		"administrator",
		"created_at",
		"updated_at",
	}
	userData = []string{
		"1",
		"ephemeris",
		"ephemeris",
		"f",
		"2014-08-25 18:55:49.111865-03",
		"2014-08-25 19:01:05.492411-03",
	}
)

func SelectUser() {
	testdb.StubQuery(sqlSelectUser, testdb.RowsFromCSVString(userFields, strings.Join(userData, ",")))
}

func SelectUserWithError() {
	testdb.StubQueryError(sqlSelectUser, errors.New("Forged error: SelectUserWithError."))
}

func SelectNonExistentUser() {
	testdb.StubQueryError(sqlSelectUser, sql.ErrNoRows)
}

func SelectUserWithPassword() {
	rows := testdb.RowsFromCSVString(userFields, strings.Join(userData, ","))
	testdb.StubQuery(sqlSelectUserWithPassword, rows)
	testdb.StubQuery(sqlSelectUserWithPasswordAlt, rows)
}

func SelectUserWithWrongPassword() {
	testdb.StubQueryError(sqlSelectUserWithPassword, sql.ErrNoRows)
	testdb.StubQueryError(sqlSelectUserWithPasswordAlt, sql.ErrNoRows)
}

func SelectUserWithPasswordAndError() {
	err := errors.New("Forged error: SelectUserWithPasswordAndError.")
	testdb.StubQueryError(sqlSelectUserWithPassword, err)
	testdb.StubQueryError(sqlSelectUserWithPasswordAlt, err)
}

func UpdateUser() {
	testdb.StubExec(sqlUpdateUser, testdb.NewResult(1, nil, 1, nil))
}
