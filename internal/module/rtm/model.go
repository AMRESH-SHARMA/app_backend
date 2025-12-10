package rtm

import "gorm.io/gorm"

type RTMMessage struct {
	gorm.Model

	MessageID  string `gorm:"uniqueIndex" json:"messageId"`
	SenderID   int64  `json:"senderId"`   // AccountID
	ReceiverID int64  `json:"receiverId"` // AccountID
	ChannelID  string `json:"channelId"`  // same as chat session
	Content    string `gorm:"type:text" json:"content"`
	Type       string `json:"type"`   // text | image | audio
	Status     string `json:"status"` // sent | delivered | seen
}
