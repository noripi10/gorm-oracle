package main

import (
	"encoding/json"
	"fmt"
	"gorm-oracle/db"
	"gorm-oracle/mail"
	"gorm-oracle/model"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("goram-oralce")
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	db, err := db.NewDB()
	if err != nil {
		log.Fatal(err)
	}

	// CREATE
	// newWkData := model.WkData{ID: 999, DATA_1: "hoge", DATA_2: "fuga", DATA_3: "hogehoge", DATA_4: "fugafuga", DATA_5: "123"}
	// result := db.Create(&newWkData)
	// fmt.Println("result", result)

	// READ
	var wkDatas []model.WkData
	db.Where("ID >= ? AND ID <= ?", 10, 11).Order("id desc").Find(&wkDatas)
	fmt.Println(wkDatas)
	for index, wkData := range wkDatas {
		fmt.Println(index, wkData.ID)
	}

	json, err := json.Marshal(wkDatas)
	if err != nil {
		log.Fatal(err)
	}

	// UPDATE
	// result := db.Model(model.WkData{}).Where("id >= ? and id <= ?", 10, 11).Updates(model.WkData{DATA_5: "hoge"})
	// fmt.Println("result", result)

	// DELETE
	// db.Exec("delete wk_data where id = 1")

	message := "WK_DATA: " + "\r\n" + string(json)
	mail.SendMail(message)
}