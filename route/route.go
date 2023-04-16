package route

import (
	"MyGram/handler"
	"MyGram/middlewares"
	"MyGram/service"

	"github.com/gin-gonic/gin"
)

func RegisterApi(r *gin.Engine, app service.ServiceInterface) {
	server := handler.NewHttpServer(app)

	user := r.Group("/user")
	{
		user.POST("register", server.UserRegister)
		user.POST("login", server.UserLogin)
	}

	photo := r.Group("/photo")
	{
		photo.Use(middlewares.Authentication())
		photo.GET("", server.PhotoGetAll)
		photo.GET(":id", server.PhotoGet)
		photo.POST("", server.PhotoCreate)
		photo.PUT(":id", server.PhotoUpdate)
		photo.DELETE(":id", server.PhotoDelete)
	}

	comment := r.Group("/comment")
	{
		comment.Use(middlewares.Authentication())
		comment.GET("photo/:id", server.CommentGetAll)
		comment.GET(":id", server.CommentGet)
		comment.POST(":id", server.CommentCreate)
		comment.PUT(":id", server.CommentUpdate)
		comment.DELETE(":id", server.CommentDelete)
	}

	sosmed := r.Group("/socialmedia")
	{
		sosmed.Use(middlewares.Authentication())
		sosmed.GET("", server.SocialMediaGetAll)
		sosmed.GET(":id", server.SocialMediaGet)
		sosmed.POST("", server.SocialMediaCreate)
		sosmed.PUT(":id", server.SocialMediaUpdate)
		sosmed.DELETE(":id", server.SocialMediaDelete)
	}
}
