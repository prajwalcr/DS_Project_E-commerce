package store

import "database/sql"

type Packet struct {
	ID         int
	ProductID  int
	IsReserved sql.NullTime
	OrderID    sql.NullString
}
