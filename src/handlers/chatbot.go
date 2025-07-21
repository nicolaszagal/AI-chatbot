package handlers

import (
	"chatbot/src/services"
	"chatbot/src/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAnswer(c *gin.Context) {
	var body struct {
		Question string `json:"question"`
	}

	if err := c.BindJSON(&body); err != nil || body.Question == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Question is required"})
		return
	}

	answer := services.GetAnswer(body.Question)
	c.JSON(http.StatusOK, gin.H{"answer": answer})

}

func ChatWithVideo(c *gin.Context) {
	var body struct {
		VideoName string `json:"video_name"`
		Question  string `json:"question"`
	}

	if err := c.BindJSON(&body); err != nil || body.VideoName == "" || body.Question == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "VideoName and Question are required"})
	}

	transcriptPath := utils.BuildTranscriptionPath(body.VideoName)

	response, err := services.ChatWithTranscript(transcriptPath, body.Question)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "IA answer could not be generated"})
		return
	}

	c.Data(http.StatusOK, "application/json", []byte(response))
}

func Index(c *gin.Context) {
	c.File("./templates/index.html")
}
