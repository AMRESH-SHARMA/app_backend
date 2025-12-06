package listener

import "gorm.io/gorm"

type Listener struct {
	ID          uint    `json:"id" gorm:"primaryKey"`
	UserID      uint    `json:"userId" gorm:"unique"`
	Avatar      string  `json:"avatar"`
	Bio         string  `json:"bio"`
	About       string  `gorm:"type:text" json:"about"`
	Experience  int     `json:"experience"`
	PricePerMin float64 `json:"pricePerMin"`
	Languages   string  `json:"languages"`
	Rating      float32 `json:"rating"`
	gorm.Model
}
