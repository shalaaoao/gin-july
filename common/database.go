package common

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"net/url"
)

var DB map[string]*gorm.DB

func InitDB() map[string]*gorm.DB {
	DB = map[string]*gorm.DB{}

	databaseList := []string{"july", "kf"}
	for _, databaseName := range databaseList {
		DB[databaseName] = baseInitDb(databaseName)
	}

	return DB
}

func baseInitDb(databaseName string) *gorm.DB {
	driverName := viper.GetString("mysql." + databaseName + ".driverName")
	host := viper.GetString("mysql." + databaseName + ".host")
	port := viper.GetString("mysql." + databaseName + ".port")
	database := viper.GetString("mysql." + databaseName + ".database")
	username := viper.GetString("mysql." + databaseName + ".username")
	password := viper.GetString("mysql." + databaseName + ".password")
	charset := viper.GetString("mysql." + databaseName + ".charset")
	loc := viper.GetString("mysql." + databaseName + ".loc")

	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=%s",
		username,
		password,
		host,
		port,
		database,
		charset,
		url.QueryEscape(loc))

	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("fail to connect database, err: " + err.Error())
	}
	//db.AutoMigrate(&model.User{})

	return db
}

func GetDB(databaseName string) *gorm.DB {
	return DB[databaseName]
}
