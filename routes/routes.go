func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	countries := api.Group("/countries")

	countries.Post("/", controllers.CreateCountry)
	countries.Get("/", controllers.GetCountries)
	countries.Get("/:id", controllers.GetCountryByID)
	countries.Delete("/:id", controllers.DeleteCountry)
}
