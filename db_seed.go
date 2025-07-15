
package main

import (
    "gorm.io/gorm" 
)


func SeedAllData(db *gorm.DB) { 
  
    SeedCountry() 
    SeedCity()    
}