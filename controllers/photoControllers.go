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

// CreatePhoto godoc
// @Summary Create for a photo
// @Description Create of photo
// @Tags photos
// @Accept json
// @Produce json
// @Param models.Photo body models.Photo true "Create Photo"
// @Success 200 {object} models.Photo
// @Router /photos [post]
func CreatePhoto(ctx *gin.Context) {
	db := database.GetDB()
	photoRepo := repositories.NewPhotoRepository(db)
	photoService := services.NewPhotoService(photoRepo)
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	contentType := helpers.GetContentType(ctx)

	var newPhoto models.Photo

	var err error
	if contentType == "application/json" {
		err = ctx.ShouldBindJSON(&newPhoto)
	} else {
		err = ctx.ShouldBind(&newPhoto)
	}

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "Failed to parse request body",
		})
		return
	}

	fmt.Println("photo url", newPhoto.PhotoUrl)
	fmt.Println("userid", userID)

	newPhoto.UserID = userID
	fmt.Println("photo : ", newPhoto.UserID)
	err = photoService.CreatePhoto(&newPhoto)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"err":     "bad Request",
			"message": err.Error(),
		})
		return
	}

	newPhoto, err = photoService.GetPhotoByID(newPhoto.ID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"err":     "cant get photo data",
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, newPhoto)
}

// GetPhotoByID godoc
// @Summary Get details for a given id
// @Description Get details of photo corresponding to the input id
// @Tags photos
// @Accept json
// @Produce json
// @Param ID path int true "ID of the photo"
// @Success 200 {object} models.Photo
// @Router /photos/{id} [get]
func GetPhotoByID(ctx *gin.Context) {
	db := database.GetDB()
	photoRepo := repositories.NewPhotoRepository(db)
	photoService := services.NewPhotoService(photoRepo)
	photoId, err := strconv.Atoi(ctx.Param("ID"))
	//userData := ctx.MustGet("userData").(jwt.MapClaims)
	//userID := uint(userData["id"].(float64))

	result, err := photoService.GetPhotoByID(uint(photoId))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error message": "Data Not Found",
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Photo Data": result,
	})

}

// GetPhotos godoc
// @Summary Get details
// @Description Get details of all photos
// @Tags photos
// @Accept json
// @Produce json
// @Success 200 {object} models.Photo
// @Router /photos [get]
func GetPhotos(ctx *gin.Context) {
	db := database.GetDB()
	photoRepo := repositories.NewPhotoRepository(db)
	photoService := services.NewPhotoService(photoRepo)

	results, err := photoService.GetPhotos()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error message": err,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Product data : ": results,
	})

}

// UpdatePhoto godoc
// @Summary Update photo identified by the given id
// @Description Update photo identified corresponding to the input id
// @Tags photos
// @Accept json
// @Produce json
// @Param ID path int true "ID of the photo to be updated"
// @Success 200 {object} models.Photo
// @Router /photos/{id} [patch]
func UpdatePhoto(ctx *gin.Context) {
	db := database.GetDB()
	photoRepo := repositories.NewPhotoRepository(db)
	photoService := services.NewPhotoService(photoRepo)
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(ctx)
	updatedPhoto := models.Photo{}
	userID := uint(userData["id"].(float64))

	var err error
	if contentType == "application/json" {
		err = ctx.ShouldBindJSON(&updatedPhoto)
	} else {
		err = ctx.ShouldBind(&updatedPhoto)
	}

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "Failed to parse request body",
		})
		return
	}

	photoId, err := strconv.Atoi(ctx.Param("ID"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "bad request",
			"message": "params invalid",
		})
	}
	updatedPhoto.ID = uint(photoId)

	//get photo data
	photoData, err := photoService.GetPhotoByID(uint(photoId))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"err":     "cant get photo data",
			"message": err.Error(),
		})
		return
	}

	// Check owner
	if photoData.UserID != userID {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "you can only update your own photos",
		})
		return
	}

	rowsAffected, err := photoService.UpdatePhotoByID(&updatedPhoto, userID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "bad Request",
			"message": err.Error(),
		})
		return
	}
	if rowsAffected == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error message :": fmt.Sprintf("Photo with id %v not found", photoId),
		})
		return
	}

	updatedPhoto, err = photoService.GetPhotoByID(updatedPhoto.ID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"err":     "cant get photo data",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":       "photo has been updated successfully",
		"rows_affected": rowsAffected,
		"Updated Data":  updatedPhoto,
	})
}

// DeletePhotoByID godoc
// @Summary Delete photo identified by the given id
// @Title Delete the order corresponding to the input id
// @Tags photos
// @Accept json
// @Produce json
// @Param ID path int true "ID of the photo to be deleted"
// @Success 204 "No Content"
// @Router /photos/{id} [delete]
func DeletePhotoByID(ctx *gin.Context) {
	db := database.GetDB()
	photoRepo := repositories.NewPhotoRepository(db)
	photoService := services.NewPhotoService(photoRepo)
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	photoId, err := strconv.Atoi(ctx.Param("ID"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "bad request",
			"message": "params invalid",
		})
	}
	//get photo data
	photoData, err := photoService.GetPhotoByID(uint(photoId))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"err":     "cant get photo data",
			"message": err.Error(),
		})
		return
	}

	// Check owner
	if photoData.UserID != userID {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "you can only update your own photos",
		})
		return
	}

	rowsAffected, err := photoService.DeletePhotoByID(uint(photoId), userID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error message": err.Error(),
		})
		return
	}
	if rowsAffected == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error message :": fmt.Sprintf("Product with id %v not found", photoId),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":       "product has been deleted successfully",
		"rows_affected": rowsAffected,
	})
}
