package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

type PrintJob struct {
	JobId     int    `json:"jobId" binding:"gte=0"`
	InvoiceId int    `json:"invoiceId" binding:"required,gte=0"`
	Format    string `json:"format" binding:"required"`
}

type Invoice struct {
	InvoiceId   int    `json:"invoiceId"`
	CustomerId  int    `json:"customerId" binding:"required,gte=0"`
	Price       int    `json:"price" binding:"required,gte=0"`
	Description string `json:"description" binding:"required"`
}

func createPrintJob(invoiceId int) {
	client := resty.New()
	var pj PrintJob
	// Call PrinterService via RESTful interface
	_, err := client.R().
		SetBody(PrintJob{Format: "A4", InvoiceId: invoiceId}).
		SetResult(&pj).
		Post("http://localhost:4000/print-jobs")

	if err != nil {
		log.Println("InvoiceGenerator: unable to connect PrinterService")
		return
	}
	log.Printf("InvoiceGenerator: created print job #%v via PrinterService", pj.JobId)
}

func handleInvoiceGeneration(ctx *gin.Context) {
	var invoice Invoice
	if err := ctx.ShouldBindJSON(&invoice); err != nil {
		ctx.JSON(400, "Invalid input")
		return
	}
	log.Printf("InvoiceService: Creating new invoice...")
	rand.Seed(time.Now().UnixNano())
	invoice.InvoiceId = rand.Intn(1000)
	log.Printf("InvoiceService: Created new print job #%v", invoice.InvoiceId)
	ctx.JSON(http.StatusOK, invoice)
}

func main() {
	// Set "release" mode
	// gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	router.POST("/invoices", handleInvoiceGeneration)
	router.Run(":5000")
}
