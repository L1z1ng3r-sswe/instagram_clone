package handler

import (
	"github.com/L1z1ng3r-sswe/instagram_clone/app/internal/middleware"
	"github.com/L1z1ng3r-sswe/instagram_clone/app/pkg/logging"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SetUpRoutes() *gin.Engine {
	router := gin.New()
	router.Use(cors.Default())
	router.Use(logging.LoggingMiddleware())
	router.Static("/post_image", "./app/pkg/storage/post_image")

	v1 := router.Group("/api/v1")

	post := v1.Group("/post")
	{
		post.POST("/", middleware.IsAuthMiddleware(), h.CreatePost)
		post.GET("/", h.GetPosts)
	}

	user := v1.Group("/user")
	{
		user.POST("/sign-up", h.SignUp)
		user.POST("/sign-in", h.SignIn)
	}

	return router
}
