package nfclift_service

import (
	"github.com/JaimePalomo/nfcliftserver-ddd/src/domain/lifts"
	"github.com/JaimePalomo/nfcliftserver-ddd/src/domain/tags"
	"github.com/federicoleon/bookstore_utils-go/rest_errors"
)

func (s *service) CallToTag(tag tags.Tag) (*lifts.Lift, rest_errors.RestErr) {
	filledTag, err := s.dbTags.GetTagById(tag.Id)
	if err != nil {
		return nil, err
	}
	if filledTag.IsFloor() {
		return nil, s.callToFloorTag(*filledTag)
	}
	return s.callToCabinTag(*filledTag)
}

func (s *service) callToFloorTag(tag tags.Tag) rest_errors.RestErr {
	////////////////////////////////////////////////////////////

	// TODO: Placeholder for an internal call of the lift to the floor

	////////////////////////////////////////////////////////////
	return nil
}

func (s *service) callToCabinTag(tag tags.Tag) (*lifts.Lift, rest_errors.RestErr) {
	lift, err := s.GetLiftByRae(lifts.Lift{Rae: tag.Rae})
	if err != nil {
		return nil, err
	}
	return lift, nil
}
