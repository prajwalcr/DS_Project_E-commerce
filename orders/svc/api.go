package api

import 
(
	"github/google/uuid"
	"errors"
	"bytes"
	"encoding/json"
	"net/http"
)

func PlaceOrder(foodID int) (*Order, error)
{
	// reserve food
	body, _ := json.Marshal(map[string]interface{}
	{
		"food_id" : 1,
	})
	
	reqBody := bytes.NewBuffer(body)
	
	resp1, err := http.Post("http://localhost:8081/store/food/reserve", "application/json", reqBody)
	
	if err != nil || resp1.StatusCode != 200
	{
		return nil, errors.New("food not available")
	}
	
	// reserve agent
	resp2, err := http.Post("http://localhost:8082/delivery/agent/reserve", "application/json", nil)
	
	if err != nil || resp2.StatusCode != 200
	{
		return nil, errors.New("delivery agent not availabe")
	}
	
	orderID := uuid.New().String()
	
	// book food
	body, _ = json.Marshal(map[string]interface{}
	{
		"order_id" : orderID,
		"food_id" : foodID,
	})
	
	reqBody = bytes.NewBuffer(body)
	resp3, err := http.Post("http://localhost:8081/store/food/book", "application/json", reqBody)
	
	if err != nil || resp3.StatusCode != 200
	{
		return nil, errors.New("could not assign food to an order")
	}
	
	body, _ = json.Marshal(map[string]interface{}
	{
		"order_id" : orderID,
	})
	
	reqBody = bytes.NewBuffer(body)
	resp4, err := http.Post("http://localhost:8082/delivery/agent/book", "application/json", reqBody)
	
	if err != nil || resp4.StatusCode != 200
	{
		return nil, errors.New("could not assign delivery agent to an order")
	}
	
	return &Order{ID: orderID}, nil		
	
}