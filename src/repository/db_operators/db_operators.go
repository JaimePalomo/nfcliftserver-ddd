package db_operators

import (
	"database/sql"
	"github.com/JaimePalomo/nfcliftserver-ddd/src/domain/operators"
	"github.com/federicoleon/bookstore_utils-go/rest_errors"
	"strings"
)

const (
	queryGetOperatorById    = "SELECT id FROM operators where id = ?"
	queryInsertOperator     = "INSERT INTO operators(id) VALUES(?)"
	queryDeleteOperatorById = "DELETE FROM operators where id = ?"
)

type DbOperatorsI interface {
	GetById(id int) (*operators.Operator, rest_errors.RestErr)
	Create(operator operators.Operator) rest_errors.RestErr
	DeleteById(id int) rest_errors.RestErr
}

type dbOperators struct {
	db *sql.DB
}

// New gets a db controller for lifts
func New(dbConnection *sql.DB) DbOperatorsI {
	return &dbOperators{db: dbConnection}
}

func (d dbOperators) GetById(id int) (*operators.Operator, rest_errors.RestErr) {
	statement, err := d.db.Prepare(queryGetOperatorById)
	if err != nil {
		return nil, rest_errors.NewInternalServerError("internal database error getting operator", err)
	}
	defer statement.Close()

	operator := &operators.Operator{}
	result := statement.QueryRow(id)
	err = result.Scan(
		&operator.Id)
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return nil, rest_errors.NewNotFoundError("no operator found with given id")
		}
		return nil, rest_errors.NewInternalServerError("error getting operator from database", err)
	}
	return operator, nil
}

func (d dbOperators) Create(operator operators.Operator) rest_errors.RestErr {
	statement, err := d.db.Prepare(queryInsertOperator)
	if err != nil {
		return rest_errors.NewInternalServerError("internal database error inserting a new operator", err)
	}
	defer statement.Close()

	_, err = statement.Exec(
		operator.Id)
	if err != nil {
		return rest_errors.NewInternalServerError("internal database error inserting a new operator", err)
	}
	return nil
}

func (d dbOperators) DeleteById(id int) rest_errors.RestErr {
	statement, err := d.db.Prepare(queryDeleteOperatorById)
	if err != nil {
		return rest_errors.NewInternalServerError("internal database error deleting a operator", err)
	}
	defer statement.Close()

	_, err = statement.Exec(id)
	if err != nil {
		return rest_errors.NewInternalServerError("internal database error deleting a operator", err)
	}
	return nil
}
