package delivery

type ReserveAgentResponse struct {
	AgentID int
}

type BookAgentRequest struct {
	OrderID string
}

type BookAgentResponse struct {
	AgentID int
	OrderID string
}