package user

type DeviceTokenRequest struct {
	UserID      string `json:"userId"`
	DeviceToken string `json:"deviceToken"`
}
