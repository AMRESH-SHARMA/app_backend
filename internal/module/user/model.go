package user

import (
	"math/rand/v2"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	AccountID   int64          `gorm:"uniqueIndex;not null" json:"accountId"`
	Name        string         `json:"name"`
	Email       string         `json:"email" gorm:"unique"`
	Password    string         `json:"-"` // hashed
	Phone       string         `gorm:"unique" json:"phone"`
	Gender      string         `json:"gender"`
	Languages   pq.StringArray `json:"languages" gorm:"type:text[]"`
	Role        string         `json:"role"` // LISTENER or CUSTOMER
	DeviceToken string         `json:"deviceToken"`

	Listener ListenerProfile `json:"listener" gorm:"foreignKey:UserID"` // OK

}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	for {
		id := int64(rand.Uint32()) // convert here
		var count int64
		tx.Model(&User{}).Where("account_id = ?", id).Count(&count)
		if count == 0 {
			u.AccountID = id // assign int64
			break
		}
	}
	return
}

type ListenerProfile struct {
	UserID      uint    `json:"userId" gorm:"primaryKey"`
	Age         int     `json:"age"`
	Avatar      string  `json:"avatar"`
	TagLine     string  `json:"TagLine"`
	About       string  `gorm:"type:text" json:"about"`
	Experience  int     `json:"experience"`
	PricePerMin float64 `json:"pricePerMin"`
	Rating      float32 `json:"rating"`
}

func (ListenerProfile) TableName() string {
	return "listeners"
}
