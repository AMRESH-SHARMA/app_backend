package testdb

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Listener struct {
	ID       uint `gorm:"primaryKey"`
	Name     string
	Language string
	Gender   string
	AgeGroup string
}

// SetupInMemoryDB initializes an in-memory SQLite DB and seeds dummy listeners
var defaultListeners = []Listener{
	{Name: "Alice", Language: "English", Gender: "female", AgeGroup: "18-25"},
	{Name: "Bob", Language: "Hindi", Gender: "male", AgeGroup: "26-35"},
	{Name: "Charlie", Language: "English", Gender: "male", AgeGroup: "18-25"},
	{Name: "Diana", Language: "French", Gender: "female", AgeGroup: "36-45"},
}

func SetupInMemoryDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if err := db.AutoMigrate(&Listener{}); err != nil {
		return nil, err
	}
	for _, l := range defaultListeners {
		db.Create(&l)
	}
	log.Printf("DB DONE")
	fmt.Printf("DB DONE")
	return db, nil
}
