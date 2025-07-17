<<<<<<< HEAD
// main.go (No necesita cambios si ya lo tienes así)
package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"backend_golang/config"
	"backend_golang/routes"
)

func main() {
	// Cargar variables de entorno
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error cargando archivo .env")
	}

	// Conectar a la base de datos
	config.ConnectDatabase()

	// Iniciar Gin
	router := gin.Default()

	// Configurar rutas
	routes.SetupRoutes(router)

	// Puerto desde .env
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Iniciar servidor
	err = router.Run(":" + port)
	if err != nil {
		log.Fatal("No se pudo iniciar el servidor")
	}
=======
package main

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Country struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var users = []User{
	{ID: 1, Name: "Alice", Email: "alice@example.com"},
	{ID: 2, Name: "Bob", Email: "bob@example.com"},
}

var countries = []Country{
	{ID: 1, Name: "España"},
	{ID: 2, Name: "Francia"},
}

func main() {
	app := fiber.New()

	storage.LoadCountriesFromFile()
	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}

	// CRUD Usuarios
	app.Post("/users", createUser)
	app.Get("/users", listUsers)
	app.Get("/users/:id", getUserByID)
	app.Put("/users/:id", updateUser)
	app.Delete("/users/:id", deleteUser)

	// CRUD Países
	app.Post("/countries", createCountry)
	app.Get("/countries", listCountries)
	app.Get("/countries/:id", getCountryByID)
	app.Put("/countries/:id", updateCountry)
	app.Delete("/countries/:id", deleteCountry)

	app.Listen(":3000")
}

// Usuarios
func createUser(c *fiber.Ctx) error {
	var newUser User
	if err := c.BodyParser(&newUser); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	newUser.ID = len(users) + 1
	users = append(users, newUser)
	return c.JSON(newUser)
}

func listUsers(c *fiber.Ctx) error {
	return c.JSON(users)
}

func getUserByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "ID inválido"})
	}
	for _, user := range users {
		if user.ID == id {
			return c.JSON(user)
		}
	}
	return c.Status(404).JSON(fiber.Map{"error": "Usuario no encontrado"})
}

func updateUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "ID inválido"})
	}
	var updatedUser User
	if err := c.BodyParser(&updatedUser); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	for i, user := range users {
		if user.ID == id {
			users[i].Name = updatedUser.Name
			users[i].Email = updatedUser.Email
			return c.JSON(users[i])
		}
	}
	return c.Status(404).JSON(fiber.Map{"error": "Usuario no encontrado"})
}

func deleteUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "ID inválido"})
	}
	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			return c.SendString(fmt.Sprintf("Usuario con ID %d eliminado", id))
		}
	}
	return c.Status(404).JSON(fiber.Map{"error": "Usuario no encontrado"})
}

// Países
func createCountry(c *fiber.Ctx) error {
	var newCountry Country
	if err := c.BodyParser(&newCountry); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	newCountry.ID = len(countries) + 1
	countries = append(countries, newCountry)
	return c.JSON(newCountry)
}

func listCountries(c *fiber.Ctx) error {
	return c.JSON(countries)
}

func getCountryByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "ID inválido"})
	}
	for _, country := range countries {
		if country.ID == id {
			return c.JSON(country)
		}
	}
	return c.Status(404).JSON(fiber.Map{"error": "País no encontrado"})
}

func updateCountry(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "ID inválido"})
	}
	var updatedCountry Country
	if err := c.BodyParser(&updatedCountry); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	for i, country := range countries {
		if country.ID == id {
			countries[i].Name = updatedCountry.Name
			return c.JSON(countries[i])
		}
	}
	return c.Status(404).JSON(fiber.Map{"error": "País no encontrado"})
}

func deleteCountry(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "ID inválido"})
	}
	for i, country := range countries {
		if country.ID == id {
			countries = append(countries[:i], countries[i+1:]...)
			return c.SendString(fmt.Sprintf("País con ID %d eliminado", id))
		}
	}
	return c.Status(404).JSON(fiber.Map{"error": "País no encontrado"})
>>>>>>> 9358fd8 (recambios)
}