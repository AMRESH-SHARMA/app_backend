package call

import (
	"app_backend/internal/database"
	"app_backend/internal/module/notification"
	"app_backend/internal/module/user"
	"app_backend/internal/response"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func StartCall(c *fiber.Ctx) error {
	req := new(StartCallRequest)
	if err := c.BodyParser(req); err != nil {
		return response.Error(c, "Invalid request", fiber.StatusBadRequest)
	}

	callID := uuid.New().String()

	Calls[callID] = &Call{
		ID:       callID,
		CallerID: req.CallerID,
		CalleeID: req.CalleeID,
		Status:   "RINGING",
	}

	// lookup callee in DB
	var callee user.User
	if err := database.DB.First(&callee, "account_id = ?", req.CalleeID).Error; err == nil {
		notification.SendToToken(callee.DeviceToken, map[string]string{
			"type":     "incoming_call",
			"callId":   callID,
			"callerId": req.CallerID,
		})
	}

	return response.Success(c, fiber.Map{
		"callId": callID,
		"status": "RINGING",
	}, "Call started", fiber.StatusOK)
}

func AcceptCall(c *fiber.Ctx) error {
	req := new(AcceptCallRequest)
	if err := c.BodyParser(req); err != nil {
		return response.Error(c, "Invalid request", fiber.StatusBadRequest)
	}

	call, ok := Calls[req.CallID]
	if !ok {
		return response.Error(c, "Call not found", fiber.StatusNotFound)
	}

	call.Status = "ACCEPTED"

	return response.Success(c, fiber.Map{
		"callId":  call.ID,
		"status":  call.Status,
		"channel": call.Channel,
	}, "Call accepted", fiber.StatusOK)
}

func RejectCall(c *fiber.Ctx) error {
	req := new(RejectCallRequest)
	if err := c.BodyParser(req); err != nil {
		return response.Error(c, "Invalid request", fiber.StatusBadRequest)
	}

	call, ok := Calls[req.CallID]
	if !ok {
		return response.Error(c, "Call not found", fiber.StatusNotFound)
	}

	call.Status = "REJECTED"

	return response.Success(c, fiber.Map{
		"callId": call.ID,
		"status": call.Status,
	}, "Call rejected", fiber.StatusOK)
}
