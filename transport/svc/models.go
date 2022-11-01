package transport

import "database/sql"

type Vehicle struct {
	ID         int
	IsReserved bool
	OrderID    sql.NullString
}