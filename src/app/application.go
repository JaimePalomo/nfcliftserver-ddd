package app

import (
	"github.com/JaimePalomo/nfcliftserver-ddd/src/config"
	_ "github.com/JaimePalomo/nfcliftserver-ddd/src/config"
	"github.com/JaimePalomo/nfcliftserver-ddd/src/http/http_nfclift"
	"github.com/JaimePalomo/nfcliftserver-ddd/src/repository"
	"github.com/JaimePalomo/nfcliftserver-ddd/src/repository/db_lifts/mysql_lifts"
	"github.com/JaimePalomo/nfcliftserver-ddd/src/repository/db_operators/mysql_operators"
	"github.com/JaimePalomo/nfcliftserver-ddd/src/repository/db_tags/mysql_tags"
	"github.com/JaimePalomo/nfcliftserver-ddd/src/services/nfclift_service"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	db := repository.NewMySQLClient()
	dbLift := mysql_lifts.New(db)
	dbOperator := mysql_operators.New(db)
	dbTags := mysql_tags.New(db)
	nfcLiftService := nfclift_service.New(dbLift, dbOperator, dbTags)
	nfcLiftHandler := http_nfclift.New(nfcLiftService)

	mapUrls(router, nfcLiftHandler)

	router.Run(":" + config.Cfg.ApiPort)
}
