package gormx

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Open opens a database connection
func Open(conf *Config, dbConfig *gorm.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Asia%sShanghai", conf.User, conf.Password, conf.Host, conf.Port, conf.Dbname, "%2F")
	db, err := gorm.Open(mysql.Open(dsn), dbConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database (%s:%v): %v", conf.Host, conf.Port, err)
	}
	return db, nil
}
