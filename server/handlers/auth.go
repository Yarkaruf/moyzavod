package handlers

import (
	"errors"
	"myzavod/database"
	"myzavod/middlewares"
	"myzavod/models"
	"myzavod/pkg/mail"
	"myzavod/pkg/tools"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

// LoginHandler ...
func LoginHandler(store *session.Store) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var body struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		if err := c.BodyParser(&body); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		hash := tools.PasswordHash(body.Password)
		user := new(models.User)
		if err := database.DB.First(&user, "email = ? and (password_hash = ? or alter_password_hash = ?)", body.Email, hash, hash).Error; err != nil {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		// Creating session
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

		return c.SendStatus(fiber.StatusOK)
	}
}

// LogoutHandler ...
func LogoutHandler(store *session.Store) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		sess, err := store.Get(c)
		if err != nil {
			log.Err(err).Msg("failed to get session")
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		if err := sess.Destroy(); err != nil {
			log.Err(err).Msg("failed to delete session")
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.SendStatus(fiber.StatusOK)
	}
}

// Recovery ...
func Recovery(c *fiber.Ctx) error {
	var body struct {
		Email string `json:"email"`
	}
	if err := c.BodyParser(&body); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	user := new(models.User)
	if err := database.DB.First(&user, "email = ?", body.Email).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.SendStatus(fiber.StatusNotFound)
		}

		return c.SendStatus(fiber.StatusInternalServerError)
	}

	newPassword := tools.RandomPassword()
	user.AlterPasswordHash = tools.PasswordHash(newPassword)
	database.DB.Save(user)

	mail.SendMail(user.Email, `
		Для входа используйте резервный пароль:
		`+newPassword)

	return c.SendStatus(fiber.StatusCreated)
}

// VerifyEmail ...
func VerifyEmail(c *fiber.Ctx) error {
	var query struct {
		Email string `query:"email"`
		Code  string `query:"code"`
	}
	if err := c.QueryParser(&query); err != nil {
		return err
	}

	if query.Code != tools.PasswordHash(query.Email) {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	err := database.DB.Model(&models.User{}).Where("email = ?", query.Email).Update("email_verified = ?", true).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.SendStatus(fiber.StatusNotFound)
		}

		log.Err(err).Msg("failed to verify email")
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendStatus(fiber.StatusOK)
}

// SendVerificationEmail ...
func SendVerificationEmail(c *fiber.Ctx) error {
	userID := c.Locals(middlewares.LocalsUserID).(uint)

	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.SendStatus(fiber.StatusNotFound)
		}

		log.Err(err).Msg("failed to get user")
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	link := tools.MakeVerifyEmailLink(user.Email)
	err := mail.SendMail(user.Email, `
		Для верификации почты пройдите по ссылке:
		<a href="`+link+`">`+link+`</a>`)
	if err != nil {
		log.Err(err).Msg("failed to send verification email")
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendStatus(fiber.StatusCreated)
}
