package client

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/koba1108/go-backend-ddd-template/internals/config"
)

func NewGorm(conf *config.DatabaseConfig) (*gorm.DB, error) {
	db, err := gorm.Open(conf.Driver, conf.DSN())
	if err != nil {
		return nil, err
	}
	return db, nil
}
