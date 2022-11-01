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
	r.POST("/store/food/reserve", func(c *gin.Context) {
		var req store.ReserveFoodRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.AbortWithStatus(400)
			return
		}

		_, err := store.ReserveFood(req.FoodID)
		if err != nil {
			c.AbortWithError(500, err)
			return
		}

		c.JSON(200, store.ReserveFoodResponse{Reserved: true})
	})

	r.POST("/store/food/book", func(c *gin.Context) {
		var req store.BookFoodRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.AbortWithStatus(400)
			return
		}

		packet, err := store.BookFood(req.OrderID, req.FoodID)
		if err != nil {
			c.AbortWithError(500, err)
			return
		}
		c.JSON(200, store.BookFoodResponse{OrderID: packet.OrderID.String})
	})
	log.Println("running the store service on port 8081")
	r.Run(":8081")
}
