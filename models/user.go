package models

type User struct {
    ID        uint   `json:"id" gorm:"primaryKey"`
    Name      string `json:"name"`
    LastName  string `json:"last_name"`
    Email     string `json:"email"`
    Password  string `json:"password"`
}