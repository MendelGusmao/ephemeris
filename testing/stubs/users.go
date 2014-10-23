package stubs

import (
	"database/sql"
	"ephemeris/core/models"
	"errors"
	"strconv"
	"strings"

	"github.com/MendelGusmao/go-testdb"
)

var (
	sqlSelectAllUsers            = "SELECT * FROM `users`"
	sqlSelectUser                = "SELECT * FROM `users` WHERE (`id` = ?) LIMIT 1"
	sqlSelectUserWithPassword    = "SELECT * FROM `users` WHERE (`password` = ?) AND (`username` = ?) LIMIT 1"
	sqlSelectUserWithPasswordAlt = "SELECT * FROM `users` WHERE (`username` = ?) AND (`password` = ?) LIMIT 1"
	sqlUpdateUser                = "UPDATE `users` SET `created_at` = ?, `password` = ?, `role` = ?, `updated_at` = ?, `username` = ?  WHERE (`id` = ?)"
	sqlInsertUser                = "INSERT INTO `users` (`created_at`,`password`,`role`,`updated_at`,`username`) VALUES (?,?,?,?,?)"
	sqlDeleteUser                = "DELETE FROM `users`  WHERE (`id` = ?)"

	userFields = []string{
		"id",
		"username",
		"password",
		"role",
		"created_at",
		"updated_at",
	}
	userData = []string{
		"1",
		"ephemeris",
		"ephemeris",
		"0",
		"2014-08-25 18:55:49.111865-03",
		"2014-08-25 19:01:05.492411-03",
	}
)

func UserData(role models.UserRole) []string {
	newUserData := make([]string, len(userData))
	copy(newUserData, userData)
	newUserData[3] = strconv.Itoa(int(role))
	return newUserData
}

func SelectAllUsers(result Result, userData []string) {
	switch result {
	case ResultSuccess:
		if userData == nil {
			panic("userData is nil")
		}

		testdb.StubQuery(sqlSelectAllUsers, testdb.RowsFromCSVString(userFields, strings.Join(userData, ",")))
	case ResultNoRows:
		testdb.StubQueryError(sqlSelectAllUsers, sql.ErrNoRows)
	case ResultError:
		testdb.StubQueryError(sqlSelectAllUsers, errors.New("Forged error: SelectAllUsers."))
	}
}

func SelectUser(result Result, userData []string) {
	switch result {
	case ResultSuccess:
		rows := testdb.RowsFromCSVString(userFields, strings.Join(userData, ","))
		testdb.StubQuery(sqlSelectUser, rows)
	case ResultNoRows:
		testdb.StubQueryError(sqlSelectUser, sql.ErrNoRows)
	case ResultError:
		testdb.StubQueryError(sqlSelectUser, errors.New("Forged error: SelectUser."))
	}
}

func SelectUserWithPassword(result Result, userData []string) {
	switch result {
	case ResultSuccess:
		rows := testdb.RowsFromCSVString(userFields, strings.Join(userData, ","))
		testdb.StubQuery(sqlSelectUserWithPassword, rows)
		testdb.StubQuery(sqlSelectUserWithPasswordAlt, rows)
	case ResultNoRows:
		testdb.StubQuery(sqlSelectUserWithPassword, sql.ErrNoRows)
		testdb.StubQuery(sqlSelectUserWithPasswordAlt, sql.ErrNoRows)
	case ResultError:
		testdb.StubQueryError(sqlSelectUserWithPassword, errors.New("Forged error: SelectUserWithPassword."))
		testdb.StubQueryError(sqlSelectUserWithPasswordAlt, errors.New("Forged error: SelectUserWithPassword."))
	}
}

func SelectUserWithPasswordAndError() {
	err := errors.New("Forged error: SelectUserWithPasswordAndError.")
	testdb.StubQueryError(sqlSelectUserWithPassword, err)
	testdb.StubQueryError(sqlSelectUserWithPasswordAlt, err)
}

func UpdateUser(result Result) {
	if result == ResultSuccess {
		testdb.StubExec(sqlUpdateUser, testdb.NewResult(1, nil, 1, nil))
		return
	}

	testdb.StubQueryError(sqlUpdateUser, errors.New("Forged error: UpdateUser."))
}

func InsertUser(result Result) {
	if result == ResultSuccess {
		testdb.StubExec(sqlInsertUser, testdb.NewResult(1, nil, 1, nil))
		return
	}

	testdb.StubQueryError(sqlInsertUser, errors.New("Forged error: InsertUser."))
}

func DeleteUser(result Result) {
	if result == ResultSuccess {
		testdb.StubExec(sqlDeleteUser, testdb.NewResult(1, nil, 1, nil))
		return
	}

	testdb.StubQueryError(sqlDeleteUser, errors.New("Forged error: DeleteUser."))
}
