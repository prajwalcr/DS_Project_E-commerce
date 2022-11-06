package delivery

import "database/sql"

type Agent struct {
	ID         int
	IsReserved sql.NullTime
	OrderID    sql.NullString
}
