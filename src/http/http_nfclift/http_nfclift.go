package http_nfclift

import (
	"fmt"
	"github.com/JaimePalomo/nfcliftserver-ddd/src/domain/lifts"
	"github.com/JaimePalomo/nfcliftserver-ddd/src/domain/operators"
	"github.com/JaimePalomo/nfcliftserver-ddd/src/domain/tags"
	"github.com/JaimePalomo/nfcliftserver-ddd/src/services/nfclift_service"
	"github.com/federicoleon/bookstore_utils-go/rest_errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

type NfcLiftHandlerI interface {
	CreateLift(*gin.Context)
	GetLift(*gin.Context)
	DeleteLift(ctx *gin.Context)

	CreateOperator(*gin.Context)
	//	GetOperator(*gin.Context)
	DeleteOperator(*gin.Context)

	CreateTag(*gin.Context)
	DeleteTag(*gin.Context)

	Call(*gin.Context)
}

type nfcLiftHandler struct {
	service nfclift_service.NfcLiftServiceI
}

func New(service nfclift_service.NfcLiftServiceI) NfcLiftHandlerI {
	return &nfcLiftHandler{service: service}
}

func (n *nfcLiftHandler) CreateLift(c *gin.Context) {
	idOp, err := getIdOpFromQuery(c)
	if err != nil {
		c.JSON(err.Status(), err.Error())
		return
	}
	_, err = n.service.GetOperator(operators.Operator{Id: idOp})
	if err != nil {
		if err.Status() == http.StatusNotFound {
			restErr := rest_errors.NewBadRequestError("invalid operator id")
			c.JSON(restErr.Status(), restErr.Error())
			return
		}
		c.JSON(err.Status(), err.Error())
		return
	}

	var lift lifts.Lift
	if err := c.ShouldBindJSON(&lift); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr.Error())
		return
	}
	result, err := n.service.CreateLift(lift)
	if err != nil {
		c.JSON(err.Status(), err.Error())
		return
	}
	c.JSON(http.StatusCreated, result)
}

func (n *nfcLiftHandler) GetLift(c *gin.Context) {
	idOp, err := getIdOpFromQuery(c)
	if err != nil {
		c.JSON(err.Status(), err.Error())
		return
	}
	_, err = n.service.GetOperator(operators.Operator{Id: idOp})
	if err != nil {
		if err.Status() == http.StatusNotFound {
			restErr := rest_errors.NewBadRequestError("invalid operator id")
			c.JSON(restErr.Status(), restErr.Error())
			return
		}
		c.JSON(err.Status(), err.Error())
		return
	}

	var lift lifts.Lift
	liftRae, idErr := getIntParam(c.Param("lift_rae"))
	if idErr != nil {
		c.JSON(idErr.Status(), idErr)
		return
	}
	lift.Rae = liftRae
	result, err := n.service.GetLiftByRae(lift)
	if err != nil {
		c.JSON(err.Status(), err.Error())
		return
	}
	c.JSON(http.StatusOK, result)
}

