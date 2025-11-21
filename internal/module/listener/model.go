package listener

import "gorm.io/gorm"

type Listener struct {
	ID       int    `json:"id"`
	Name     string `gorm:"size:100" json:"name"`
	Phone    string `gorm:"unique" json:"phone"`
	Email    string `gorm:"unique" json:"email"`
	Avatar   string `json:"avatar"`
	Gender   string `json:"gender"`
	Language string `json:"language"`
	AboutMe  string `gorm:"type:text" json:"about_me"`
	gorm.Model
}
