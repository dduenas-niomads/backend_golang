// SeedCity.go
package main

import (
    "go-gin-crud/config"
    "go-gin-crud/models"
    
)

func SeedCity() { // <-- ¡Así es como debe quedar!
    cities := []models.City{
        {Name: "Buenos Aires", CountryID: 1},
        {Name: "Córdoba", CountryID: 1},
        {Name: "Rosario", CountryID: 1},

        {Name: "La Paz", CountryID: 2},
        {Name: "Santa Cruz de la Sierra", CountryID: 2},
        {Name: "Cochabamba", CountryID: 2},

        {Name: "Brasilia", CountryID: 3},
        {Name: "São Paulo", CountryID: 3},
        {Name: "Rio de Janeiro", CountryID: 3},

        {Name: "Santiago", CountryID: 4},
        {Name: "Valparaíso", CountryID: 4},
        {Name: "Concepción", CountryID: 4},

        {Name: "Bogotá", CountryID: 5},
        {Name: "Medellín", CountryID: 5},
        {Name: "Cali", CountryID: 5},

        {Name: "San José", CountryID: 6},
        {Name: "Alajuela", CountryID: 6},
        {Name: "Cartago", CountryID: 6},

        {Name: "La Habana", CountryID: 7},
        {Name: "Santiago de Cuba", CountryID: 7},
        {Name: "Camagüey", CountryID: 7},

        {Name: "Quito", CountryID: 8},
        {Name: "Guayaquil", CountryID: 8},
        {Name: "Cuenca", CountryID: 8},

        {Name: "San Salvador", CountryID: 9},
        {Name: "Santa Ana", CountryID: 9},
        {Name: "San Miguel", CountryID: 9},

        {Name: "Ciudad de Guatemala", CountryID: 10},
        {Name: "Quetzaltenango", CountryID: 10},
        {Name: "Escuintla", CountryID: 10},

        {Name: "Tegucigalpa", CountryID: 11},
        {Name: "San Pedro Sula", CountryID: 11},
        {Name: "La Ceiba", CountryID: 11},

        {Name: "Ciudad de México", CountryID: 12},
        {Name: "Guadalajara", CountryID: 12},
        {Name: "Monterrey", CountryID: 12},

        {Name: "Managua", CountryID: 13},
        {Name: "León", CountryID: 13},
        {Name: "Granada", CountryID: 13},

        {Name: "Panamá", CountryID: 14},
        {Name: "Colón", CountryID: 14},
        {Name: "David", CountryID: 14},

        {Name: "Asunción", CountryID: 15},
        {Name: "Ciudad del Este", CountryID: 15},
        {Name: "Encarnación", CountryID: 15},

        {Name: "Lima", CountryID: 16},
        {Name: "Arequipa", CountryID: 16},
        {Name: "Trujillo", CountryID: 16},

        {Name: "San Juan", CountryID: 17},
        {Name: "Bayamón", CountryID: 17},
        {Name: "Ponce", CountryID: 17},

        {Name: "Santo Domingo", CountryID: 18},
        {Name: "Santiago de los Caballeros", CountryID: 18},
        {Name: "La Romana", CountryID: 18},

        {Name: "Montevideo", CountryID: 19},
        {Name: "Salto", CountryID: 19},
        {Name: "Paysandú", CountryID: 19},

        {Name: "Caracas", CountryID: 20},
        {Name: "Maracaibo", CountryID: 20},
        {Name: "Valencia", CountryID: 20},

        {Name: "Madrid", CountryID: 21},
        {Name: "Barcelona", CountryID: 21},
        {Name: "Valencia", CountryID: 21},
    }
    for _, city := range cities {
        config.DB.Where(models.City{Name: city.Name, CountryID: city.CountryID}).FirstOrCreate(&city)
    }
}