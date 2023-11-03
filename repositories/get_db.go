package repositories

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func getDb(name string, models ...any) *gorm.DB {
	if len(name) == 0 {
		panic("DB name can't have 0 length!")
	}

	db, err := gorm.Open(sqlite.Open(name+".db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}

	db.AutoMigrate(models...)
	return db
}
