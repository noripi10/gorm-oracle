package db

import (
	"context"
	"fmt"
	"gorm-oracle/config"
	"log"
	"time"

	"github.com/dzwvip/oracle"
	"gorm.io/gorm"

	_ "database/sql"
)

func New(ctx context.Context) (*gorm.DB, func(), error) {
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
	if err != nil {
		return nil, nil, err
	}
	sqlDB, _ := db.DB()

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := sqlDB.PingContext(ctx); err != nil {
		return nil, func() { _ = sqlDB.Close() }, err
	}

	return db, func() { _ = sqlDB.Close() }, err
}

func CloseDB(db *gorm.DB) {
	sqlDB, _ := db.DB()
	if err := sqlDB.Close(); err != nil {
		log.Fatalln(err)
	}
}
