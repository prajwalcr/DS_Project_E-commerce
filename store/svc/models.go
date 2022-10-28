package store

import "database/sql"

type Packet struct {
	ID         int
	FoodID     int
	IsReserved bool
	OrderID    sql.NullString
}