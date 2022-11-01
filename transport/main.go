package main


import (
	"log"

	"github.com/gin-gonic/gin"
	transport "github.com/prajwalcr/DS_Project_E-commerce/transport/svc"
	"github.com/prajwalcr/DS_Project_E-commerce/io"
)

func main() {
	io.Connect()
	transport.Clean()
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.POST("/transport/vehicle/reserve", func(c *gin.Context) {

		vehicle, err := transport.ReserveVehicle()
		if err != nil {
			c.AbortWithError(500, err)
			return
		}

		c.JSON(200, transport.ReserveVehicleResponse{VehicleID: vehicle.ID})
	})

	r.POST("/transport/vehicle/book", func(c *gin.Context) {

		var req transport.BookVehicleRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.AbortWithStatus(400)
		}

		vehicle, err := transport.BookVehicle(req.OrderID)
		if err != nil {
			c.AbortWithError(500, err)
			return
		}

		c.JSON(200, transport.BookVehicleResponse{VehicleID: vehicle.ID, OrderID: req.OrderID})
	})
	log.Println("running the transport service on port 8083")
	r.Run(":8083")
}