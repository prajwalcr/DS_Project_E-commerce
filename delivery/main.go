package main

import (
	"log"

	"github.com/gin-gonic/gin"
	delivery "github.com/prajwalcr/DS_Project_E-commerce/delivery/svc"
	"github.com/prajwalcr/DS_Project_E-commerce/io"
)

func main() {
	io.Connect()
	delivery.Clean()
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.POST("/delivery/agent/reserve", func(c *gin.Context) {

		agent, err := delivery.ReserveAgent()
		if err != nil {
			c.AbortWithError(500, err)
			return
		}

		c.JSON(200, delivery.ReserveAgentResponse{AgentID: agent.ID})
	})

	r.POST("/delivery/agent/book", func(c *gin.Context) {

		var req delivery.BookAgentRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.AbortWithStatus(400)
		}

		agent, err := delivery.BookAgent(req.OrderID)
		if err != nil {
			c.AbortWithError(500, err)
			return
		}

		c.JSON(200, delivery.BookAgentResponse{AgentID: agent.ID, OrderID: req.OrderID})
	})
	log.Println("running the delivery service on port 8082")
	r.Run(":8082")
}
