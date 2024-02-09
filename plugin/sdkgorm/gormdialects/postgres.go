package gormdialects

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Ex: postgresql://username:password@localhost:3306/db_name?sslmode=disable
func PostgresDB(uri string) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(uri))
}
