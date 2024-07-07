package main

import (
	"context"
	"fmt"
	"gorm-oracle/db"
	"gorm-oracle/model"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("goram-oralce")
	// .env
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	run(context.Background())
}

func run(ctx context.Context) {
	db, cleanup, err := db.New(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer cleanup()
	// READ
	var wkDatas []model.WkData
	db.Where("ID >= ? AND ID <= ?", 1, 2).Order("id desc").Find(&wkDatas)
	fmt.Println(wkDatas)
	for index, wkData := range wkDatas {
		fmt.Println(index, wkData.ID)
	}

	// CREATE
	// newWkData := model.WkData{ID: 999, DATA_1: "hoge", DATA_2: "fuga", DATA_3: "hogehoge", DATA_4: "fugafuga", DATA_5: "123"}
	// result := db.Create(&newWkData)
	// fmt.Println("result", result)

	// UPDATE
	// result := db.Model(&model.WkData{}).Where("id = ?", 999).Updates(model.WkData{DATA_1: strconv.Itoa(rand.Intn(100))})
	// fmt.Println("result", result)

	// DELETE
	// db.Exec("delete wk_data where id = 1")

	// mail.SendMail(wkDatas)
}
