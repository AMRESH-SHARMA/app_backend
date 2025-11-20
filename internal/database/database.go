package database

import (
	"app_backend/internal/module/listener"
	"fmt"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := viper.GetString("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Failed to connect to database:")
	}
	DB = db

	fmt.Println("✅ Database connection established")

	if err := db.AutoMigrate(&listener.Listener{}); err != nil {
		log.Fatal("❌ Migration failed:", err)
	} else {
		fmt.Println("✅ Migration done")
	}

	seed(db)
}
