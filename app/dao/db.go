package dao

import (
	"log"

	"github.com/ProgrammingLab/toshokan/app/conf"
	"github.com/ProgrammingLab/toshokan/app/util"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"golang.org/x/crypto/bcrypt"
)

var db *gorm.DB

func ConnectDB() error {
	const driver = "mysql"
	cfg := conf.GetConfig()
	var err error
	db, err = gorm.Open(driver, cfg.DB.DataSourceName)
	db.LogMode(cfg.Debug)
	if err != nil {
		return err
	}

	if err := migrate(); err != nil {
		return err
	}

	return insertAdminUser()
}

func utf8mb4() *gorm.DB {
	return db.Set("gorm:table_options", "CHARACTER SET utf8mb4")
}

func migrate() error {
	if err := utf8mb4().AutoMigrate(User{}).Error; err != nil {
		return err
	}

	if err := utf8mb4().AutoMigrate(Session{}).Error; err != nil {
		return err
	}
	if err := db.Model(Session{}).AddForeignKey("user_id", "users(user_id)", "CASCADE", "CASCADE").Error; err != nil {
		return err
	}

	return nil
}

func insertAdminUser() error {
	tx := db.Begin()

	cfg := conf.GetConfig()

	pwd, err := util.GenerateRandomString(20)
	if err != nil {
		return err
	}
	digest, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	admin := User{
		Name:           "admin",
		DisplayName:    "admin",
		Email:          cfg.AdminEmail,
		PasswordDigest: string(digest),
		IsAdmin:        true,
	}

	u := User{}
	res := tx.Model(u).Where("name = ?", "admin").Scan(&u)
	if !res.RecordNotFound() {
		tx.Rollback()
		return res.Error
	}

	if err := tx.Create(&admin).Error; err != nil {
		tx.Rollback()
		return err
	}

	log.Printf("admin email: %v  admin password: %v", cfg.AdminEmail, pwd)

	return tx.Commit().Error
}
