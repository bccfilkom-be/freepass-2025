package mariadb

import (
	"freepass-bcc/pkg/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDatabase() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(config.LoadDataSourceName()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		return nil, err
	}

	return db, nil
}
