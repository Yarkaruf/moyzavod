package handlers

import (
	"errors"
	"myzavod/database"
	"myzavod/dto"
	"myzavod/middlewares"
	"myzavod/models"
	"myzavod/pkg/mail"
	"myzavod/pkg/tools"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

// RegisterHandler ...
func RegisterHandler(store *session.Store) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		createUser := new(dto.CreateUser)

		if err := c.BodyParser(&createUser); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		user := createUser.User()
		if err := database.DB.Create(&user).Error; err != nil {
			log.Err(err).Msg("failed to create user")
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		// Send mail to verify
		link := tools.MakeVerifyEmailLink(user.Email)
		err := mail.SendMail(user.Email, `
		Cпасибо за регистрацию!
		Для подтверждения почты пройдите по ссылке:
		<a href="`+link+`">`+link+`</a>`)
		if err != nil {
			log.Err(err).Str("email", user.Email).Str("name", user.Name).Msg("failed to send verification email")
			// return c.SendStatus(fiber.StatusInternalServerError)
		}

		// Creating new session
		sess, err := store.Get(c)
		if err != nil {
			log.Err(err).Msg("failed to get session")
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		sess.Set(middlewares.LocalsUserID, user.ID)
		if err := sess.Save(); err != nil {
			log.Err(err).Msg("failed to save session")
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.SendStatus(fiber.StatusCreated)
	}
}

// GetUser ...
func GetUser(c *fiber.Ctx) error {
	userID := c.Locals(middlewares.LocalsUserID).(uint)

	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		log.Err(err).Msg("failed to get current user")
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.JSON(user)
}

// UpdateUser ...
func UpdateUser(c *fiber.Ctx) error {
	userID := c.Locals(middlewares.LocalsUserID).(uint)

	var update models.User
	if err := c.BodyParser(&update); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	err := database.DB.
		First(&models.User{}, userID).
		Select("name", "tel", "pay_service").
		Updates(update).
		Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.SendStatus(fiber.StatusNotFound)
		}

		log.Err(err).Msg("failed to update current user")
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendStatus(fiber.StatusOK)
}

// CreateLocation ...
func CreateLocation(c *fiber.Ctx) error {
	userID := c.Locals(middlewares.LocalsUserID).(uint)

	location := new(models.Location)
	if err := c.BodyParser(&location); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	location.UserID = userID

	if err := database.DB.Create(location).Error; err != nil {
		log.Err(err).Msg("failed to create location")
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	err := database.DB.
		First(&models.User{}, userID).
		Updates(models.User{PreferredLocation: location}).
		Error
	if err != nil {
		log.Err(err).Msg("failed to update preferred location")
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendStatus(fiber.StatusOK)
}

// UpdateLocation ...
func UpdateLocation(c *fiber.Ctx) error {
	userID := c.Locals(middlewares.LocalsUserID).(uint)

	var query struct {
		ID uint `query:"id"`
	}
	if err := c.QueryParser(&query); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	location := new(models.Location)
	if err := c.BodyParser(&location); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	location.ID = query.ID
	location.UserID = userID

	err := database.DB.Model(location).Updates(location).Error
	if err != nil {
		log.Err(err).Msg("failed to update location")
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendStatus(fiber.StatusOK)
}

// DeleteLocation ...
func DeleteLocation(c *fiber.Ctx) error {
	userID := c.Locals(middlewares.LocalsUserID).(uint)

	var query struct {
		ID uint `query:"id"`
	}
	if err := c.QueryParser(&query); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	err := database.DB.Delete(&models.Location{}, "id = ? AND user_id = ?", query.ID, userID).Error
	if err != nil {
		log.Err(err).Msg("failed to delete location")
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendStatus(fiber.StatusOK)
}
