package listener

import (
	"app_backend/internal/testdb"

	"gorm.io/gorm"
)

// SearchListeners performs a global search by name, language, or ID
func SearchListeners(db *gorm.DB, query string) ([]testdb.Listener, error) {
	var listeners []testdb.Listener
	err := db.Where("name LIKE ? OR language LIKE ? OR id = ?", "%"+query+"%", "%"+query+"%", query).Find(&listeners).Error
	return listeners, err
}

// CustomSearchListeners searches by language, gender, and age group
func CustomSearchListeners(db *gorm.DB, language, gender, ageGroup string) ([]testdb.Listener, error) {
	var listeners []testdb.Listener
	err := db.Where("language = ? AND gender = ? AND age_group = ?", language, gender, ageGroup).Find(&listeners).Error
	return listeners, err
}
