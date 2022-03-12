package middlewares

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func IsLogged() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		cookie := session.Get("user")
	
		if cookie == nil {
			ctx.JSON(401, gin.H{
				"message":"Log in to Acess this page!",
			})
			return
		}

		ctx.Next()
	}
}
