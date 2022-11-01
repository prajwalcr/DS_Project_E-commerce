package transport

type ReserveVehicleResponse struct {
	VehicleID int
}

type BookVehicleRequest struct {
	OrderID string
}

type BookVehicleResponse struct {
	VehicleID int
	OrderID string
}