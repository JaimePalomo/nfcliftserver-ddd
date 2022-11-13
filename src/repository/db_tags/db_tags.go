package db_tags

import (
	"database/sql"
	"github.com/JaimePalomo/nfcliftserver-ddd/src/domain/tags"
	"github.com/federicoleon/bookstore_utils-go/rest_errors"
	"strings"
)

const (
	queryInsertFloorTag = "INSERT INTO tags(id, rae, planta) VALUES(?, ?, ?)"
	queryInsertCabinTag = "INSERT INTO tags(id, rae) VALUES(?, ?)"
	queryGetTag         = "SELECT id, rae, planta FROM tags WHERE id=?"
	queryDeleteTag      = "DELETE FROM lifts WHERE id = ?"
)

type DbTagsI interface {
	RegisterFloorTag(tag tags.Tag) rest_errors.RestErr
	RegisterCabinTag(tag tags.Tag) rest_errors.RestErr
	GetTagById(id string) (*tags.Tag, rest_errors.RestErr)
	DeleteTag(tag tags.Tag) rest_errors.RestErr
}

type dbTags struct {
	db *sql.DB
}

func New(dbConnection *sql.DB) DbTagsI {
	return &dbTags{db: dbConnection}
}

func (d dbTags) RegisterFloorTag(tag tags.Tag) rest_errors.RestErr {
	statement, err := d.db.Prepare(queryInsertFloorTag)
	if err != nil {
		return rest_errors.NewInternalServerError("internal database error registering a new tag", err)
	}
	defer statement.Close()

	_, err = statement.Exec(tag.Id, tag.Rae, tag.Planta)
	if err != nil {
		return rest_errors.NewInternalServerError("internal database error registering a new tag", err)
	}
	return nil
}

func (d dbTags) RegisterCabinTag(tag tags.Tag) rest_errors.RestErr {
	statement, err := d.db.Prepare(queryInsertCabinTag)
	if err != nil {
		return rest_errors.NewInternalServerError("internal database error registering a new tag", err)
	}
	defer statement.Close()

	_, err = statement.Exec(tag.Id, tag.Rae)
	if err != nil {
		return rest_errors.NewInternalServerError("internal database error registering a new tag", err)
	}
	return nil
}

func (d dbTags) GetTagById(id string) (*tags.Tag, rest_errors.RestErr) {
	statement, err := d.db.Prepare(queryGetTag)
	if err != nil {
		return nil, rest_errors.NewInternalServerError("internal database error getting tag", err)
	}
	defer statement.Close()

	tag := &tags.Tag{}
	result := statement.QueryRow(id)
	err = result.Scan(&tag.Id, &tag.Rae, &tag.Planta)
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return nil, rest_errors.NewNotFoundError("no tag found with given id")
		}
		return nil, rest_errors.NewInternalServerError("error getting tag from database", err)
	}
	return tag, nil
}

func (d dbTags) DeleteTag(tag tags.Tag) rest_errors.RestErr {
	statement, err := d.db.Prepare(queryDeleteTag)
	if err != nil {
		return rest_errors.NewInternalServerError("internal database error deleting a tag", err)
	}
	defer statement.Close()

	_, err = statement.Exec(tag.Id)
	if err != nil {
		return rest_errors.NewInternalServerError("internal database error deleting a tag", err)
	}
	return nil
}
