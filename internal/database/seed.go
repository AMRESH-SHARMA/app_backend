package database

import (
	"app_backend/internal/module/listener"
	"fmt"

	"gorm.io/gorm"
)

var defaultListeners = []listener.Listener{
	{Name: "Alice", Email: "a@g.com", Avatar: "https://i.pravatar.cc/150?img=48", Phone: "1234567890", Language: "English", Gender: "female"},
	{Name: "Bob", Email: "b@g.com", Avatar: "https://i.pravatar.cc/150?img=48", Phone: "2234567890", Language: "Hindi", Gender: "male"},
	{Name: "Charlie", Email: "c@g.com", Avatar: "https://i.pravatar.cc/150?img=48", Phone: "3234567890", Language: "English", Gender: "male"},
	{Name: "Diana", Email: "d@g.com", Avatar: "https://i.pravatar.cc/150?img=48", Phone: "4234567890", Language: "French", Gender: "female"},
}

func seed(db *gorm.DB) {
	var count int64
	db.Model(&listener.Listener{}).Count(&count)
	if count == 0 {
		db.Create(&defaultListeners)
		fmt.Println("✅ Seeded listeners")
	} else {
		fmt.Println("⚠️ Listeners already seeded")
	}

}
