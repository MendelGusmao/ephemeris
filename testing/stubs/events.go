package stubs

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/rafaeljusto/go-testdb"
)

var (
	sqlSelectAllEvents = "SELECT * FROM `events`"
	sqlInsertEvent     = "INSERT INTO `events` (`beginning`,`created_at`,`description`,`end`,`logo_u_r_l`,`name`,`place`,`registration_beginning`,`registration_end`,`status`,`u_r_l`,`updated_at`,`user_id`,`visibility`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?)"

	eventFields = []string{
		"id",
		"name",
		"place",
		"description",
		"u_r_l",
		"logo_u_r_l",
		"beginning",
		"end",
		"registration_beginning",
		"registration_end",
		"visibility",
		"status",
		"user_id",
		"created_at",
		"updated_at",
	}
	eventData = []string{
		"1",
		"Test Event",
		"Test Place",
		"Test Description",
		"http://test.url",
		"http://test.url/test-logo.jpg",
		"2014-09-29 11:44:49.115155-03",
		"2014-09-29 11:44:49.115155-03",
		"2014-09-29 11:44:49.115155-03",
		"2014-09-29 11:44:49.115155-03",
		"public",
		"open",
		"1",
		"2014-08-25 19:01:05.492933-03",
		"2014-08-25 19:01:05.492933-03",
	}
)

func SelectAllEvents(hasEvents bool) {
	if hasEvents {
		testdb.StubQuery(sqlSelectAllEvents, testdb.RowsFromCSVString(eventFields, strings.Join(eventData, ",")))
		return
	}

	testdb.StubQueryError(sqlSelectAllEvents, sql.ErrNoRows)
}

func SelectAllEventsWithError() {
	testdb.StubQueryError(sqlInsertEvent, errors.New("Forged error: SelectAllEventsWithError."))
}

func InsertEvent() {
	testdb.StubExec(sqlInsertEvent, testdb.NewResult(1, nil, 1, nil))
}

func InsertEventWithError() {
	testdb.StubQueryError(sqlInsertEvent, errors.New("Forged error: InsertEventWithError."))
}
