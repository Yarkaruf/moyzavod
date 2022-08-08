package middlewares

import (
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/rs/zerolog/log"
)

// LocalsUserID ...
const LocalsUserID = "userid"

// RetrieveUser ...
func RetrieveUser(store *session.Store) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		sess, err := store.Get(c)
		if err != nil {
			log.Info().Msg("no session")
			return err
		}

		if sess.Fresh() {
			log.Info().Msg("Fresh session")
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		log.Info().Msg(`"` + strings.Join(sess.Keys(), `","`) + `"`)

		// Getting user id
		userID, ok := sess.Get(LocalsUserID).(uint)
		log.Info().Msg(strconv.Itoa(int(userID)))
		if !ok {
			log.Info().Msg("no userid")
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		// Setting locals
		c.Locals(LocalsUserID, userID)
		return c.Next()
	}
}
