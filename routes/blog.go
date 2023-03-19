package routes

import (
	"tugas4/controller"
	"tugas4/middleware"
	"tugas4/service"

	"github.com/gin-gonic/gin"
)

func BlogRoutes(router *gin.Engine, BlogController controller.BlogController, jwtService service.JWTService) {
	blogRoutes := router.Group("/api/blog")
	{
		blogRoutes.POST("", middleware.Authenticate(jwtService), BlogController.CreateBlog)
		blogRoutes.PUT("/:id", middleware.Authenticate(jwtService), BlogController.LikeBlog)
		blogRoutes.GET("", BlogController.GetAllBlog)
		blogRoutes.GET("/:id", BlogController.GetBlogByID)

	}
}
