package chat

type CreateSessionRequest struct {
	CustomerID int64 `json:"customerId"`
	ListenerID int64 `json:"listenerId"`
}

type SendMessageRequest struct {
	SessionID       string `json:"sessionId"`
	SenderID        int64  `json:"senderId"`
	ReceiverID      int64  `json:"receiverId"`
	Type            string `json:"type"`
	Content         string `json:"content"`
	MediaUrl        string `json:"mediaUrl"`
	ClientMessageID string `json:"clientMessageId"`
}

type UpdateStatusRequest struct {
	MessageID string `json:"messageId"`
	SessionID string `json:"sessionId"`
	Status    string `json:"status"` // delivered | seen
}
