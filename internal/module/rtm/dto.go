package rtm

type SendRTMRequest struct {
	SenderID   int64  `json:"senderId"`
	ReceiverID int64  `json:"receiverId"`
	ChannelID  string `json:"channelId"`
	Type       string `json:"type"`
	Content    string `json:"content"`
}

type UpdateRTMStatusRequest struct {
	MessageID string `json:"messageId"`
	Status    string `json:"status"`
}

type RTMTokenRequest struct {
	UserID  string `json:"userId"`
	Channel string `json:"channel"`
}

type RTMTokenResponse struct {
	Token string `json:"token"`
}
