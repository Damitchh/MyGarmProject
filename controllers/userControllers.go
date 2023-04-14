package controllers

import (
	"MygarmProject/database"
	"MygarmProject/helpers"
	"MygarmProject/models"
	"MygarmProject/repositories"
	"MygarmProject/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	appJSON = "application/JSON"
)

// UserRegister godoc
// @Summary Register for a User
// @Description Register of User
// @Tags users
// @Accept json
// @Produce json
// @Param models.User body models.User true "Register User"
// @Success 200 {object} models.User
// @Router /users/register [post]
func UserRegister(ctx *gin.Context) {
	db := database.GetDB()
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	contentType := helpers.GetContentType(ctx)
	var NewUser models.User

	if contentType == appJSON {
		ctx.ShouldBindJSON(&NewUser)
	} else {
		ctx.ShouldBind(&NewUser)
	}

	err := userService.CreateUser(&NewUser)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"id":        NewUser.ID,
		"email":     NewUser.Email,
		"full_name": NewUser.Username,
	})
}

// UserLogin godoc
// @Summary Login for a User
// @Description Login of User
// @Tags users
// @Accept json
// @Produce json
// @Success 200 "token"
// @Router /users/login [post]
func UserLogin(ctx *gin.Context) {
	db := database.GetDB()
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	contentType := helpers.GetContentType(ctx)
	User := models.User{}
	password := ""

	if contentType == appJSON {
		ctx.ShouldBindJSON(&User)
	} else {
		ctx.ShouldBind(&User)
	}
	password = User.Password

	err := userService.LoginUser(&User)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "invalid username/password",
		})
		return
	}

	comparePass := helpers.ComparePass([]byte(User.Password), []byte(password))
	if !comparePass {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "invalid username/password",
		})
		return
	}

	token := helpers.GenerateToken(User.ID, User.Email, User.Username)

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
