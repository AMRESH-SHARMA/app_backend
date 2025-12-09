package call

type Call struct {
	ID       string `json:"id"`
	CallerID string `json:"callerId"`
	CalleeID string `json:"calleeId"`
	Status   string `json:"status"`
	Channel  string `json:"channel"`
}

type StartCallRequest struct {
	CallerID string `json:"callerId"`
	CalleeID string `json:"calleeId"`
}

type AcceptCallRequest struct {
	CallID string `json:"callId"`
}

type RejectCallRequest struct {
	CallID string `json:"callId"`
}
