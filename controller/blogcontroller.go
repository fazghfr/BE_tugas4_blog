package controller

import (
	"net/http"
	"tugas4/common"
	"tugas4/dto"
	"tugas4/service"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type BlogController interface {
	CreateBlog(ctx *gin.Context)
	GetAllBlog(ctx *gin.Context)
	GetBlogByID(ctx *gin.Context)
	LikeBlog(crx *gin.Context)
}

type blogController struct {
	jwtService  service.JWTService
	blogService service.BlogService
}

func NewBlogController(bs service.BlogService, jwts service.JWTService) BlogController {
	return &blogController{
		blogService: bs,
		jwtService:  jwts,
	}
}

func (bc *blogController) CreateBlog(ctx *gin.Context) {
	var blog dto.BlogCreateDTO
	err := ctx.ShouldBind(&blog)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Membuat Blog", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	token := ctx.MustGet("token").(string)
	userID, err := bc.jwtService.GetUserIDByToken(token)
	if err != nil {
		response := common.BuildErrorResponse("Gagal Memproses Request", "Token Tidak Valid", nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	blog.UserID = userID
	result, err := bc.blogService.CreateBlog(ctx.Request.Context(), blog)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Membuat Blog", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "Berhasil Membuat Blog", result)
	ctx.JSON(http.StatusOK, res)
}

func (bc *blogController) GetAllBlog(ctx *gin.Context) {
	result, err := bc.blogService.GetAllBlogs(ctx.Request.Context())
	if err != nil {
		res := common.BuildErrorResponse("Gagal Mendapatkan List Blog", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "Berhasil Mendapatkan List Blog", result)
	ctx.JSON(http.StatusOK, res)
}

func (bc *blogController) LikeBlog(ctx *gin.Context) {
	blogID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid blog ID"})
		return
	}

	result, err := bc.blogService.LikeBlog(ctx, blogID)
	if err != nil {
		res := common.BuildErrorResponse("Gagal memberi like pada Blog", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "Berhasil Mendapatkan List Blog", result)
	ctx.JSON(http.StatusOK, res)
}
func (bc *blogController) GetBlogByID(ctx *gin.Context) {
	blogID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid blog ID"})
		return
	}

	token := ctx.MustGet("token").(string)
	_, err = bc.jwtService.GetUserIDByToken(token)
	if err != nil {
		response := common.BuildErrorResponse("Gagal Memproses Request", "Token Tidak Valid", nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	result, err := bc.blogService.FindBlogByID(ctx, blogID)
	if err != nil {
		res := common.BuildErrorResponse("Gagal Mendapatkan List Blog", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "Berhasil Mendapatkan List Blog", result)
	ctx.JSON(http.StatusOK, res)
}
