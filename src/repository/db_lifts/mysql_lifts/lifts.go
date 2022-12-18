package mysql_lifts

import (
	"database/sql"
	"github.com/JaimePalomo/nfcliftserver-ddd/src/domain/lifts"
	"github.com/JaimePalomo/nfcliftserver-ddd/src/repository/db_lifts"
	"github.com/JaimePalomo/nfcliftserver-ddd/toolkit/log"
	"github.com/federicoleon/bookstore_utils-go/rest_errors"
	"strings"
)

const (
	queryGetLiftByRae    = "SELECT rae, stops, description, address, company, appDescription, stopTexts, distance FROM lifts where rae = ?"
	queryInsertLift      = "INSERT INTO lifts(rae, stops, description, address, company, appDescription, stopTexts, distance) VALUES(?, ?, ?, ?, ?, ?, ?, ?)"
	queryDeleteLiftByRae = "DELETE FROM lifts where rae = ?"
)

type dbLift struct {
	db *sql.DB
}

// New gets a db controller for lifts
func New(dbConnection *sql.DB) db_lifts.DbLiftsI {
	return &dbLift{db: dbConnection}
}

// GetByRae gets a lift by its RAE from database
func (d dbLift) GetByRae(rae int) (*lifts.Lift, rest_errors.RestErr) {

	var restErr rest_errors.RestErr

	statement, err := d.db.Prepare(queryGetLiftByRae)
	if err != nil {
		restErr = rest_errors.NewInternalServerError("internal database error getting a new lift", err)
		log.Error(restErr)
		return nil, restErr
	}
	defer statement.Close()

	lift := &lifts.Lift{}
	result := statement.QueryRow(rae)
	err = result.Scan(
		&lift.Rae,
		&lift.Stops,
		&lift.Description,
		&lift.Address,
		&lift.Company,
		&lift.AppDescription,
		&lift.StopTexts,
		&lift.Distance)
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			restErr = rest_errors.NewNotFoundError("no lift found with given rae")
			log.Error(restErr)
			return nil, restErr
		}
		restErr = rest_errors.NewInternalServerError("error getting lift from database", err)
		return nil, restErr
	}
	return lift, nil
}

// Create inserts a lift in the database
func (d dbLift) Create(lift lifts.Lift) (restErr rest_errors.RestErr) {
	statement, err := d.db.Prepare(queryInsertLift)
	if err != nil {
		restErr = rest_errors.NewInternalServerError("internal database error inserting a new lift", err)
		log.Error(restErr)
		return
	}
	defer statement.Close()

	_, err = statement.Exec(
		lift.Rae,
		lift.Stops,
		lift.Description,
		lift.Address,
		lift.Company,
		lift.AppDescription,
		lift.StopTexts,
		lift.Distance)
	if err != nil {
		return rest_errors.NewInternalServerError("internal database error inserting a new lift", err)
	}
	return nil
}

// DeleteByRae deletes a lift in the database by its RAE
func (d dbLift) DeleteByRae(rae int) rest_errors.RestErr {
	statement, err := d.db.Prepare(queryDeleteLiftByRae)
	if err != nil {
		return rest_errors.NewInternalServerError("internal database error deleting a lift", err)
	}
	defer statement.Close()

	_, err = statement.Exec(rae)
	if err != nil {
		return rest_errors.NewInternalServerError("internal database error deleting a lift", err)
	}
	return nil
}
