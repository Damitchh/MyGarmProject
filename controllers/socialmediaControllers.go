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

// CreateSocialMedia godoc
// @Summary Create for a socialmedia
// @Description Create of socialmedia
// @Tags socialmedias
// @Accept json
// @Produce json
// @Param models.SocialMedia body models.SocialMedia true "Create Social Media"
// @Success 200 {object} models.SocialMedia
// @Router /socialmedias [post]
func CreateSocialMedia(ctx *gin.Context) {
	db := database.GetDB()
	socialmediaRepo := repositories.NewSocialMediaRepository(db)
	socialmediaService := services.NewSocialMediaService(socialmediaRepo)
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	contentType := helpers.GetContentType(ctx)

	var newSocialMedia models.SocialMedia

	var err error
	if contentType == appJSON {
		err = ctx.ShouldBindJSON(&newSocialMedia)
	} else {
		err = ctx.ShouldBind(&newSocialMedia)
	}

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "Failed to parse request body",
		})
		return
	}

	newSocialMedia.UserID = userID
	err = socialmediaService.CreateSocialMedia(&newSocialMedia)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"err":     "bad Request",
			"message": err.Error(),
		})
		return
	}

	newSocialMedia, err = socialmediaService.GetSocialMediaByID(newSocialMedia.ID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"err":     "cant get social media data",
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, newSocialMedia)
}

// GetSocialMediaByID godoc
// @Summary Get details for a given id
// @Description Get details of socialmedia corresponding to the input id
// @Tags socialmedias
// @Accept json
// @Produce json
// @Param id path int true "ID of the socialmedia"
// @Success 200 {object} models.SocialMedia
// @Router /socialmedias/{id} [get]
func GetSocialMediaByID(ctx *gin.Context) {
	db := database.GetDB()
	socialmediaRepo := repositories.NewSocialMediaRepository(db)
	socialmediaService := services.NewSocialMediaService(socialmediaRepo)
	photoId, err := strconv.Atoi(ctx.Param("ID"))

	result, err := socialmediaService.GetSocialMediaByID(uint(photoId))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error message": "Data Not Found",
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Photo Data": result,
	})

}

// GetSocialMedias godoc
// @Summary Get details
// @Description Get details of all socialmedias
// @Tags socialmedias
// @Accept json
// @Produce json
// @Success 200 {object} models.SocialMedia
// @Router /socialmedias [get]
func GetSocialMedias(ctx *gin.Context) {
	db := database.GetDB()
	SocialMediaRepo := repositories.NewSocialMediaRepository(db)
	SocialMediaService := services.NewSocialMediaService(SocialMediaRepo)

	results, err := SocialMediaService.GetAllSocialMedias()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error message": err,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Social Media data : ": results,
	})

}

// UpdateSocialMedia godoc
// @Summary Update socialmedia identified by the given id
// @Description Update socialmedia identified corresponding to the input id
// @Tags socialmedias
// @Accept json
// @Produce json
// @Param id path int true "ID of the socialmedia to be updated"
// @Success 200 {object} models.SocialMedia
// @Router /socialmedias/{id} [patch]
func UpdateSocialMedia(ctx *gin.Context) {
	db := database.GetDB()
	socialmediaRepo := repositories.NewSocialMediaRepository(db)
	socialmediaService := services.NewSocialMediaService(socialmediaRepo)
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(ctx)
	updatedSocialMedia := models.SocialMedia{}
	userID := uint(userData["id"].(float64))

	var err error
	if contentType == appJSON {
		err = ctx.ShouldBindJSON(&updatedSocialMedia)
	} else {
		err = ctx.ShouldBind(&updatedSocialMedia)
	}

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "Failed to parse request body",
		})
		return
	}

	socialmediaId, err := strconv.Atoi(ctx.Param("ID"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "bad request",
			"message": "params invalid",
		})
	}
	updatedSocialMedia.ID = uint(socialmediaId)

	//get photo data
	socialmediaData, err := socialmediaService.GetSocialMediaByID(uint(socialmediaId))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"err":     "cant get social media data",
			"message": err.Error(),
		})
		return
	}

	// Check owner
	if socialmediaData.UserID != userID {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "you can only update your own social media",
		})
		return
	}

	rowsAffected, err := socialmediaService.UpdateSocialMediaByID(&updatedSocialMedia, userID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "bad Request",
			"message": err.Error(),
		})
		return
	}
	if rowsAffected == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error message :": fmt.Sprintf("Social Media with id %v not found", socialmediaId),
		})
		return
	}

	updatedSocialMedia, err = socialmediaService.GetSocialMediaByID(updatedSocialMedia.ID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"err":     "cant get social media data",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":       "social media has been updated successfully",
		"rows_affected": rowsAffected,
		"Updated Data":  updatedSocialMedia,
	})
}

// DeleteSocialMediaByID godoc
// @Summary Delete socialmedia identified by the given id
// @Data Delete the order corresponding to the input id
// @Tags socialmedias
// @Accept json
// @Produce json
// @Param id path int true "ID of the socialmedia to be deleted"
// @Success 204 "No Content"
// @Router /socialmedias/{id} [delete]
func DeleteSocialMediaByID(ctx *gin.Context) {
	db := database.GetDB()
	socialmediaRepo := repositories.NewSocialMediaRepository(db)
	socialmediaService := services.NewSocialMediaService(socialmediaRepo)
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	socialmediaId, err := strconv.Atoi(ctx.Param("ID"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "bad request",
			"message": "params invalid",
		})
	}
	//get photo data
	socialmediaData, err := socialmediaService.GetSocialMediaByID(uint(socialmediaId))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"err":     "cant get social media data",
			"message": err.Error(),
		})
		return
	}

	// Check owner
	if socialmediaData.UserID != userID {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "you can only update your own social media",
		})
		return
	}

	rowsAffected, err := socialmediaService.DeleteSocialMediaByID(uint(socialmediaId), userID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error message": err.Error(),
		})
		return
	}
	if rowsAffected == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error message :": fmt.Sprintf("social media with id %v not found", socialmediaId),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":       "social media has been deleted successfully",
		"rows_affected": rowsAffected,
	})
}
