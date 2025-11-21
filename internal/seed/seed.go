package seed

import (
	"app_backend/internal/module/listener"
	"app_backend/internal/module/user"
	"fmt"

	"gorm.io/gorm"
)

func Run(db *gorm.DB) {
	defer fmt.Println("✅ Seeded")
	db.Create(&listeners)
	db.Create(&users)
}

var listeners = []listener.Listener{
	{Name: "Alice", Phone: "1234567890", Email: "a@g.com", Avatar: "https://i.pravatar.cc/150?img=48", Gender: "female", Language: "English", AboutMe: "Hello, I'm Alice."},
	{Name: "Bob", Phone: "2234567890", Email: "b@g.com", Avatar: "https://i.pravatar.cc/150?img=48", Gender: "male", Language: "Hindi", AboutMe: "Hi, I'm Bob."},
	{Name: "Charlie", Phone: "3234567890", Email: "c@g.com", Avatar: "https://i.pravatar.cc/150?img=48", Gender: "male", Language: "English", AboutMe: "Hey there, I'm Charlie."},
	{Name: "Diana", Phone: "4234567890", Email: "d@g.com", Avatar: "https://i.pravatar.cc/150?img=48", Gender: "female", Language: "French", AboutMe: "Bonjour, I'm Diana."},
}

var users = []user.User{
	{Name: "Alice", Phone: "1234567890", Gender: "female", Language: "English"},
	{Name: "Bob", Phone: "2234567890", Gender: "male", Language: "Hindi"},
	{Name: "Charlie", Phone: "3234567890", Gender: "male", Language: "English"},
	{Name: "Diana", Phone: "4234567890", Gender: "female", Language: "French"},
}
