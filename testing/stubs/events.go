package stubs

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/rafaeljusto/go-testdb"
)

var (
	sqlSelectAllEvents = "SELECT * FROM `events`"
	sqlSelectEvent     = "SELECT  * FROM `events`  WHERE (`id` = ?) LIMIT 1"
	sqlInsertEvent     = "INSERT INTO `events` (`beginning`,`created_at`,`description`,`end`,`logo_u_r_l`,`name`,`place`,`registration_beginning`,`registration_end`,`status`,`u_r_l`,`updated_at`,`user_id`,`visibility`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	sqlUpdateEvent     = "UPDATE `events` SET `beginning` = ?, `created_at` = ?, `description` = ?, `end` = ?, `logo_u_r_l` = ?, `name` = ?, `place` = ?, `registration_beginning` = ?, `registration_end` = ?, `status` = ?, `u_r_l` = ?, `updated_at` = ?, `user_id` = ?, `visibility` = ?  WHERE (`id` = ?)"
	sqlDeleteEvent     = "DELETE FROM `events`  WHERE (`id` = ?)"

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

func SelectAllEvents(result Result) {
	switch result {
	case ResultSuccess:
		testdb.StubQuery(sqlSelectAllEvents, testdb.RowsFromCSVString(eventFields, strings.Join(eventData, ",")))
	case ResultNoRows:
		testdb.StubQueryError(sqlSelectAllEvents, sql.ErrNoRows)
	case ResultError:
		testdb.StubQueryError(sqlSelectAllEvents, errors.New("Forged error: SelectAllEvents."))
	}
}

func SelectEvent(result Result) {
	switch result {
	case ResultSuccess:
		rows := testdb.RowsFromCSVString(eventFields, strings.Join(eventData, ","))
		testdb.StubQuery(sqlSelectEvent, rows)
	case ResultNoRows:
		testdb.StubQueryError(sqlSelectEvent, sql.ErrNoRows)
	case ResultError:
		testdb.StubQueryError(sqlSelectEvent, errors.New("Forged error: SelectEvent."))
	}
}

func InsertEvent(result Result) {
	if result == ResultSuccess {
		testdb.StubExec(sqlInsertEvent, testdb.NewResult(1, nil, 1, nil))
		return
	}

	testdb.StubQueryError(sqlInsertEvent, errors.New("Forged error: InsertEvent."))
}

func UpdateEvent(result Result) {
	if result == ResultSuccess {
		testdb.StubExec(sqlUpdateEvent, testdb.NewResult(1, nil, 1, nil))
		return
	}

	testdb.StubQueryError(sqlUpdateEvent, errors.New("Forged error: UpdateEvent."))
}

func DeleteEvent(result Result) {
	if result == ResultSuccess {
		testdb.StubExec(sqlDeleteEvent, testdb.NewResult(1, nil, 1, nil))
		return
	}

	testdb.StubQueryError(sqlDeleteEvent, errors.New("Forged error: DeleteEvent."))
}
