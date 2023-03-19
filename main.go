package main

import (
	"os"
	"tugas4/config"
	"tugas4/controller"
	"tugas4/repository"
	"tugas4/routes"
	"tugas4/service"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.Connect()
	defer config.Close(db)

	jwtService := service.NewJWTService()

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService, jwtService)

	commentRepository := repository.NewCommentRepository(db)
	commentService := service.NewCommentService(commentRepository)
	commentController := controller.NewCommentController(commentService)

	blogRepository := repository.NewBlogRepository(db)
	blogService := service.NewBlogService(blogRepository)
	blogController := controller.NewBlogController(blogService, jwtService)

	server := gin.Default()
	routes.UserRoutes(server, userController, jwtService)
	routes.CommentRoutes(server, commentController, jwtService)
	routes.BlogRoutes(server, blogController, jwtService)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server.Run(":" + port)
}
