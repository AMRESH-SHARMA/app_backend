package seed

import (
	"app_backend/internal/module/listener"
	"app_backend/internal/module/user"
	"fmt"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

func Run(db *gorm.DB) {
	defer fmt.Println("✅ Seeded")
	db.Create(&users)
	db.Create(&listeners)

}

var users = []user.User{
	{AccountID: 1, Name: "Alice", Email: "alice@g.com", Phone: "0123456789", Gender: "F", Languages: pq.StringArray{"English", "Hindi"}, Role: "CUSTOMER"},
	{AccountID: 2, Name: "Bob", Email: "bob@g.com", Phone: "1234567890", Gender: "M", Languages: pq.StringArray{"English", "Hindi"}, Role: "LISTENER"},
	{AccountID: 3, Name: "Charlie", Email: "charlie@g.com", Phone: "0113456789", Gender: "M", Languages: pq.StringArray{"English", "Hindi"}, Role: "LISTENER"},
	{AccountID: 4, Name: "Diana", Email: "diana@g.com", Phone: "1113456789", Gender: "F", Languages: pq.StringArray{"English", "Hindi"}, Role: "LISTENER"},
}

var listeners = []listener.Listener{
	{UserID: users[0].ID, Age: 20, Avatar: "https://i.pravatar.cc/150?img=48", TagLine: "Professional listener for life guidance.", About: "Hello, I'm Alice.", Experience: 3, PricePerMin: 2.50, Rating: 4.5},
	{UserID: users[1].ID, Age: 20, Avatar: "https://i.pravatar.cc/150?img=48", TagLine: "Experienced listener specializing in mindset coaching.", About: "Hi, I'm Bob.", Experience: 5, PricePerMin: 3.75, Rating: 4.2},
	{UserID: users[2].ID, Age: 20, Avatar: "https://i.pravatar.cc/150?img=48", TagLine: "I love helping people through conversations.", About: "Hey there, I'm Charlie.", Experience: 4, PricePerMin: 4.99, Rating: 4.0},
	{UserID: users[3].ID, Age: 20, Avatar: "https://i.pravatar.cc/150?img=48", TagLine: "Multilingual listener with international experience.", About: "Bonjour, I'm Diana.", Experience: 6, PricePerMin: 5.50, Rating: 4.8},
}
