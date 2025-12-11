package call

import (
	"fmt"
	"time"

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

	if req.CallerID == 0 || req.CalleeID == 0 {
		return response.Error(c, "callerId and calleeId are required", fiber.StatusBadRequest)
	}

	callID := uuid.New().String()
	now := time.Now().Unix()

	call := Call{
		CallID:    callID,       // CallId is ChannelName custom generate by backend
		CallerID:  req.CallerID, // from android
		CalleeID:  req.CalleeID, // from android
		Status:    CallStatusRequested,
		StartedAt: now, // backend
	}

	if err := database.DB.Create(&call).Error; err != nil {
		return response.Error(c, "DB error", fiber.StatusInternalServerError)
	}

	// lookup listener device
	var callee user.User
	if err := database.DB.First(&callee, "account_id = ?", req.CalleeID).Error; err == nil {
		notification.SendToToken(callee.DeviceToken, map[string]string{
			"type":     "incoming_call",
			"callId":   callID,
			"callerId": fmt.Sprintf("%d", req.CallerID),
			"calleeId": fmt.Sprintf("%d", req.CalleeID),
		})
	}

	return response.Success(c, fiber.Map{
		"callId": callID,
		"status": call.Status,
	}, "Call started", fiber.StatusOK)
}

func AcceptCall(c *fiber.Ctx) error {
	req := new(AcceptCallRequest)
	if err := c.BodyParser(req); err != nil {
		return response.Error(c, "Invalid request", fiber.StatusBadRequest)
	}

	var call Call
	if err := database.DB.First(&call, "id = ?", req.CallID).Error; err != nil {
		return response.Error(c, "Call not found", fiber.StatusNotFound)
	}

	call.Status = CallStatusOngoing
	if err := database.DB.Save(&call).Error; err != nil {
		return response.Error(c, "DB error", fiber.StatusInternalServerError)
	}

	return response.Success(c, fiber.Map{
		"callId": call.CallID,
		"status": call.Status,
	}, "Call accepted", fiber.StatusOK)
}

func RejectCall(c *fiber.Ctx) error {
	req := new(RejectCallRequest)
	if err := c.BodyParser(req); err != nil {
		return response.Error(c, "Invalid request", 400)
	}

	var call Call
	if err := database.DB.First(&call, "call_id = ?", req.CallID).Error; err != nil {
		return response.Error(c, "Call not found", 404)
	}

	call.Status = CallStatusRejected
	call.EndedAt = time.Now().Unix()

	if err := database.DB.Save(&call).Error; err != nil {
		return response.Error(c, "DB error", 500)
	}

	// Notify caller
	var caller user.User
	if err := database.DB.First(&caller, "account_id = ?", call.CallerID).Error; err == nil {
		notification.SendToToken(caller.DeviceToken, map[string]string{
			"type":   "call_rejected",
			"callId": call.CallID,
		})
	}

	// Notify callee too (important)
	var callee user.User
	if err := database.DB.First(&callee, "account_id = ?", call.CalleeID).Error; err == nil {
		notification.SendToToken(callee.DeviceToken, map[string]string{
			"type":   "call_rejected",
			"callId": call.CallID,
		})
	}

	return response.Success(c, fiber.Map{
		"callId": call.CallID,
		"status": call.Status,
	}, "Call rejected", 200)
}
