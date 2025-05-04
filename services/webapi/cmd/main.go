package main

import (
	"fmt"
	"log"

	"github.com/bit365/iotcloud/services/webapi/config"
	"github.com/bit365/iotcloud/services/webapi/internal/database"
	"github.com/bit365/iotcloud/services/webapi/internal/middleware"
	"github.com/bit365/iotcloud/services/webapi/internal/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	appConfig := config.LoadConfig()

	router := gin.Default()

	database_connection_string := appConfig.GetString("database.iotcloud")
	database.RunMigrations(database_connection_string)
	db, err := database.InitDB(database_connection_string)

	if err != nil {
		log.Fatalf("Could not initialize database: %v", err)
	}

	router.Use(middleware.Logger())
	routers.SetupRouter(router, db)

	port := fmt.Sprintf(":%d", appConfig.GetInt("server.port"))
	router.Run(port)
}
