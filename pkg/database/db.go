package database

import (
	"github.com/go-practice/config"
	"github.com/go-practice/logger"
	mysql "gorm.io/driver/mysql"
	gorm "gorm.io/gorm"
)

var Db *gorm.DB

func InitDb() *gorm.DB {
	Db = connectDB()
	return Db
}

func connectDB() *gorm.DB {
	db_username := config.Config.GetString("db.username")
	db_password := config.Config.GetString("db.assword")
	db_host := config.Config.GetString("db.host")
	db_port := config.Config.GetString("db.port")
	db_name := config.Config.GetString("db.name")

	dsn := db_username + ":" + db_password + "@tcp" + "(" + db_host + ":" + db_port + ")/" + db_name + "?" + "parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		logger.Fatal("Error while initializing the db")
		return nil
	}

	dbInstace, err := db.DB()

	if err != nil {
		logger.Fatal("Error while initializing the db")
		return nil
	}

	dbInstace.SetConnMaxIdleTime(200) // set idle connections to 20
	dbInstace.SetMaxIdleConns(100)
	dbInstace.SetMaxOpenConns(200)

	return db
}
