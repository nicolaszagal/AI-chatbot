package routes

import (
	"chatbot/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/", handlers.Index)
	router.POST("/get_answer", handlers.GetAnswer)
	router.GET("/list_videos", handlers.ListVideos)
	router.POST("/transcribe", handlers.Transcribe)
	router.POST("/chat-with-video", handlers.ChatWithVideo)
}
