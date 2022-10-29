package main

import (
	"log"
	"sync"
	"time"

	orders "github.com/prajwalcr/DS_Project_E-commerce/orders/svc"
)

func main() {
	foodID := 1
	var wg sync.WaitGroup

	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			order, err := orders.PlaceOrder(foodID)
			time.Sleep(1)
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
