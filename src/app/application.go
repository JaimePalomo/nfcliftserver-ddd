package app

import (
	"github.com/JaimePalomo/nfcliftserver-ddd/src/http/http_nfclift"
	"github.com/JaimePalomo/nfcliftserver-ddd/src/http/http_ping"
	"github.com/JaimePalomo/nfcliftserver-ddd/src/repository"
	"github.com/JaimePalomo/nfcliftserver-ddd/src/repository/db_lifts"
	"github.com/JaimePalomo/nfcliftserver-ddd/src/repository/db_operators"
	"github.com/JaimePalomo/nfcliftserver-ddd/src/repository/db_tags"
	"github.com/JaimePalomo/nfcliftserver-ddd/src/services/nfclift_service"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	db := repository.New()
	dbLift := db_lifts.New(db)
	dbOperator := db_operators.New(db)
	dbTags := db_tags.New(db)
	nfcLiftService := nfclift_service.New(dbLift, dbOperator, dbTags)
	nfcLiftHandler := http_nfclift.New(nfcLiftService)

	router.POST("/nfclift/lift", nfcLiftHandler.CreateLift)
	router.GET("/nfclift/lift/:lift_rae", nfcLiftHandler.GetLift)
	router.DELETE("/nfclift/lift/:lift_rae", nfcLiftHandler.DeleteLift)

	router.POST("/nfclift/operator", nfcLiftHandler.CreateOperator)
	//router.GET("/nfclift/operator/:operator_id", nfcLiftHandler.GetOperator)
	router.DELETE("/nfclift/operator/:operator_id", nfcLiftHandler.DeleteOperator)

	router.POST("/nfclift/tag", nfcLiftHandler.CreateTag)
	router.DELETE("/nfclift/tag/:tag_id", nfcLiftHandler.DeleteTag)

	router.GET("/nfclift/tag/:tag_id", nfcLiftHandler.Call)

	router.GET("/ping", http_ping.Pong)

	router.Run(":8080")
}
