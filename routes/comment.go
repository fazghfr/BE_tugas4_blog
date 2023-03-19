package routes

import (
	"tugas4/controller"
	"tugas4/middleware"
	"tugas4/service"

	"github.com/gin-gonic/gin"
)

func CommentRoutes(router *gin.Engine, ComController controller.CommentController, jwtservice service.JWTService) {
	commentRoutes := router.Group("/api/comment")
	{
		commentRoutes.POST("/:id", middleware.Authenticate(jwtservice), ComController.AddComment)
	}
}
