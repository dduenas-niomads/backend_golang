// SeedCountry.go
package main

import (
    "backend_golang/config"
    "backend_golang/models"
    
)


func SeedCountry() { // <-- ¡Así es como debe quedar!
    countries := []models.Country{
        {ID: 1, Name: "Argentina", ISOCode: "AR", Flag: "🇦🇷"},
        {ID: 2, Name: "Bolivia", ISOCode: "BO", Flag: "🇧🇴"},
        {ID: 3, Name: "Brasil", ISOCode: "BR", Flag: "🇧🇷"},
        {ID: 4, Name: "Chile", ISOCode: "CL", Flag: "🇨🇱" },
        {ID: 5, Name: "Colombia", ISOCode: "CO", Flag: "🇨🇴" },
        {ID: 6, Name: "Costa Rica", ISOCode: "CR", Flag: "🇨🇷" },
        {ID: 7, Name: "Cuba", ISOCode: "CU", Flag: "🇨🇺" },
        {ID: 8, Name: "Ecuador", ISOCode: "EC", Flag: "🇪🇨" },
        {ID: 9, Name: "El Salvador", ISOCode: "SV", Flag: "🇸🇻" },
        {ID: 10, Name: "Guatemala", ISOCode: "GT", Flag: "🇬🇹" },
        {ID: 11, Name: "Honduras", ISOCode: "HN", Flag: "🇭🇳" },
        {ID: 12, Name: "México", ISOCode: "MX", Flag: "🇲🇽" },
        {ID: 13, Name: "Nicaragua", ISOCode: "NI", Flag: "🇳🇮" },
        {ID: 14, Name: "Panamá", ISOCode: "PA", Flag: "🇵🇦" },
        {ID: 15, Name: "Paraguay", ISOCode: "PY", Flag: "🇵🇾" },
        {ID: 16, Name: "Perú", ISOCode: "PE", Flag: "🇵🇪" },
        {ID: 17, Name: "Puerto Rico", ISOCode: "PR", Flag: "🇵🇷" },
        {ID: 18, Name: "República Dominicana", ISOCode: "DO", Flag: "🇩🇴" },
        {ID: 19, Name: "Uruguay", ISOCode: "UY", Flag: "🇺🇾" },
        {ID: 20, Name: "Venezuela", ISOCode: "VE", Flag: "🇻🇪" },
        {ID: 21, Name: "España", ISOCode: "ES", Flag: "🇪🇸" },
    }

    for _, country := range countries {
        config.DB.Where(models.Country{Name: country.Name}).FirstOrCreate(&country)
    }
}