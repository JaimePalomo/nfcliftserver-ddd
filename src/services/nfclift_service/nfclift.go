package nfclift_service

import (
	"github.com/JaimePalomo/nfcliftserver-ddd/src/domain/lifts"
	"github.com/JaimePalomo/nfcliftserver-ddd/src/domain/operators"
	"github.com/JaimePalomo/nfcliftserver-ddd/src/domain/tags"
	"github.com/JaimePalomo/nfcliftserver-ddd/src/repository/db_lifts"
	"github.com/JaimePalomo/nfcliftserver-ddd/src/repository/db_operators"
	"github.com/JaimePalomo/nfcliftserver-ddd/src/repository/db_tags"
	"github.com/federicoleon/bookstore_utils-go/rest_errors"
)

type NfcLiftServiceI interface {
	CreateLift(lift lifts.Lift) (*lifts.Lift, rest_errors.RestErr)
	GetLiftByRae(lift lifts.Lift) (*lifts.Lift, rest_errors.RestErr)
	DeleteLiftByRae(lift lifts.Lift) rest_errors.RestErr

	CreateOperator(operator operators.Operator) (*operators.Operator, rest_errors.RestErr)
	GetOperator(operator operators.Operator) (*operators.Operator, rest_errors.RestErr)
	DeleteOperator(operator operators.Operator) rest_errors.RestErr

	RegisterTag(tag tags.Tag) (*tags.Tag, rest_errors.RestErr)
	DeleteTag(tag tags.Tag) rest_errors.RestErr

	CallToTag(tag tags.Tag) (*lifts.Lift, rest_errors.RestErr)
}

type service struct {
	dbLift      db_lifts.DbLiftsI
	dbOperators db_operators.DbOperatorsI
	dbTags      db_tags.DbTagsI
}

func New(dbLift db_lifts.DbLiftsI, dbOperators db_operators.DbOperatorsI, dbTags db_tags.DbTagsI) NfcLiftServiceI {
	return &service{
		dbLift:      dbLift,
		dbOperators: dbOperators,
		dbTags:      dbTags,
	}
}

// CreateLift inserts a lift in the database
func (s *service) CreateLift(lift lifts.Lift) (*lifts.Lift, rest_errors.RestErr) {
	err := lift.Validate()
	if err != nil {
		return nil, err
	}
	if err = s.dbLift.Create(lift); err != nil {
		return nil, err
	}
	return &lift, nil
}

// GetLiftByRae gets a lift by its rae
func (s *service) GetLiftByRae(lift lifts.Lift) (*lifts.Lift, rest_errors.RestErr) {
	result, err := s.dbLift.GetByRae(lift.Rae)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteLiftByRae deletes a lift by its rae
func (s *service) DeleteLiftByRae(lift lifts.Lift) rest_errors.RestErr {
	return s.dbLift.DeleteByRae(lift.Rae)
}

// CreateOperator inserts a lift in the database
func (s *service) CreateOperator(operator operators.Operator) (*operators.Operator, rest_errors.RestErr) {
	if err := s.dbOperators.Create(operator); err != nil {
		return nil, err
	}
	return &operator, nil
}

// GetOperator gets a operator
func (s *service) GetOperator(operator operators.Operator) (*operators.Operator, rest_errors.RestErr) {
	result, err := s.dbOperators.GetById(operator.Id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteOperator deletes a operator
func (s *service) DeleteOperator(operator operators.Operator) rest_errors.RestErr {
	return s.dbOperators.DeleteById(operator.Id)
}

func (s *service) RegisterTag(tag tags.Tag) (*tags.Tag, rest_errors.RestErr) {
	lift, err := s.dbLift.GetByRae(tag.Rae)
	if err != nil {
		return nil, err
	}
	if tag.IsFloor() {
		if lift.IsValidFloor(tag.Planta) == false {
			return nil, rest_errors.NewBadRequestError("invalid floor")
		}
		err = s.dbTags.RegisterFloorTag(tag)
		if err != nil {
			return nil, err
		}
		return &tag, nil
	}
	err = s.dbTags.RegisterCabinTag(tag)
	if err != nil {
		return nil, err
	}
	return &tag, nil
}

func (s *service) DeleteTag(tag tags.Tag) rest_errors.RestErr {
	return s.dbTags.DeleteTag(tag)
}
