package models

import "time"

type Concept struct {
    ID          uint      `gorm:"primaryKey"`
    Name        string    `gorm:"unique"`
    Description string
    URL         string
    CreatedAt   time.Time
    UpdatedAt   time.Time
}
