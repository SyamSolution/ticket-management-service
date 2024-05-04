package handler

import (
	"github.com/SyamSolution/ticket-management-service/config"
	"github.com/SyamSolution/ticket-management-service/internal/usecase"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"net/url"
)

type ticketHandler struct {
	ticketUsecase usecase.TicketExecutor
	logger        config.Logger
}

type TicketHandler interface {
	GetAvailableTicketByContinent(c *fiber.Ctx) error
	GetAvailableTicketByType(c *fiber.Ctx) error
}

func NewTicketHandler(ticketUsecase usecase.TicketExecutor, logger config.Logger) TicketHandler {
	return &ticketHandler{ticketUsecase: ticketUsecase, logger: logger}
}

func (handler *ticketHandler) GetAvailableTicketByContinent(c *fiber.Ctx) error {
	continent := c.Params("continent")
	decodedContinent, err := url.QueryUnescape(continent)
	if err != nil {
		handler.logger.Error("Error when unescaping continent", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	tickets, err := handler.ticketUsecase.GetAvailableTicketByContinent(decodedContinent)
	if err != nil {
		handler.logger.Error("Error when getting available ticket by continent", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success",
		"data":    tickets,
	})
}

func (handler *ticketHandler) GetAvailableTicketByType(c *fiber.Ctx) error {
	ticketType := c.Params("type")
	decodedTicketType, err := url.QueryUnescape(ticketType)
	if err != nil {
		handler.logger.Error("Error when unescaping ticket type", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	tickets, err := handler.ticketUsecase.GetAvailableTicketByType(decodedTicketType)
	if err != nil {
		handler.logger.Error("Error when getting available ticket by type", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success",
		"data":    tickets,
	})
}