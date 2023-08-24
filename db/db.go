package db

import (
	"context"
	"fmt"
	"gorm-oracle/config"
	"gorm-oracle/model"
	"log"
	"os"
	"time"

	"github.com/dzwvip/oracle"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

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

	_logger := logger.New(
		log.New(os.Stdout, "\n[LOG]", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			Colorful:                  true,
			IgnoreRecordNotFoundError: false,
			ParameterizedQueries:      true,
			LogLevel:                  logger.Info,
		},
	)

	db, err := gorm.Open(
		oracle.Open(dsn),
		&gorm.Config{
			Logger: _logger,
			// DisableForeignKeyConstraintWhenMigrating: true,
		})
	if err != nil {
		return nil, nil, err
	}

	merr := db.AutoMigrate(&model.WkData{})
	if merr != nil {
		log.Fatal("DB Migration Error")
	}

	sqlDB, _ := db.DB()

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := sqlDB.PingContext(ctx); err != nil {
		return nil, func() { _ = sqlDB.Close() }, err
	}

	return db, func() { _ = sqlDB.Close() }, nil
}

func CloseDB(db *gorm.DB) {
	sqlDB, _ := db.DB()
	if err := sqlDB.Close(); err != nil {
		log.Fatalln(err)
	}
}
