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

	err = db.AutoMigrate(models...)
	if err != nil {
		panic("Database migration failed")
	}
	return db
}
