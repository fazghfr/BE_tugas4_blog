package controller

import (
	"net/http"
	"tugas4/dto"
	"tugas4/service"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CommentController interface {
	AddComment(ctx *gin.Context)
}

type commentController struct {
	commentService service.CommentService
}

func NewCommentController(cs service.CommentService) CommentController {
	return &commentController{
		commentService: cs,
	}
}

func (cc *commentController) AddComment(ctx *gin.Context) {
	var commentDTO dto.CommentDTO
	if err := ctx.ShouldBindJSON(&commentDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	blogID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid blog ID"})
		return
	}

	comment, err := cc.commentService.AddComment(ctx, blogID, commentDTO)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, comment)
}
