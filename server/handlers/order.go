package handlers

import (
	"encoding/json"
	"myzavod/database"
	"myzavod/dto"
	"myzavod/middlewares"
	"myzavod/models"
	"myzavod/pkg/calculator"
	"myzavod/pkg/tools"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/hschendel/stl"
	"github.com/rs/zerolog/log"
)

// CalculateWithoutCost ...
func CalculateWithoutCost(c *fiber.Ctx) error {
	// Getting model
	model, err := c.FormFile("model")
	if err != nil {
		log.Err(err).Msg("failed to read form file")
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	file, err := model.Open()
	if err != nil {
		log.Err(err).Msg("failed to open model")
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	defer file.Close()

	solid, err := stl.ReadAll(file)
	if err != nil {
		log.Err(err).Msg("failed to read stl model")
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	volume := calculator.CalculateVolume(solid)
	area := calculator.CalculateSurface(solid)
	l := solid.Measure().Len

	return c.JSON(fiber.Map{
		"volume": volume,
		"area":   area,
		"bounds": fiber.Map{
			"x": l[0],
			"y": l[1],
			"z": l[2],
		},
	})
}

// Calculate ...
func Calculate(c *fiber.Ctx) error {
	createOrder := new(dto.CreateOrder)
	json.Unmarshal([]byte(c.FormValue("order")), &createOrder)

	// Getting model
	model, err := c.FormFile("model")
	if err != nil {
		log.Err(err).Msg("failed to read model")
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	file, err := model.Open()
	if err != nil {
		log.Err(err).Msg("failed to open model")
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	defer file.Close()

	var serviceInfo models.Service
	if err := database.DB.First(&serviceInfo).Error; err != nil {
		log.Err(err).Msg("service info does not exists, replace with relevant info in admin panel")
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	var technology models.Technology
	if err := database.DB.First(&technology, createOrder.Technology).Error; err != nil {
		log.Err(err).Msg("service info does not exists, replace with relevant info in admin panel")
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	var platic models.Variation
	if err := database.DB.First(&platic, createOrder.Plastic).Error; err != nil {
		log.Err(err).Msg("service info does not exists, replace with relevant info in admin panel")
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// Report for model
	report, err := calculator.Calculate(calculator.Details{
		TechnologyFilling: createOrder.Quality,
		Quality:           createOrder.Quality,
		PlasticCost:       platic.Cost,
		TechnologyCost:    technology.Cost,
		ServiceCost:       serviceInfo.Cost,
	}, file)
	if err != nil {
		log.Err(err).Msg("failed calculate volume for model")
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(report)
}

// CreateOrder ...
func CreateOrder(c *fiber.Ctx) error {
	userID := c.Locals(middlewares.LocalsUserID).(uint)

	order := new(models.Order)
	json.Unmarshal([]byte(c.FormValue("order")), &order)
	order.UserID = userID
	if err := database.DB.Create(&order).Error; err != nil {
		log.Err(err).Msg("failed to create order")
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// Saving model
	model, err := c.FormFile("model")
	if err != nil {
		log.Err(err).Msg("failed to read model")
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	file, err := model.Open()
	if err != nil {
		log.Err(err).Msg("failed to open model")
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	defer file.Close()

	var serviceInfo models.Service
	if err := database.DB.First(&serviceInfo).Error; err != nil {
		log.Err(err).Msg("service info does not exists, replace with relevant info in admin panel")
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	var technology models.Technology
	if err := database.DB.First(&technology, order.Technology).Error; err != nil {
		log.Err(err).Msg("service info does not exists, replace with relevant info in admin panel")
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	var platic models.Variation
	if err := database.DB.First(&platic, order.Plastic).Error; err != nil {
		log.Err(err).Msg("service info does not exists, replace with relevant info in admin panel")
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// Report for model
	report, err := calculator.Calculate(calculator.Details{
		TechnologyFilling: order.Quality,
		Quality:           order.Quality,
		PlasticCost:       platic.Cost,
		TechnologyCost:    technology.Cost,
		ServiceCost:       serviceInfo.Cost,
	}, file)
	if err != nil {
		log.Err(err).Msg("failed calculate volume for model")
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	order.Cost = uint(report.Cost)

	order.FileName = model.Filename
	// Creates unique filename that depends on orderID
	name, err := tools.HashID(int64(order.ID))
	if err != nil {
		log.Err(err).Msg("failed to encode hashid")
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	path := "./files_storage/" + name + filepath.Ext(model.Filename)
	err = c.SaveFile(model, path)
	if err != nil {
		log.Err(err).Msg("failed to save model")
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// Saving preview
	preview, err := c.FormFile("preview")
	if err != nil {
		log.Err(err).Msg("failed to read preview")
		return err
	}

	path = "./files_storage/" + name + "_preview" + filepath.Ext(preview.Filename)
	err = c.SaveFile(preview, path)
	if err != nil {
		log.Err(err).Msg("failed to save model")
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendStatus(fiber.StatusCreated)
}

// GetOrders ...
func GetOrders(c *fiber.Ctx) error {
	userID := c.Locals(middlewares.LocalsUserID).(uint)

	result := []*models.Order{}
	database.DB.Where("user_id = ?", userID).Order("created_at desc").Find(&result)
	return c.JSON(result)
}

// DeleteOrder ...
func DeleteOrder(c *fiber.Ctx) error {
	userID := c.Locals(middlewares.LocalsUserID).(uint)

	id := c.Params("id")
	if id == "" {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	if err := database.DB.Delete(&models.Order{}, "user_id = ? and id = ?", userID, id).Error; err != nil {
		log.Err(err).Msg("failed to delete order")
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendStatus(fiber.StatusOK)
}
