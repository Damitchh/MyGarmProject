package controllers

import (
	"MygarmProject/database"
	"MygarmProject/helpers"
	"MygarmProject/models"
	"MygarmProject/repositories"
	"MygarmProject/services"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// CreateComment godoc
// @Summary Create for a comment
// @Description Create of comment
// @Tags comments
// @Accept json
// @Produce json
// @Param models.Comment body models.Comment true "Create Comment"
// @Success 200 {object} models.Comment
// @Router /comments [post]
func CreateComment(ctx *gin.Context) {
	db := database.GetDB()
	commentRepo := repositories.NewCommentRepository(db)
	commentService := services.NewCommentService(commentRepo)
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	contentType := helpers.GetContentType(ctx)

	var newComment models.Comment

	var err error
	if contentType == appJSON {
		err = ctx.ShouldBindJSON(&newComment)
	} else {
		err = ctx.ShouldBind(&newComment)
	}

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "Failed to parse request body",
		})
		return
	}

	newComment.UserID = userID

	err = commentService.CreateComment(&newComment)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"err":     "bad Request",
			"message": err.Error(),
		})
		return
	}

	newComment, err = commentService.GetCommentByID(newComment.ID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"err":     "cant get comment data",
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, newComment)
}

// GetCommentByID godoc
// @Summary Get details for a given id
// @Description Get details of comment corresponding to the input id
// @Tags comments
// @Accept json
// @Produce json
// @Param id path int true "ID of the comment"
// @Success 200 {object} models.Comment
// @Router /comments/{id} [get]
func GetCommentByID(ctx *gin.Context) {
	db := database.GetDB()
	commentRepo := repositories.NewCommentRepository(db)
	commentService := services.NewCommentService(commentRepo)
	commentId, err := strconv.Atoi(ctx.Param("ID"))

	result, err := commentService.GetCommentByID(uint(commentId))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error message": "Data Not Found",
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Photo Data": result,
	})

}

// GetComments godoc
// @Summary Get details
// @Description Get details of all comment
// @Tags comments
// @Accept json
// @Produce json
// @Success 200 {object} models.Comment
// @Router /comments [get]
func GetComments(ctx *gin.Context) {
	db := database.GetDB()
	commentRepo := repositories.NewCommentRepository(db)
	commentService := services.NewCommentService(commentRepo)

	results, err := commentService.GetAllComments()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error message": err,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Comments data : ": results,
	})

}

// UpdateComment godoc
// @Summary Update comment identified by the given id
// @Description Update comment identified corresponding to the input id
// @Tags comments
// @Accept json
// @Produce json
// @Param id path int true "ID of the comment to be updated"
// @Success 200 {object} models.Comment
// @Router /comments/{id} [patch]
func UpdateComment(ctx *gin.Context) {
	db := database.GetDB()
	commentRepo := repositories.NewCommentRepository(db)
	commentService := services.NewCommentService(commentRepo)
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(ctx)
	updatedComment := models.Comment{}
	userID := uint(userData["id"].(float64))

	var err error
	if contentType == appJSON {
		err = ctx.ShouldBindJSON(&updatedComment)
	} else {
		err = ctx.ShouldBind(&updatedComment)
	}

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "Failed to parse request body",
		})
		return
	}

	commentId, err := strconv.Atoi(ctx.Param("ID"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "bad request",
			"message": "params invalid",
		})
	}
	updatedComment.ID = uint(commentId)

	rowsAffected, err := commentService.UpdateCommentByID(&updatedComment, userID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "bad Request",
			"message": err.Error(),
		})
		return
	}
	if rowsAffected == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error message :": fmt.Sprintf("Photo with id %v not found", commentId),
		})
		return
	}

	updatedComment, err = commentService.GetCommentByID(updatedComment.ID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"err":     "cant get photo data",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":       "comment has been updated successfully",
		"rows_affected": rowsAffected,
		"Updated Data":  updatedComment,
	})
}

// DeleteCommentByID godoc
// @Summary Delete comment identified by the given id
// @Title Delete the order corresponding to the input id
// @Tags comments
// @Accept json
// @Produce json
// @Param id path int true "ID of the comment to be deleted"
// @Success 204 "No Content"
// @Router /comments/{id} [delete]
func DeleteCommentByID(ctx *gin.Context) {
	db := database.GetDB()
	commentRepo := repositories.NewCommentRepository(db)
	commentService := services.NewCommentService(commentRepo)
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	commentId, err := strconv.Atoi(ctx.Param("ID"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "bad request",
			"message": "params invalid",
		})
	}

	rowsAffected, err := commentService.DeleteCommentByID(uint(commentId), userID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error message": err.Error(),
		})
		return
	}
	if rowsAffected == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error message :": fmt.Sprintf("Comment with id %v not found", commentId),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":       "comment has been deleted successfully",
		"rows_affected": rowsAffected,
	})
}
