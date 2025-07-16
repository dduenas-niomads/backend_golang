package models

type Country struct {
    ID      uint   `json:"id" gorm:"primaryKey"`
    Name    string `json:"name"`
    ISOCode string `json:"iso_code"`
    Flag    string `json:"flag"`
}