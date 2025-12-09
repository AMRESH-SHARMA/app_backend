package listener

import "github.com/lib/pq"

type ListenerGetResponse struct {
	AccountID   string         `json:"accountId"`
	Name        string         `json:"name"`
	Age         int            `json:"age"`
	Gender      string         `json:"gender"`
	Avatar      string         `json:"avatar"`
	TagLine     string         `json:"tagLine"`
	About       string         `json:"about"`
	Experience  int            `json:"experience"`
	PricePerMin float64        `json:"pricePerMin"`
	Languages   pq.StringArray `json:"languages" gorm:"type:text[]"`
	Rating      float32        `json:"rating"`
}
