package chat

import (
	"app_backend/internal/database"
	"math/rand"
)

func CreateOrGetSessionR(customerID, listenerID int64) (*ChatSession, error) {
	var session ChatSession

	// check existing session
	err := database.DB.
		Where("customer_id = ? AND listener_id = ?", customerID, listenerID).
		First(&session).Error

	if err == nil {
		return &session, nil
	}

	// create new session
	session = ChatSession{
		SessionID:  RandString(20), // now generating real random ID
		CustomerID: customerID,
		ListenerID: listenerID,
	}

	if err := database.DB.Create(&session).Error; err != nil {
		return nil, err
	}

	return &session, nil
}

func SaveMessageR(msg *Message) error {
	return database.DB.Create(msg).Error
}

func GetMessagesR(sessionID string, limit int, before string) ([]Message, error) {
	var msgs []Message
	q := database.DB.Where("session_id = ?", sessionID)

	if before != "" {
		q = q.Where("created_at < ?", before)
	}

	err := q.Order("created_at desc").Limit(limit).Find(&msgs).Error
	return msgs, err
}

func UpdateMessageStatusR(messageID string, status string) error {
	return database.DB.
		Model(&Message{}).
		Where("id = ?", messageID).
		Update("status", status).
		Error
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func RandString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
