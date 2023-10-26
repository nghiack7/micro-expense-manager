/*
This service use postgresql to save and manage database connections
GORM is orm used to manage database connections and prevent sql injection
*/

package models

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitConnection() (*gorm.DB, error) {
	dsn := os.Getenv("POSTGRES_DATABASE")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db, nil
}
