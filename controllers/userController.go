package controllers

import (
	"net/http"

	"github.com/caiostarke/studyToAdvanced/database"
	"github.com/caiostarke/studyToAdvanced/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func UserPage(ctx *gin.Context) {
	username := ctx.Param("username")
	var user models.User
	var posts []models.Post

	var userPage bool

	session := sessions.Default(ctx)
	cookie := session.Get("user")

	if cookie != username {
		userPage = false
	} else {
		userPage = true
	}

	database.DB.Where("first_name = ?", username).First(&user)
	if user.ID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "User Not Found",
		})
		return
	}

	database.DB.Where("user_id = ?", user.ID).Find(&posts)

	ctx.HTML(http.StatusOK, "userpage.html", gin.H{
		"user":       user,
		"posts":      posts,
		"isuserpage": userPage,
	})
}

func RegisterGet(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "register.html", nil)
}

func RegisterPost(ctx *gin.Context) {
	var data models.RegisterForm
	ctx.ShouldBind(&data)

	if data.Password != data.ConfirmPassword {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Wrong Credentials!",
		})
		return
	}

	user := models.User{
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Email:     data.Email,
		Password:  []byte(data.Password), // Refactor it to a model that encrypts data
	}

	path := "/u/" + user.FirstName

	database.DB.Create(&user)
	ctx.Redirect(http.StatusFound, path)
}

func LoginController(ctx *gin.Context) {
	var data models.LoginForm
	ctx.ShouldBind(&data)
	var user models.User

	database.DB.Where("email = ?", data.Email).Find(&user)
	if user.ID == 0 {
		ctx.JSON(401, gin.H{
			"message": "Credentials wrong",
		})
		return
	}

	if data.Password != string(user.Password) {
		ctx.JSON(401, gin.H{
			"message": "Credentials wrong",
		})
		return
	}

	session := sessions.Default(ctx)

	if session.Get("user") == nil {
		session.Set("user", user.FirstName)
		session.Save()
	}

	ctx.JSON(200, gin.H{"hello": session.Get("user")})
}

func UserPublishGet(ctx *gin.Context) {
	username := ctx.Param("username")
	var userPage bool
	var user models.User

	session := sessions.Default(ctx)
	cookieUser := session.Get("user")
	database.DB.Where("first_name = ?", username).First(&user)

	if cookieUser != username {
		userPage = false
	} else {
		userPage = true
	}

	ctx.HTML(http.StatusOK, "user.html", gin.H{
		"userpage": userPage,
		"user":     user,
	})
}

func UserPublishPost(ctx *gin.Context) {
	username := ctx.Param("username")
	var user models.User
	var post models.Post

	session := sessions.Default(ctx)
	cookie := session.Get("user")

	if cookie != username {
		ctx.JSON(401, gin.H{
			"message": "Unauthorized",
		})
	} else {
		database.DB.Where("first_name = ?", username).First(&user)
		if user.ID == 0 {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": "User Not Found",
			})
			return
		}

		ctx.ShouldBind(&post)
		database.DB.Create(&post)
		ctx.Redirect(302, "/u/"+username)
	}

}
