package provider

import (
	"fmt"

	"app/pkg/gormx"
	"app/services/internal/config"

	"gorm.io/gorm"
)

func NewDb(conf *config.Config) *gorm.DB {
	dbConf := conf.Database["travel"]

	db, err := gormx.Open(dbConf, &gorm.Config{
		Logger: gormx.NewLogger(),
	})
	if err != nil {
		panic(fmt.Sprintf("failed to connect database: %v", err))
	}
	sqlDb, err := db.DB()
	if err != nil {
		panic(fmt.Errorf("db open error: %w", err))
	}
	sqlDb.SetMaxIdleConns(5)
	sqlDb.SetMaxOpenConns(10)

	if err != nil {
		panic(fmt.Sprintf("failed to connect to database (%s:%v): %v", dbConf.Host, dbConf.Port, err))
	}
	return db
}
