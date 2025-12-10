package rtm

import "fmt"

func SendRTMMessageService(req SendRTMRequest) (*RTMMessage, error) {

	msg := RTMMessage{
		MessageID:  fmt.Sprintf("msg_%d", req.SenderID), // generate UUID later
		SenderID:   req.SenderID,
		ReceiverID: req.ReceiverID,
		ChannelID:  req.ChannelID,
		Content:    req.Content,
		Type:       req.Type,
		Status:     "sent",
	}

	if err := SaveRTMMessage(&msg); err != nil {
		return nil, err
	}

	// TODO: call Agora Chat SDK/REST to deliver realtime message

	return &msg, nil
}
