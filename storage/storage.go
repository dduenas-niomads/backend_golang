package storage

import (
	"encoding/json"
	"os"

	"backend_golang/models"
)

var Countries []models.Country // ✅ Exportada si la primera letra es MAYÚSCULA

func LoadCountriesFromFile() {
	data, err := os.ReadFile("data/countries.json")
	if err == nil {
		json.Unmarshal(data, &Countries)
	}
}

func SaveCountriesToFile() {
	data, _ := json.MarshalIndent(Countries, "", "  ")
	_ = os.WriteFile("data/countries.json", data, 0644)
}
