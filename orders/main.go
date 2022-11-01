package main

import (
	"log"
	"sync"

	orders "github.com/prajwalcr/DS_Project_E-commerce/orders/svc"
)

func main() {
	foodID := 1
	var wg sync.WaitGroup

	numberOfOrders := 10

	wg.Add(numberOfOrders)

	for i := 0; i < numberOfOrders; i++ {
		go func() {
			order, err := orders.PlaceOrder(foodID)
			if err != nil {
				log.Println("order not placed : ", err.Error())
			} else {
				log.Println("order placed : ", order.ID)
			}
			wg.Done()
		}()
	}

	wg.Wait()
}
