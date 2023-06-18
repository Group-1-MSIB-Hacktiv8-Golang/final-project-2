package handler

import (
	"final-project-2/database"
	"final-project-2/repository/comment_repository/comment_pg"
	"final-project-2/repository/photo_repository/photo_pg"
	"final-project-2/repository/socialmedia_repository/socialmedia_pg"
	"final-project-2/repository/user_repository/user_pg"
	"final-project-2/service"

	"github.com/gin-gonic/gin"
)

func StartApp() {
	database.InitiliazeDatabase()

	db := database.GetDatabaseInstance()

	userRepo := user_pg.NewUserPG(db)
	userService := service.NewUserService(userRepo)
	userHandler := NewUserHandler(userService)

	photoRepo := photo_pg.NewPhotoPG(db)
	photoService := service.NewPhotoService(photoRepo)
	photoHandler := NewPhotoHandler(photoService)

	commentRepo := comment_pg.NewCommentPG(db)
	commentService := service.NewCommentService(commentRepo)
	commentHandler := NewCommentHandler(commentService)

	socialmediaRepo := socialmedia_pg.NewSocialMediaPG(db)
	socialmediaService := service.NewSocialmediaService(socialmediaRepo)
	socialmediaHandler := NewSocialMediaHandler(socialmediaService)

	authService := service.NewAuthService(userRepo)

	route := gin.Default()

	userRoute := route.Group("/users")
	{
		// userRoute.Use(authService.Authentication())

		userRoute.POST("/register", userHandler.Register)
		userRoute.POST("/login", userHandler.Login)
		userRoute.PUT("/", authService.Authentication(), userHandler.UpdateUser)
		userRoute.DELETE("/", authService.Authentication(), userHandler.DeleteUser)
	}

	photoRoute := route.Group("/photos")
	{
		photoRoute.Use(authService.Authentication())

		photoRoute.POST("/", photoHandler.CreateNewPhoto)
		photoRoute.GET("/", photoHandler.GetAllPhotoByUserId)
		photoRoute.PUT("/:photoId", photoHandler.UpdatePhotoById)
		photoRoute.DELETE("/:photoId", photoHandler.DeletePhotoById)
	}

	commentRoute := route.Group("/comments")
	{
		commentRoute.Use(authService.Authentication())

		commentRoute.POST("/", commentHandler.CreateNewComment)
		commentRoute.GET("/", commentHandler.GetAllCommentByUserId)
		commentRoute.PUT("/:commentId", commentHandler.UpdateCommentById)
		commentRoute.DELETE("/:commentId", commentHandler.DeleteCommentById)
	}

	socialMediaRoute := route.Group("/socialmedias")
	{
		socialMediaRoute.Use(authService.Authentication())

		socialMediaRoute.POST("/", socialmediaHandler.CreateNewSocialMedia)
		socialMediaRoute.GET("/", socialmediaHandler.GetAllSocialMediaByUserId)
		socialMediaRoute.PUT("/:socialMediaId", socialmediaHandler.UpdateSocialMediaById)
		socialMediaRoute.DELETE("/:socialMediaId", socialmediaHandler.DeleteSocialMediaById)
	}

	route.Run()
}