package store

type ReserveFoodRequest struct {
	FoodID int
}

type ReserveFoodResponse struct {
	Reserved bool
}

type BookFoodRequest struct {
	OrderID string
	FoodID  int
}

type BookFoodResponse struct {
	OrderID string
}