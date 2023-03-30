// Ref: https://blog.logrocket.com/building-microservices-go-gin/
package main

import (
	"fmt"
	"runtime"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Numbers struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type Product struct {
	Id   int    `json:"id" yaml:"id" xml:"Id"`
	Name string `json:"name" yaml:"name" xml:"Name"`
}

func v2endpointHandler(ctx *gin.Context) {
	ctx.String(200, "%s %s", ctx.Request.Method, ctx.Request.URL)
}

func getAddHandler(ctx *gin.Context) {
	x, _ := strconv.ParseFloat(ctx.Param("x"), 64)
	y, _ := strconv.ParseFloat(ctx.Param("y"), 64)
	ctx.String(200, fmt.Sprintf("%f", x+y))
}

func postAddHandler(ctx *gin.Context) {
	var nums Numbers
	if err := ctx.ShouldBind(&nums); err != nil {
		ctx.JSON(400, gin.H{"error": "Calculation Error"})
		return
	}
	ctx.JSON(200, gin.H{"sum": nums.X + nums.Y})
}

func getProductsHandler(ctx *gin.Context) {
	format := ctx.Param("format")

	switch format {
	case "json":
		product := Product{100, "Apple"}
		ctx.JSON(200, product)
		break
	case "yaml":
		product := Product{200, "Apple"}
		ctx.YAML(200, product)
		break
	case "xml":
		product := Product{300, "Apple"}
		ctx.XML(200, product)
		break
	default:
		ctx.String(200, fmt.Sprintf("Invalid format \"%s\"", format))
	}
}

func GettingStarted() {
	router := gin.Default()

	router.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello world!",
		})
	})

	// Route group
	v1 := router.Group("/v1")
	v1.GET("/os", func(ctx *gin.Context) {
		ctx.String(200, runtime.GOOS)
	})

	// Structuring microservices with routes
	// Accepting data via URL params
	// Route group
	v2 := router.Group("/v2")
	v2.GET("/products", v2endpointHandler)
	v2.GET("/products/:productId", v2endpointHandler)
	v2.POST("/products", v2endpointHandler)
	v2.PUT("/products/:productId", v2endpointHandler)
	v2.DELETE("/products/:productId", v2endpointHandler)

	v3 := router.Group("/v3")
	v3.GET("/add/:x/:y", getAddHandler)
	v3.POST("/add", postAddHandler)

	// Returning data in JSON, YAML, and XML formats
	v4 := router.Group("/v4")
	// http://localhost:5000/v4/product.[json|xml|yaml]
	v4.GET("/product.:format", getProductsHandler)

	// Validating incoming requests
	v5 := router.Group("/v5")
	// http://localhost:5000/v4/product.[json|xml|yaml]
	v5.GET("/product.:format", getProductsHandler)

	router.Run(":5000")
}
