package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/SyamSolution/ticket-management-service/config"
	"github.com/SyamSolution/ticket-management-service/config/middleware"
	"github.com/SyamSolution/ticket-management-service/internal/consumer"
	"github.com/SyamSolution/ticket-management-service/internal/handler"
	"github.com/SyamSolution/ticket-management-service/internal/repository"
	"github.com/SyamSolution/ticket-management-service/internal/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus"
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
	fiberProm := middleware.NewWithRegistry(prometheus.DefaultRegisterer, "ticket-management-service", "", "", map[string]string{})

	//=== repository lists start ===//
	ticketRepo := repository.NewTicketRepository(DB, baseDep.Logger)
	//=== repository lists end ===//

	//=== usecase lists start ===//
	ticketUsecase := usecase.NewTicketUsecase(ticketRepo, baseDep.Logger)
	//=== usecase lists end ===//

	//=== handler lists start ===//
	ticketHandler := handler.NewTicketHandler(ticketUsecase, baseDep.Logger)
	//=== handler lists end ===//

	go consumer.StartConsumer(ticketUsecase)

	app := fiber.New()

	app.Use(recover.New())
	app.Use(cors.New())
	app.Use(pprof.New())
	app.Use(logger.New(logger.Config{
		// Format: "[${time}] ${status} - ${latency} ${method} ${path}\n",
		Format:       `${time} {"router_activity" : [${status},"${latency}","${method}","${path}"], "query_param":${queryParams}, "body_param":${body}}` + "\n",
		TimeInterval: time.Millisecond,
		TimeFormat:   "02-01-2006 15:04:05",
		TimeZone:     "Indonesia/Jakarta",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	//=== metrics route
	fiberProm.RegisterAt(app, "/metrics")
	app.Use(fiberProm.Middleware)

	//=== ticket routes ===//
	app.Get("/continent/tickets/:continent", ticketHandler.GetTicketByContinent)
	app.Get("/tickets/continent-stock", ticketHandler.GetStockTicketGroupByContinent)
	app.Get("/event/ticket/:ticket_id", ticketHandler.GetTicketEventByTicketID)
	app.Group("/", middleware.Auth())
	app.Get("/tickets/continent/:continent", ticketHandler.GetAvailableTicketByContinent)
	app.Get("/tickets/type/:type", ticketHandler.GetAvailableTicketByType)

	//=== listen port ===//
	if err := app.Listen(fmt.Sprintf(":%s", os.Getenv("APP_PORT"))); err != nil {
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
