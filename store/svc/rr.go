package store

type ReserveProductRequest struct {
	ProductID int
}

type ReserveProductResponse struct {
	Reserved bool
}

type BookProductRequest struct {
	OrderID string
	ProductID  int
}

type BookProductResponse struct {
	OrderID string
}