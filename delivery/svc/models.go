package delivery

import "database/sql"

type Agent struct {
	ID         int
	IsReserved bool
	OrderID    sql.NullString
}