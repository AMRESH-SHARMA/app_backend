package chat

import (
	"gorm.io/gorm"
)

type ChatSession struct {
	gorm.Model
	SessionID  string `gorm:"uniqueIndex;not null" json:"sessionId"`
	CustomerID int64  `json:"customerId"`
	ListenerID int64  `json:"listenerId"`
}

type Message struct {
	gorm.Model
	SessionID       string `json:"sessionId"`
	SenderID        int64  `json:"senderId"`
	ReceiverID      int64  `json:"receiverId"`
	Type            string `json:"type"` // text | image | audio
	Content         string `gorm:"type:text" json:"content"`
	MediaUrl        string `json:"mediaUrl"`
	Status          string `json:"status"` // sent | delivered | seen
	ClientMessageID string `json:"clientMessageId"`
}
