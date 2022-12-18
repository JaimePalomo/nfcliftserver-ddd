package app

import (
	"github.com/JaimePalomo/nfcliftserver-ddd/src/http/http_nfclift"
	"github.com/JaimePalomo/nfcliftserver-ddd/src/http/http_ping"
	"github.com/gin-gonic/gin"
)

func mapUrls(router *gin.Engine, nfcLiftHandler http_nfclift.NfcLiftHandlerI) {
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
}
