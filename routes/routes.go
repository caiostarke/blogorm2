package routes

import (
	"github.com/caiostarke/studyToAdvanced/controllers"
	"github.com/caiostarke/studyToAdvanced/middlewares"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func Setup(router *gin.Engine) {
	router.GET("/", controllers.HomeController)

	router.GET("/register", controllers.RegisterGet)
	router.POST("/register", controllers.RegisterPost)

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("session", store))
	router.POST("/login", controllers.LoginController)

	user := router.Group("/u")
	user.Use(middlewares.IsLogged())
	{
		user.GET("/:username", controllers.UserPage)
		user.GET("/:username/publish", controllers.UserPublishGet)
		user.POST("/:username/post", controllers.UserPublishPost)
		user.GET("/:username/delete/post/:id", controllers.DeletePost)
		user.GET("/:username/update/post/:id", controllers.UpdatePostPage)
		user.POST("/:username/update/post/:id", controllers.UpdatePost)
	}

	// admin := router.Group("/admin")
	// admin.Use(middlewares.IsLogged())1w
	// {
	// 	admin.GET("/:username", controllers.AdminHandler)
	// }

	post := router.Group("/post")
	{
		post.GET("/:id", controllers.GetPost)
	}
}
