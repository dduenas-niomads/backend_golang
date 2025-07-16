func CreateCountry(c *fiber.Ctx) error {
	country := new(models.Country)
	if err := c.BodyParser(country); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}
	if country.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Country name is required",
		})
	}

	// Agrega ID autoincremental
	country.ID = len(storage.Countries) + 1
	storage.Countries = append(storage.Countries, *country)
	storage.SaveCountriesToFile()

	return c.Status(fiber.StatusCreated).JSON(country)
}

