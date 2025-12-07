package user

type DeviceTokenRequest struct {
	UserID      string `json:"userId"`
	DeviceToken string `json:"deviceToken"`
}

type RefreshDeviceTokenRequest struct {
	UserID   string `json:"userId"`
	OldToken string `json:"oldToken"`
	NewToken string `json:"newToken"`
}
