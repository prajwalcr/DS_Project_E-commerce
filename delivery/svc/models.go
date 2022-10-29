package delivery

type Agent struct {
	ID int
	IsReserved bool
	OrderID sql.NullString
}