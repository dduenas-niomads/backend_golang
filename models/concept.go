package models

import "gorm.io/gorm"

type Concept struct {
	gorm.Model
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Url         string `json:"url" binding:"required,url"`
}
