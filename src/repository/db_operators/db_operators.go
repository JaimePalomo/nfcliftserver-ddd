package db_operators

import (
	"github.com/JaimePalomo/nfcliftserver-ddd/src/domain/operators"
	"github.com/federicoleon/bookstore_utils-go/rest_errors"
)

type DbOperatorsI interface {
	GetById(id string) (*operators.Operator, rest_errors.RestErr)
	Create(operator operators.Operator) rest_errors.RestErr
	DeleteById(id string) rest_errors.RestErr
}
