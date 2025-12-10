package rtm

import "app_backend/internal/database"

func SaveRTMMessage(msg *RTMMessage) error {
	return database.DB.Create(msg).Error
}

func GetRTMMessages(channelID string, limit int) ([]RTMMessage, error) {
	var msgs []RTMMessage
	err := database.DB.
		Where("channel_id = ?", channelID).
		Order("created_at asc").
		Limit(limit).
		Find(&msgs).Error

	return msgs, err
}

func UpdateRTMStatus(messageID string, status string) error {
	return database.DB.
		Model(&RTMMessage{}).
		Where("message_id = ?", messageID).
		Update("status", status).Error
}
