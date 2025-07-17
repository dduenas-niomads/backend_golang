package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetConceptos(c *gin.Context) {
	conceptos := []string{"Concepto 1", "Concepto 2", "Concepto 3"}
	c.JSON(http.StatusOK, gin.H{"conceptos": conceptos})
}

func GetCiudades(c *gin.Context) {
	ciudades := []string{"Ciudad A", "Ciudad B", "Ciudad C"}
	c.JSON(http.StatusOK, gin.H{"ciudades": ciudades})
}

func GetPaises(c *gin.Context) {
	paises := []string{"País X", "País Y", "País Z"}
	c.JSON(http.StatusOK, gin.H{"paises": paises})
}

