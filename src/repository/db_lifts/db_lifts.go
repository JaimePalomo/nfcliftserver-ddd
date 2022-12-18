package db_lifts

import (
	"github.com/JaimePalomo/nfcliftserver-ddd/src/domain/lifts"
	"github.com/federicoleon/bookstore_utils-go/rest_errors"
)

type DbLiftsI interface {
	GetByRae(rae int) (*lifts.Lift, rest_errors.RestErr)
	Create(lift lifts.Lift) rest_errors.RestErr
	DeleteByRae(rae int) rest_errors.RestErr
}
