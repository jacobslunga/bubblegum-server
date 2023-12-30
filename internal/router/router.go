package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jacobslunga/bubblegum-server/internal/controller"
	"github.com/jacobslunga/bubblegum-server/internal/repository"
	"github.com/jacobslunga/bubblegum-server/internal/service"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	userRepo := repository.NewUserRepo(db)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	v1 := router.Group("api/v1")
	{
		v1.GET("/users/:userId", userController.GetUserById)
		v1.GET("/users/me", userController.GetMe)
		v1.POST("/users/auth/register", userController.Register)
		v1.POST("/users/auth/login", userController.Login)
		v1.PUT("/users", userController.UpdateUser)
		v1.DELETE("/users", userController.DeleteUser)
	}

	return router
}
