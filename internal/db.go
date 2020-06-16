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

// DBClient database client interface
type DBClient interface {
	Close() error
	Find(out interface{}, where ...interface{}) error
	Save(value interface{}) error
}

// NewDBClient generate new DBClient implement
func NewDBClient(dsn DSN) (DBClient, error) {
	db, err := gorm.Open("mysql", string(dsn))
	if err != nil {
		return nil, fmt.Errorf("gorm.Open failed: %v", err)
	}
	err = db.DB().Ping()
	if err != nil {
		_ = db.Close()
		return nil, fmt.Errorf("sql.DB.Ping failed: %v", err)
	}
	return &dbClient{
		db: db,
	}, nil
}

type dbClient struct {
	db *gorm.DB
}

func (p *dbClient) Close() error {
	return p.db.Close()
}

func (p *dbClient) Find(out interface{}, where ...interface{}) error {
	return p.db.Find(out, where...).Error
}

func (p *dbClient) Save(value interface{}) error {
	return p.db.Save(value).Error
}
