package handlers

import (
	"myzavod/database"
	"myzavod/dto"
	"myzavod/models"

	"github.com/gofiber/fiber/v2"
)

// CreateTech ...
func CreateTech(c *fiber.Ctx) error {
	createTech := new(dto.CreateTech)
	if err := c.BodyParser(createTech); err != nil {
		return c.SendStatus(fiber.StatusBadGateway)
	}

	tech := createTech.Technology()
	if err := database.DB.Create(&tech).Error; err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(tech)
}

// GetTechList ...
func GetTechList(c *fiber.Ctx) error {
	result := []*models.Technology{}
	if err := database.DB.Model(&models.Technology{}).Find(&result).Error; err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(result)
}

// EditTech ...
func EditTech(c *fiber.Ctx) error {
	tech := new(models.Technology)
	if err := c.BodyParser(tech); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	if err := database.DB.First(tech, tech.ID).Select("Name", "Cost").Updates(tech).Error; err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(tech)
}

// DeleteTech ...
func DeleteTech(c *fiber.Ctx) error {
	tech := new(models.Technology)
	if err := c.BodyParser(tech); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	if err := database.DB.Delete(tech, tech.ID).Error; err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendStatus(fiber.StatusOK)
}

// AddVariant ...
func AddVariant(c *fiber.Ctx) error {
	createVariant := new(dto.CreateVariant)
	if err := c.BodyParser(&createVariant); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	variation := createVariant.Variation()
	if err := database.DB.Create(&variation).Error; err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(variation)
}

// EditVariant ...
func EditVariant(c *fiber.Ctx) error {
	variation := new(models.Variation)
	if err := c.BodyParser(variation); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	if err := database.DB.First(variation, variation.ID).Select("Name", "Cost").Updates(variation).Error; err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(variation)
}

// DeleteVariant ...
func DeleteVariant(c *fiber.Ctx) error {
	variation := new(models.Variation)
	if err := c.BodyParser(variation); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	if err := database.DB.Delete(&models.Variation{}, variation.ID).Error; err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendStatus(fiber.StatusOK)
}
