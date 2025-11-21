package listener

import (
	"app_backend/internal/database"
)

func GetAllListenerR() ([]Listener, error) {
	var listeners []Listener
	err := database.DB.Find(&listeners).Error
	return listeners, err
}

func GlobalSearchR(accountID, name, gender, lang string) ([]Listener, error) {
	var listeners []Listener
	query := database.DB.Model(&Listener{})

	if accountID != "" {
		query = query.Where("account_id = ?", accountID)
	}
	if name != "" {
		query = query.Where("name ILIKE ?", "%"+name+"%")
	}
	if gender != "" {
		query = query.Where("gender ILIKE ?", "%"+gender+"%")
	}
	if lang != "" {
		query = query.Where("language ILIKE ?", "%"+lang+"%")
	}

	err := query.Find(&listeners).Error
	return listeners, err
}

// // CustomSearchListeners searches by language, gender, and age group
// func CustomSearchListeners(db *gorm.DB, language, gender, ageGroup string) ([]testdb.Listener, error) {
// 	var listeners []testdb.Listener
// 	err := db.Where("language = ? AND gender = ? AND age_group = ?", language, gender, ageGroup).Find(&listeners).Error
// 	return listeners, err
// }
