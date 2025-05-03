package routers

import (
	"github.com/bit365/iotcloud/services/webapi/internal/controllers"
	"github.com/bit365/iotcloud/services/webapi/internal/repositories"
	"github.com/bit365/iotcloud/services/webapi/internal/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(router *gin.Engine, db *gorm.DB) {

	var apiGroup = router.Group("/api")
	{
		// Add User Routes
		{
			userGroup := apiGroup.Group("/users")

			var userRepository = repositories.NewUserRepository(db)
			var userService = services.NewUserService(userRepository)
			var userController = controllers.NewUserController(userService)

			userGroup.GET("/:id", userController.GetUserByID)
			userGroup.POST("/", userController.CreateUser)
			userGroup.PUT("/:id", userController.UpdateUser)
			userGroup.DELETE("/:id", userController.DeleteUser)
			userGroup.GET("/", userController.GetAllUsers)
		}
	}
}
