package call

import "gorm.io/gorm"

type CallStatus string

const (
	CallStatusRequested CallStatus = "requested"
	CallStatusRinging   CallStatus = "ringing"
	CallStatusOngoing   CallStatus = "ongoing"
	CallStatusCompleted CallStatus = "completed"
	CallStatusRejected  CallStatus = "rejected"
	CallStatusMissed    CallStatus = "missed"
)

type Call struct {
	gorm.Model

	CallID    string     `gorm:"uniqueIndex;type:text" json:"callId"`
	CallerID  int64      `json:"callerId"` // account id
	CalleeID  int64      `json:"calleeId"` // account id
	Status    CallStatus `json:"status"`
	Duration  int64      `json:"duration"`  // in seconds
	StartedAt int64      `json:"startedAt"` // unix ts
	EndedAt   int64      `json:"endedAt"`   // unix ts
}

type StartCallRequest struct {
	CallerID int64 `json:"callerId"`
	CalleeID int64 `json:"calleeId"`
}

type AcceptCallRequest struct {
	CallID string `json:"callId"`
}

type RejectCallRequest struct {
	CallID string `json:"callId"`
}
