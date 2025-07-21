package handlers

import (
	"chatbot/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListVideos(c *gin.Context) {
	videos, err := services.ListVideos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Videos could not be listed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"videos": videos})
}

func Transcribe(c *gin.Context) {
	var body struct {
		VideoName string `json:"video_name"`
	}

	if err := c.BindJSON(&body); err != nil || body.VideoName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Video name is required"})
		return
	}

	text, err := services.TranscribeVideo(body.VideoName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Video could not be transcribed"})
		return
	}

	c.Data(200, "application/json", []byte(text))
}
