package model

import "time"

type WkData struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	DATA_1      string    `json:"data_1" gorm:"size:100"`
	DATA_2      string    `json:"data_2" gorm:"size:100"`
	DATA_3      string    `json:"data_3" gorm:"size:100"`
	DATA_4      string    `json:"data_4" gorm:"size:100"`
	DATA_5      string    `json:"data_5" gorm:"size:100"`
	DATA_6      string    `json:"data_6" gorm:"size:100"`
	DATA_7      string    `json:"data_7" gorm:"size:100"`
	DATA_8      string    `json:"data_8" gorm:"size:100"`
	INSERT_DATE time.Time `json:"insert_date" gorm:"type:Date"`
	UPDATE_DATE time.Time `json:"update_date" gorm:"type:Date"`
}

// Tabler struct
func (wk *WkData) TableName() string {
	return "WK_DATA"
}
