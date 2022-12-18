package db_tags

import (
	"github.com/JaimePalomo/nfcliftserver-ddd/src/domain/tags"
	"github.com/federicoleon/bookstore_utils-go/rest_errors"
)

type DbTagsI interface {
	RegisterFloorTag(tag tags.Tag) rest_errors.RestErr
	RegisterCabinTag(tag tags.Tag) rest_errors.RestErr
	GetTagById(id string) (*tags.Tag, rest_errors.RestErr)
	DeleteTag(tag tags.Tag) rest_errors.RestErr
}
