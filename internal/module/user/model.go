package user

import (
	"app_backend/internal/module/listener"

	"gorm.io/gorm"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"-"` // hashed
	Phone    string `gorm:"unique" json:"phone"`
	Gender   string `json:"gender"`
	Language string `json:"language"`
	Role     string `json:"role"` // LISTENER or CUSTOMER
	// role ENUM('listener','customer')
	DeviceToken string `json:"deviceToken"`

	Listener listener.Listener `json:"listener" gorm:"foreignKey:UserID"`
	gorm.Model
}
