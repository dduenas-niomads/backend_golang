package models

type User struct {
    ID       uint   `gorm:"primaryKey" json:"id"`
    Email    string `gorm:"unique" json:"email"`
    Name     string `json:"name"`
    Password string `json:"-"`
}