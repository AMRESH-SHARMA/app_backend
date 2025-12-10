package chat

import (
	"app_backend/internal/response"

	"github.com/gofiber/fiber/v2"
)

/*
SESSION HANDLER
*/
func CreateSession(c *fiber.Ctx) error {
	req := new(CreateSessionRequest)

	if err := c.BodyParser(req); err != nil {
		return response.Error(c, "Invalid request body", fiber.StatusBadRequest)
	}

	session, err := CreateOrGetSessionR(req.CustomerID, req.ListenerID)
	if err != nil {
		return response.Error(c, "Failed to create or fetch session", fiber.StatusInternalServerError)
	}

	return response.Success(c, session, "Session ready", fiber.StatusOK)
}

/*
SEND MESSAGE HANDLER
*/
func SendMessage(c *fiber.Ctx) error {
	req := new(SendMessageRequest)

	if err := c.BodyParser(req); err != nil {
		return response.Error(c, "Invalid request body", fiber.StatusBadRequest)
	}

	msg := Message{
		SessionID:       req.SessionID,
		SenderID:        req.SenderID,   // already int64
		ReceiverID:      req.ReceiverID, // already int64
		Type:            req.Type,
		Content:         req.Content,
		MediaUrl:        req.MediaUrl,
		Status:          "sent",
		ClientMessageID: req.ClientMessageID,
	}

	if err := SaveMessageR(&msg); err != nil {
		return response.Error(c, "Failed to send message", fiber.StatusInternalServerError)
	}

	return response.Success(c, msg, "Message sent", fiber.StatusOK)
}

/*
GET MESSAGES
*/
func GetMessages(c *fiber.Ctx) error {
	sessionID := c.Query("sessionId")
	limit := c.QueryInt("limit", 50)
	before := c.Query("before")

	msgs, err := GetMessagesR(sessionID, limit, before)
	if err != nil {
		return response.Error(c, "Failed to load messages", fiber.StatusInternalServerError)
	}

	return response.Success(c, msgs, "Messages loaded", fiber.StatusOK)
}

/*
MESSAGE STATUS UPDATE
*/
func UpdateMessageStatus(c *fiber.Ctx) error {
	req := new(UpdateStatusRequest)

	if err := c.BodyParser(req); err != nil {
		return response.Error(c, "Invalid request body", fiber.StatusBadRequest)
	}

	if err := UpdateMessageStatusR(req.MessageID, req.Status); err != nil {
		return response.Error(c, "Failed to update status", fiber.StatusInternalServerError)
	}

	return response.Success(c, fiber.Map{"status": "ok"}, "Status updated", fiber.StatusOK)
}
