package controllers

import (
	"net/http"
	"strconv"

	"github.com/caiostarke/studyToAdvanced/database"
	"github.com/caiostarke/studyToAdvanced/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func GetPost(ctx *gin.Context) {
	var post models.Post
	var user models.User
	var tag models.Tag
	id := ctx.Param("id")

	database.DB.Where("id = ?", id).First(&post)
	if post.ID == 0 {
		ctx.JSON(404, gin.H{
			"message": "Post not found, are u sure did u put correct ID?",
		})
		return
	}
	post = LinkUserAndTag(post, tag, user)

	ctx.HTML(http.StatusOK, "singlepost.gotmpl", gin.H{
		"post": post,
	})
}

func LinkUserAndTag(post models.Post, tag models.Tag, user models.User) models.Post {
	user.ID = post.UserID
	tag.ID = post.TagID

	database.DB.Where("id = ?", user.ID).First(&user)
	database.DB.Where("id = ?", tag.ID).First(&tag)

	post.User = user
	post.Tag = tag
	return post
}

func DeletePost(ctx *gin.Context) {
	id := ctx.Param("id")
	userid := ctx.Param("username")
	var user models.User

	userIdInt, _ := strconv.Atoi(userid)
	database.DB.Where("id = ?", userIdInt).First(&user)

	session := sessions.Default(ctx)
	cookie := session.Get("user")

	if cookie != user.FirstName {
		ctx.JSON(405, gin.H{
			"message": "Not Allowed",
		})
		return
	}

	idInt, _ := strconv.Atoi(id)
	post := models.Post{}
	post.ID = idInt

	database.DB.Where("id = ?", post.ID).First(&post)

	if post.UserID != userIdInt {
		ctx.JSON(405, gin.H{
			"message": "Not Allowed",
		})
		return
	} else {
		database.DB.Delete(&post)
		ctx.Redirect(301, "/u/"+user.FirstName)
	}
}

func UpdatePostPage(ctx *gin.Context) {
	id := ctx.Param("id")
	userid := ctx.Param("username")

	var user models.User

	userIdInt, _ := strconv.Atoi(userid)
	database.DB.Where("id = ?", userIdInt).First(&user)

	session := sessions.Default(ctx)
	cookie := session.Get("user")

	if cookie != user.FirstName {
		ctx.JSON(405, gin.H{
			"message": "Not Allowed",
		})
		return
	}

	idInt, _ := strconv.Atoi(id)
	post := models.Post{}

	database.DB.Where("id = ?", idInt).First(&post)
	if post.ID == 0 {
		ctx.JSON(405, gin.H{
			"message": "Post not found",
		})
		return
	}

	if post.UserID != user.ID {
		ctx.JSON(405, gin.H{
			"message": "Not Allowed",
		})
		return
	} else {
		ctx.HTML(http.StatusOK, "updatepage.gotmpl", gin.H{
			"post":     post,
			"title":    "Update Page",
			"image":    "https://pbs.twimg.com/profile_images/1497164191798603803/yoLtCnFO_400x400.jpg",
			"jobautor": "Student to Infinity",
			"user":     user,
		})
	}
}

func UpdatePost(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id")) // Id Post
	userid := ctx.Param("username")        // Id User

	var user models.User
	post := models.Post{}

	database.DB.Where("id = ?", id).First(&post)

	if post.ID == 0 {
		ctx.JSON(405, gin.H{
			"message": "Post not found",
		})
		return
	}

	userIdInt, _ := strconv.Atoi(userid)
	database.DB.Where("id = ?", userIdInt).First(&user)

	session := sessions.Default(ctx)
	cookie := session.Get("user")

	if cookie != user.FirstName {
		ctx.JSON(405, gin.H{
			"message": "Not Allowed",
		})
		return
	}

	if post.UserID != user.ID {
		ctx.JSON(405, gin.H{
			"message": "Not Allowed",
		})
		return
	}

	ctx.ShouldBind(&post)
	post.User.ID = post.UserID
	post.Tag.ID = post.TagID
	database.DB.Save(&post)
	ctx.Redirect(301, "/u/"+user.FirstName)
}
