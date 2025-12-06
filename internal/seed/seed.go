package seed

import (
	"app_backend/internal/module/listener"
	"app_backend/internal/module/user"
	"fmt"

	"gorm.io/gorm"
)

func Run(db *gorm.DB) {
	defer fmt.Println("✅ Seeded")
	db.Create(&users)
	db.Create(&listeners)

}

var users = []user.User{
	{ID: 1, Name: "Alice", Email: "alice@g.com", Phone: "0123456789", Gender: "female", Language: "English", Role: "CUSTOMER"},
	{ID: 1, Name: "Bob", Email: "bob@g.com", Phone: "1234567890", Gender: "male", Language: "Hindi", Role: "CUSTOMER"},
	{ID: 1, Name: "Charlie", Email: "charlie@g.com", Phone: "0113456789", Gender: "male", Language: "English", Role: "CUSTOMER"},
	{ID: 1, Name: "Diana", Email: "diana@g.com", Phone: "1113456789", Gender: "female", Language: "French", Role: "CUSTOMER"},
}

var listeners = []listener.Listener{
	{
		UserID:      1,
		Avatar:      "https://i.pravatar.cc/150?img=48",
		Bio:         "Professional listener for life guidance.",
		About:       "Hello, I'm Alice.",
		Experience:  3,
		PricePerMin: 2.50,
		Languages:   "English",
		Rating:      4.5,
	},
	{
		UserID:      2,
		Avatar:      "https://i.pravatar.cc/150?img=48",
		Bio:         "Experienced listener specializing in mindset coaching.",
		About:       "Hi, I'm Bob.",
		Experience:  5,
		PricePerMin: 3.75,
		Languages:   "Hindi",
		Rating:      4.2,
	},
	{
		UserID:      3,
		Avatar:      "https://i.pravatar.cc/150?img=48",
		Bio:         "I love helping people through conversations.",
		About:       "Hey there, I'm Charlie.",
		Experience:  4,
		PricePerMin: 4.99,
		Languages:   "English",
		Rating:      4.0,
	},
	{
		UserID:      4,
		Avatar:      "https://i.pravatar.cc/150?img=48",
		Bio:         "Multilingual listener with international experience.",
		About:       "Bonjour, I'm Diana.",
		Experience:  6,
		PricePerMin: 5.50,
		Languages:   "French",
		Rating:      4.8,
	},
}
