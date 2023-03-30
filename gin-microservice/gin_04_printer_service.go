package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type PrintJob struct {
	JobId     int    `json:"jobId" binding:"gte=0"`
	InvoiceId int    `json:"invoiceId" binding:"required,gte=0"`
	Format    string `json:"format" binding:"required"`
}

func handlePrintJobs(ctx *gin.Context) {
	var job PrintJob
	if err := ctx.ShouldBindJSON(&job); err != nil {
		ctx.JSON(400, "Invalid input")
		return
	}
	log.Printf("PrintService: Creating new print job from invoice #%v", job.InvoiceId)
	rand.Seed(time.Now().UnixNano())
	job.JobId = rand.Intn(1000)
	log.Printf("PrintService: Created new print job #%v", job.JobId)
	ctx.JSON(http.StatusOK, job)
}

func PrinterService() {
	router := gin.Default()
	router.POST("/print-jobs", handlePrintJobs)
	router.Run(":4000")
}
