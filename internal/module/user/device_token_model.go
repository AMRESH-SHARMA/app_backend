package user

type DeviceTokenRequest struct {
	AccountID   int64  `json:"accountId"`
	DeviceToken string `json:"deviceToken"`
}

// type RefreshDeviceTokenRequest struct {
// 	UserID   string `json:"userId"`
// 	OldToken string `json:"oldToken"`
// 	NewToken string `json:"newToken"`
// }
