package orders

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/google/uuid"
)

func PlaceOrder(productID int) (*Order, error) {
	// reserve product
	body, _ := json.Marshal(map[string]interface{}{
		"ProductID": productID,
	})

	reqBody := bytes.NewBuffer(body)

	resp1, err := http.Post("http://localhost:8081/store/product/reserve", "application/json", reqBody)
	if err != nil {
		return nil, err
	}
	if resp1.StatusCode != 200 {
		return nil, errors.New("product not available")
	}

	// reserve agent
	resp2, err := http.Post("http://localhost:8082/delivery/agent/reserve", "application/json", nil)
	if err != nil {
		return nil, err
	}
	if resp2.StatusCode != 200 {
		return nil, errors.New("delivery agent not availabe")
	}

	// reserve vehicle
	resp5, err := http.Post("http://localhost:8083/transport/vehicle/reserve", "application/json", nil)
	if err != nil {
		return nil, err
	}
	if resp5.StatusCode != 200 {
		return nil, errors.New("transport vehicle not availabe")
	}

	orderID := uuid.New().String()

	// book product
	body, _ = json.Marshal(map[string]interface{}{
		"OrderID": orderID,
		"ProductID":  productID,
	})

	reqBody = bytes.NewBuffer(body)
	resp3, err := http.Post("http://localhost:8081/store/product/book", "application/json", reqBody)
	
	if err != nil {
		return nil, err
	}
	if resp3.StatusCode != 200 {
		return nil, errors.New("could not assign product to an order")
	}

	// book agent
	body, _ = json.Marshal(map[string]interface{}{
		"OrderID": orderID,
	})

	reqBody = bytes.NewBuffer(body)
	resp4, err := http.Post("http://localhost:8082/delivery/agent/book", "application/json", reqBody)

	if err != nil {
		return nil, err
	}
	if resp4.StatusCode != 200 {
		return nil, errors.New("could not assign delivery agent to an order")
	}

	// book vehicle
	body, _ = json.Marshal(map[string]interface{}{
		"OrderID": orderID,
	})

	reqBody = bytes.NewBuffer(body)
	resp6, err := http.Post("http://localhost:8083/transport/vehicle/book", "application/json", reqBody)

	if err != nil {
		return nil, err
	}
	if resp6.StatusCode != 200 {
		return nil, errors.New("could not assign transport vehicle to an order")
	}

	return &Order{ID: orderID}, nil

}
