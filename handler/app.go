package handler

import (
	"final-project-2/database"
	"final-project-2/repository/user_repository/user_pg"
	"final-project-2/service"

	"github.com/gin-gonic/gin"
)

func StartApp() {
	var port = "8080"
	database.InitiliazeDatabase()

	db := database.GetDatabaseInstance()

	userRepo := user_pg.NewUserPG(db)
	userService := service.NewUserService(userRepo)
	userHandler := NewUserHandler(userService)

	route := gin.Default()

	userRoute := route.Group("/users")
	{
		userRoute.POST("/register", userHandler.Register)

		userRoute.POST("/login", userHandler.Login)
	}

	route.Run(":" + port)
}