package settings

import (
	"fmt"
	"log"
	"micro-ocr-service/database"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DBConn *gorm.DB
)

// connectDb
func ConnectMySqlDb() {

	dsn := os.Getenv("DB_SQL_USERNAME") + ":" + os.Getenv("DB_SQL_PASSWORD") + "@tcp(" + os.Getenv("DB_SQL_HOST") + ":" + os.Getenv("DB_SQL_PORT") + ")/" + os.Getenv("DB_SQL_DATABASE") + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("MySQL Database Connection Error $s", err)
		os.Exit(2)
	}

	fmt.Println("MySQL connection successfully!")

	// migrate database & seeder
	database.MigrateDown(db) // remove tables
	database.MigrateUp(db)   // create tables
	database.SeedAll(db)

	fmt.Println("MySQL auto migrate successfully!")

	DBConn = db

}
