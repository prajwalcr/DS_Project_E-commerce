package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/prajwalcr/DS_Project_E-commerce/io"
	store "github.com/prajwalcr/DS_Project_E-commerce/store/svc"
)

func main() {
	io.Connect()
	store.Clean()
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.POST("/store/product/reserve", func(c *gin.Context) {
		var req store.ReserveProductRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.AbortWithStatus(400)
			return
		}

		_, err := store.ReserveProduct(req.ProductID)
		if err != nil {
			c.AbortWithError(500, err)
			return
		}

		c.JSON(200, store.ReserveProductResponse{Reserved: true})
	})

	r.POST("/store/product/book", func(c *gin.Context) {
		var req store.BookProductRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.AbortWithStatus(400)
			return
		}

		packet, err := store.BookProduct(req.OrderID, req.ProductID)
		if err != nil {
			c.AbortWithError(500, err)
			return
		}
		c.JSON(200, store.BookProductResponse{OrderID: packet.OrderID.String})
	})
	log.Println("running the store service on port 8081")
	r.Run(":8081")
}
