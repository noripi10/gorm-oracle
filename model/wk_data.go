package model

import "time"

type WkData struct {
	ID          uint      `json:"id"`
	DATA_1      string    `json:"data_1"`
	DATA_2      string    `json:"data_2"`
	DATA_3      string    `json:"data_3"`
	DATA_4      string    `json:"data_4"`
	DATA_5      string    `json:"data_5"`
	INSERT_DATE time.Time `json:"insert_date" gorm:"type:Date"`
	UPDATE_DATE time.Time `json:"update_date" gorm:"type:Date"`
}
