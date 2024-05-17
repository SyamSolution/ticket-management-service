package main

import (
	"fmt"
	"github.com/SyamSolution/ticket-management-service/config"
	"github.com/SyamSolution/ticket-management-service/config/middleware"
	middleware2 "github.com/SyamSolution/ticket-management-service/config/middleware"
	"github.com/SyamSolution/ticket-management-service/internal/handler"
	"github.com/SyamSolution/ticket-management-service/internal/repository"
	"github.com/SyamSolution/ticket-management-service/internal/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus"
	"log"
	"os"
)

func main() {
	baseDep := config.NewBaseDep()
	loadEnv(baseDep.Logger)
	DB, err := config.NewDbPool(baseDep.Logger)
	if err != nil {
		os.Exit(1)
	}

	dbCollector := middleware.NewStatsCollector("assesment", DB)
	prometheus.MustRegister(dbCollector)
	fiberProm := middleware.NewWithRegistry(prometheus.DefaultRegisterer, "smilley", "", "", map[string]string{})

	//=== repository lists start ===//
	ticketRepo := repository.NewTicketRepository(DB, baseDep.Logger)
	//=== repository lists end ===//

	//=== usecase lists start ===//
	ticketUsecase := usecase.NewTicketUsecase(ticketRepo, baseDep.Logger)
	//=== usecase lists end ===//

	//=== handler lists start ===//
	ticketHandler := handler.NewTicketHandler(ticketUsecase, baseDep.Logger)
	//=== handler lists end ===//

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	//=== metrics route
	fiberProm.RegisterAt(app, "/metrics")
	app.Use(fiberProm.Middleware)

	//=== ticket routes ===//
	app.Group("/", middleware2.Auth())
	app.Get("/tickets/continent/:continent", ticketHandler.GetAvailableTicketByContinent)
	app.Get("/tickets/type/:type", ticketHandler.GetAvailableTicketByType)

	//=== listen port ===//
	if err := app.Listen(fmt.Sprintf(":%s", "3001")); err != nil {
		log.Fatal(err)
	}
}

func loadEnv(logger config.Logger) {
	_, err := os.Stat(".env")
	if err == nil {
		err = godotenv.Load()
		if err != nil {
			logger.Error("no .env files provided")
		}
	}
}
