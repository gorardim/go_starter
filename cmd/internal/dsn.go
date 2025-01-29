package internal

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func OpenDB() (*gorm.DB, error) {
	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Asia%sShanghai", "root", "123456", "localhost", 3306, "travel", "%2F")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Asia%sShanghai", "root", "Y7yzT0wl4taQAUmQqRyR", "travel-test.c1087wljfsc8.ap-southeast-1.rds.amazonaws.com", 3306, "travel", "%2F")
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database %s: %w", dsn, err)
	}
	return db, nil
}
