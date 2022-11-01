package store

import "database/sql"

type Packet struct {
	ID         int
	ProductID     int
	IsReserved bool
	OrderID    sql.NullString
}