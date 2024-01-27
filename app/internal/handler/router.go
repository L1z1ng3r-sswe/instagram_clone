package handler

import "github.com/gin-gonic/gin"

func (h *Handler) SetUpRoutes() *gin.Engine {
	router := gin.New()

	{
		post := router.Group("/post")
		post.POST("/", h.CreatePost)
	}

	return router
}
