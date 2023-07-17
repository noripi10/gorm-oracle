package db

import (
	"fmt"
	"gorm-oracle/config"
	"log"

	"github.com/dzwvip/oracle"
	"gorm.io/gorm"

	_ "database/sql"
)

func NewDB() (*gorm.DB, error) {
	// user := os.Getenv("ORACLE_DB_USER")
	// password := os.Getenv("ORACLE_DB_PASSWORD")
	// address := os.Getenv("ORACLE_DB_ADDRESS")
	// port := os.Getenv("ORACLE_DB_PORT")
	// service := os.Getenv("ORACLE_DB_SERVICE")

	conf, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	dsn := fmt.Sprintf("%s/%s@%s:%v/%s",
		conf.DBUser,
		conf.DBPassword,
		conf.DBAddress,
		conf.DBPort,
		conf.DBService,
	)
	db, err := gorm.Open(oracle.Open(dsn), &gorm.Config{})

	return db, err
}

func CloseDB(db *gorm.DB) {
	sqlDB, _ := db.DB()
	if err := sqlDB.Close(); err != nil {
		log.Fatalln(err)
	}
}
