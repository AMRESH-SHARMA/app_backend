package rtc

type TokenRequest struct {
	Channel string `json:"channel"`
	Uid     string `json:"uid"`
}

type TokenResponse struct {
	Token string `json:"token"`
}
