package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PrintPagesJob struct {
	JobId int `json:"jobId" binding:"required,gte=1000"`
	Pages int `json:"pages" binding:"required,gte=1,lte=1000"`
}

func Validations() {
	router := gin.Default()
	router.POST("/print", func(ctx *gin.Context) {
		var job PrintPagesJob
		if err := ctx.ShouldBind(&job); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Print job #%v started", job.JobId)})
	})
	router.Run(":5000")
}
