package delivery

type ReserveAgentResponse struct {
	AgentID int
}

type BookAgentRequest struct {
	OrderID sql.NullString
}

type BookAgentResponse struct {
	AgentID int
	OrderID sql.NullString
}