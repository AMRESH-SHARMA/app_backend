package listener

type Listener struct {
	UserID      uint    `json:"userId" gorm:"primaryKey"`
	Age         int     `json:"age"`
	Avatar      string  `json:"avatar"`
	TagLine     string  `json:"TagLine"`
	About       string  `gorm:"type:text" json:"about"`
	Experience  int     `json:"experience"`
	PricePerMin float64 `json:"pricePerMin"`
	Rating      float32 `json:"rating"`
}
