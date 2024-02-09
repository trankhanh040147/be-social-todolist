package gormdialects

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Ex: /tmp/gorm.db
func SQLiteDB(uri string) (db *gorm.DB, err error) {
	return gorm.Open(sqlite.Open(uri))
}