func (n *nfcLiftHandler) DeleteLift(c *gin.Context) {
	idOp, err := getIdOpFromQuery(c)
	if err != nil {
		c.JSON(err.Status(), err.Error())
		return
	}
	_, err = n.service.GetOperator(operators.Operator{Id: idOp})
	if err != nil {
		if err.Status() == http.StatusNotFound {
			restErr := rest_errors.NewBadRequestError("invalid operator id")
			c.JSON(restErr.Status(), restErr.Error())
			return
		}
		c.JSON(err.Status(), err.Error())
		return
	}

	var lift lifts.Lift
	liftRae, idErr := getIntParam(c.Param("lift_rae"))
	if idErr != nil {
		c.JSON(idErr.Status(), idErr)
		return
	}
	lift.Rae = liftRae
	err = n.service.DeleteLiftByRae(lift)
	if err != nil {
		c.JSON(err.Status(), err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

func (n *nfcLiftHandler) CreateOperator(c *gin.Context) {
	var operator operators.Operator
	if err := c.ShouldBindJSON(&operator); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr.Error())
		return
	}
	result, err := n.service.CreateOperator(operator)
	if err != nil {
		c.JSON(err.Status(), err.Error())
		return
	}
	c.JSON(http.StatusCreated, result)
}

/*
func (n *nfcLiftHandler) GetOperator(c *gin.Context) {
	var operator operators.Operator
	operatorId, idErr := getOperatorId(c.Param("operator_id"))
	if idErr != nil {
		c.JSON(idErr.Status(), idErr)
		return
	}
	operator.Id = operatorId
	result, err := n.service.GetOperator(operator)
	if err != nil {
		c.JSON(err.Status(), err.Error())
		return
	}
	c.JSON(http.StatusOK, result)
} */

func (n *nfcLiftHandler) DeleteOperator(c *gin.Context) {
	var operator operators.Operator
	idOp, err := getIdOpFromQuery(c)
	if err != nil {
		c.JSON(err.Status(), err.Error())
		return
	}
	operator.Id = idOp
	err = n.service.DeleteOperator(operator)
	if err != nil {
		c.JSON(err.Status(), err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

func (n *nfcLiftHandler) CreateTag(c *gin.Context) {
	idOp, err := getIdOpFromQuery(c)
	if err != nil {
		c.JSON(err.Status(), err.Error())
		return
	}
	_, err = n.service.GetOperator(operators.Operator{Id: idOp})
	if err != nil {
		if err.Status() == http.StatusNotFound {
			restErr := rest_errors.NewBadRequestError("invalid operator id")
			c.JSON(restErr.Status(), restErr.Error())
			return
		}
		c.JSON(err.Status(), err.Error())
		return
	}
	var tag tags.Tag
	if err := c.ShouldBindJSON(&tag); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr.Error())
		return
	}
	result, err := n.service.RegisterTag(tag)
	if err != nil {
		c.JSON(err.Status(), err.Error())
		return
	}
	c.JSON(http.StatusCreated, result)
}

func (n *nfcLiftHandler) DeleteTag(c *gin.Context) {
	idOp, err := getIdOpFromQuery(c)
	if err != nil {
		c.JSON(err.Status(), err.Error())
		return
	}
	_, err = n.service.GetOperator(operators.Operator{Id: idOp})
	if err != nil {
		if err.Status() == http.StatusNotFound {
			restErr := rest_errors.NewBadRequestError("invalid operator id")
			c.JSON(restErr.Status(), restErr.Error())
			return
		}
		c.JSON(err.Status(), err.Error())
		return
	}
	var tag tags.Tag
	tagId, idErr := getStringParam(c.Param("tag_id"))
	if idErr != nil {
		c.JSON(idErr.Status(), idErr)
		return
	}
	tag.Id = tagId
	err = n.service.DeleteTag(tag)
	if err != nil {
		c.JSON(err.Status(), err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

func (n *nfcLiftHandler) Call(c *gin.Context) {
	var tag tags.Tag
	tagId, idErr := getStringParam(c.Param("tag_id"))
	if idErr != nil {
		c.JSON(idErr.Status(), idErr)
		return
	}
	tag.Id = tagId
	lift, err := n.service.CallToTag(tag)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	if lift == nil {
		c.JSON(http.StatusOK, map[string]string{"status": "called"})
		return
	}
	c.JSON(http.StatusOK, lift)
}

//Util functions

func getIntParam(param string) (int, rest_errors.RestErr) {
	intParam, parseErr := strconv.ParseInt(param, 10, 64)
	if parseErr != nil {
		return 0, rest_errors.NewBadRequestError(fmt.Sprintf("error parsing parameter"))
	}
	return int(intParam), nil
}

func getStringParam(param string) (string, rest_errors.RestErr) {
	if strings.TrimSpace(param) == "" {
		return "", rest_errors.NewBadRequestError(fmt.Sprintf("error parsing parameter"))
	}
	return param, nil
}

func getIdOpFromQuery(c *gin.Context) (string, rest_errors.RestErr) {
	var restErr rest_errors.RestErr
	value, exist := c.GetQuery("idOp")
	if !exist {
		restErr = rest_errors.NewBadRequestError("id operator is needed")
		return "", restErr
	}
	return value, restErr
}
