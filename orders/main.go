package main

import 
(
	"log"
	"sync"
	// This path has to be changed later
	orders "/home/rithvik/DS_Project_E-commerce/orders/svc"
)


func main()
{
	foodID := 1
	var wg sync.WaitGroup
	
	wg.Add(10)
	
	for i := 0; i < 10; i++
	{
		go func()
		{
			order, err := orders.PlaceOrder(foodID)
			wg.Done()
			if err != nil
			{
				log.Println("order not placed : ", err.Error())
			}
			else
			{
				log.Println("order placed : ", order.ID);	
			}
		}()
	}
	
	wg.Wait()
}
