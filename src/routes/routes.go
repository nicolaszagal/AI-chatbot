package routes

import (
	handlers2 "chatbot/src/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/", handlers2.Index)
	router.POST("/get_answer", handlers2.GetAnswer)
	router.GET("/list_videos", handlers2.ListVideos)
	router.POST("/transcribe", handlers2.Transcribe)
	router.POST("/chat-with-video", handlers2.ChatWithVideo)
}
