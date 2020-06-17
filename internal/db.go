package internal

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// DSN Data source name
type DSN string

// GetDSN get DSN from environment variable
func GetDSN() DSN {
	return DSN(os.Getenv("DSN"))
}

// NewDBClient generate new DBClient implement
func NewDBClient(dsn DSN) (*gorm.DB, error) {
	db, err := gorm.Open("mysql", string(dsn))
	if err != nil {
		return nil, fmt.Errorf("gorm.Open failed: %v", err)
	}
	err = db.DB().Ping()
	if err != nil {
		_ = db.Close()
		return nil, fmt.Errorf("sql.DB.Ping failed: %v", err)
	}
	return db, nil
}
