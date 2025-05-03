package main

import (
	"fmt"

	"github.com/bit365/iotcloud/services/webapi/config"
	"github.com/bit365/iotcloud/services/webapi/internal/database"
	"github.com/bit365/iotcloud/services/webapi/internal/middleware"
	"github.com/bit365/iotcloud/services/webapi/internal/routers"

	"github.com/gin-gonic/gin"
)

func main() {

	appConfig := config.LoadConfig()

	router := gin.Default()

	dsn := appConfig.GetString("database.dsn")
	database.RunMigrations(dsn)
	db, err := database.InitDB(dsn)

	if err != nil {
		panic(err)
	}

	router.Use(middleware.Logger())
	routers.SetupRouter(router, db)

	port := fmt.Sprintf(":%d", appConfig.GetInt("server.port"))
	router.Run(port)
}
