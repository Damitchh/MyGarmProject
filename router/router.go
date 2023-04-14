package router

import (
	"MygarmProject/controllers"
	_ "MygarmProject/docs"
	"MygarmProject/middlewares"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
	}

	photoRouter := r.Group("/photos")
	{
		photoRouter.Use(middlewares.Authentication())
		photoRouter.GET("/", controllers.GetPhotos)
		photoRouter.GET("/:ID", controllers.GetPhotoByID)
		photoRouter.POST("/", controllers.CreatePhoto)
		photoRouter.PUT("/:ID", controllers.UpdatePhoto)
		photoRouter.DELETE("/:ID", controllers.DeletePhotoByID)
	}

	commentRouter := r.Group("/comments")
	{
		commentRouter.Use(middlewares.Authentication())
		commentRouter.GET("/", controllers.GetComments)
		commentRouter.GET("/:ID", controllers.GetCommentByID)
		commentRouter.POST("/", controllers.CreateComment)
		commentRouter.PUT("/:ID", controllers.UpdateComment)
		commentRouter.DELETE("/:ID", controllers.DeleteCommentByID)
	}

	socialmediaRouter := r.Group("/socialmedias")
	{
		socialmediaRouter.Use(middlewares.Authentication())
		socialmediaRouter.GET("/", controllers.GetSocialMedias)
		socialmediaRouter.GET("/:ID", controllers.GetSocialMediaByID)
		socialmediaRouter.POST("/", controllers.CreateSocialMedia)
		socialmediaRouter.PUT("/:ID", controllers.UpdateSocialMedia)
		socialmediaRouter.DELETE("/:ID", controllers.DeleteSocialMediaByID)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
