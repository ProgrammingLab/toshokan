package dao

import (
	"github.com/ProgrammingLab/toshokan/app/conf"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func ConnectDB() error {
	const driver = "mysql"
	cfg := conf.GetConfig().DB
	var err error
	db, err = gorm.Open(driver, cfg.DataSourceName)
	if err != nil {
		return err
	}
	return migrate()
}

func migrate() error {
	return nil
}
