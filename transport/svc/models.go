package transport

import "database/sql"

type Vehicle struct {
	ID         int
	IsReserved sql.NullTime
	OrderID    sql.NullString
}
