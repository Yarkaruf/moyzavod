package main

import (
	"crypto/tls"

	"myzavod/database"
	"myzavod/handlers"
	"myzavod/middlewares"
	"myzavod/models"
	"net"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/badger"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/acme/autocert"
)

func main() {
	if err := database.Init(); err != nil {
		log.Err(err).Msg("failed to init database")
		panic(err)
	}

	store := session.New(session.Config{
		Storage: badger.New(badger.Config{
			Database: "./sessions.badger",
		}),
	})

	app := fiber.New()

	app.Use(logger.New())

	api := app.Group("/api")

	// Authentication
	api.Post("/register", handlers.RegisterHandler(store))
	api.Post("/login", handlers.LoginHandler(store))
	api.Get("/logout", handlers.LogoutHandler(store))
	api.Post("/recovery", handlers.Recovery)
	api.Post("/verify", handlers.SendVerificationEmail)
	api.Get("/verify", handlers.VerifyEmail)

	// Users
	users := api.Group("/user", middlewares.RetrieveUser(store))
	users.Get("/", handlers.GetUser)
	users.Post("/", handlers.UpdateUser)
	users.Post("/location", handlers.CreateLocation)
	users.Patch("/location/:id", handlers.UpdateLocation)
	users.Delete("/location/:id", handlers.DeleteLocation)

	// Orders
	orders := api.Group("/orders", middlewares.RetrieveUser(store))
	orders.Post("/", handlers.CreateOrder)
	orders.Post("/calculate", handlers.Calculate)
	orders.Get("/", handlers.GetOrders)
	orders.Delete("/:id", handlers.DeleteOrder)

	// Tech info
	api.Post("/tech", handlers.CreateTech)
	api.Get("/tech", handlers.GetTechList)
	api.Patch("/tech", handlers.EditTech)
	api.Delete("/tech", handlers.DeleteTech)

	// Plastic variants
	api.Post("/variant", handlers.AddVariant)
	api.Patch("/variant", handlers.EditVariant)
	api.Delete("/variant", handlers.DeleteVariant)

	// DEPRECATED: Old and bad params
	api.Get("/params", func(c *fiber.Ctx) error {
		p := new(models.Params)
		database.DB.First(p)
		return c.JSON(p)
	})
	api.Post("/params", func(c *fiber.Ctx) error {
		p := new(models.Params)
		if err := c.BodyParser(p); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		p.ID = 1
		database.DB.Create(p)
		database.DB.Save(p)
		return c.SendStatus(fiber.StatusOK)
	})

	api.Post("/calculate-model", handlers.CalculateWithoutCost)

	// API route not found
	api.Get("/*", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusNotFound)
	})

	// Static like user models
	app.Static("/file/*", "./files_storage")
	// Dashboard frontend hosted from here
	app.Static("/", "./public")
	app.Get("/*", func(ctx *fiber.Ctx) error {
		return ctx.SendFile("./public/index.html")
	})

	app.Listen(":80")
	// app.Listener(TLSListener(":443"))
}

// TLSListener returns listener which is created with autocert
func TLSListener(addr string) net.Listener {
	// Let’s Encrypt has rate limits: https://letsencrypt.org/docs/rate-limits/
	// It's recommended to use it's staging environment to test the code:
	// https://letsencrypt.org/docs/staging-environment/

	// Certificate manager
	m := &autocert.Manager{
		Prompt: autocert.AcceptTOS,
		// Replace with your domain
		HostPolicy: autocert.HostWhitelist("example.com"),
		// Folder to store the certificates
		Cache: autocert.DirCache("./certs"),
	}

	// TLS Config
	tlsCert := &tls.Config{
		// Get Certificate from Let's Encrypt
		GetCertificate: m.GetCertificate,
		// By default NextProtos contains the "h2"
		// This has to be removed since Fasthttp does not support HTTP/2
		// Or it will cause a flood of PRI method logs
		// http://webconcepts.info/concepts/http-method/PRI
		NextProtos: []string{
			"http/1.1", "acme-tls/1",
		},
	}

	ln, err := net.Listen("tcp", addr)
	if err != nil {
		log.Err(err).Msg("failed to listen")
	}

	return tls.NewListener(ln, tlsCert)
}
