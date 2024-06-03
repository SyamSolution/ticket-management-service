package handler

import (
	"github.com/SyamSolution/ticket-management-service/config"
	"github.com/SyamSolution/ticket-management-service/internal/model"
	"github.com/SyamSolution/ticket-management-service/internal/usecase"
	"github.com/SyamSolution/ticket-management-service/internal/util"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"net/url"
	"strconv"
)

type ticketHandler struct {
	ticketUsecase usecase.TicketExecutor
	logger        config.Logger
}

type TicketHandler interface {
	GetAvailableTicketByContinent(c *fiber.Ctx) error
	GetAvailableTicketByType(c *fiber.Ctx) error
	GetTicketByContinent(c *fiber.Ctx) error
	GetStockTicketGroupByContinent(c *fiber.Ctx) error
	GetTicketEventByTicketID(c *fiber.Ctx) error
}

func NewTicketHandler(ticketUsecase usecase.TicketExecutor, logger config.Logger) TicketHandler {
	return &ticketHandler{ticketUsecase: ticketUsecase, logger: logger}
}

func (handler *ticketHandler) GetAvailableTicketByContinent(c *fiber.Ctx) error {
	continent := c.Params("continent")
	decodedContinent, err := url.QueryUnescape(continent)
	if err != nil {
		handler.logger.Error("Error when unescaping continent", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(model.ResponseWithoutData{
			Meta: model.Meta{
				Code:    fiber.StatusInternalServerError,
				Message: util.ERROR_BASE_MSG,
			},
		})
	}

	tickets, err := handler.ticketUsecase.GetAvailableTicketByContinent(decodedContinent)
	if err != nil {
		handler.logger.Error("Error when getting available ticket by continent", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(model.ResponseWithoutData{
			Meta: model.Meta{
				Code:    fiber.StatusInternalServerError,
				Message: util.ERROR_BASE_MSG,
			},
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.Response{
		Data: tickets,
		Meta: model.Meta{
			Code:    fiber.StatusOK,
			Message: "Success",
		},
	})
}

func (handler *ticketHandler) GetAvailableTicketByType(c *fiber.Ctx) error {
	ticketType := c.Params("type")
	decodedTicketType, err := url.QueryUnescape(ticketType)
	if err != nil {
		handler.logger.Error("Error when unescaping ticket type", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(model.ResponseWithoutData{
			Meta: model.Meta{
				Code:    fiber.StatusInternalServerError,
				Message: util.ERROR_BASE_MSG,
			},
		})
	}

	tickets, err := handler.ticketUsecase.GetAvailableTicketByType(decodedTicketType)
	if err != nil {
		handler.logger.Error("Error when getting available ticket by type", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(model.ResponseWithoutData{
			Meta: model.Meta{
				Code:    fiber.StatusInternalServerError,
				Message: util.ERROR_BASE_MSG,
			},
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.Response{
		Data: tickets,
		Meta: model.Meta{
			Code:    fiber.StatusOK,
			Message: "Success",
		},
	})
}

func (handler *ticketHandler) GetTicketByContinent(c *fiber.Ctx) error {
	continent := c.Params("continent")
	decodedContinent, err := url.QueryUnescape(continent)
	if err != nil {
		handler.logger.Error("Error when unescaping continent", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(model.ResponseWithoutData{
			Meta: model.Meta{
				Code:    fiber.StatusInternalServerError,
				Message: util.ERROR_BASE_MSG,
			},
		})
	}

	tickets, err := handler.ticketUsecase.GetTicketByContinent(decodedContinent)
	if err != nil {
		handler.logger.Error("Error when getting ticket by continent", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(model.ResponseWithoutData{
			Meta: model.Meta{
				Code:    fiber.StatusInternalServerError,
				Message: util.ERROR_BASE_MSG,
			},
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.Response{
		Data: tickets,
		Meta: model.Meta{
			Code:    fiber.StatusOK,
			Message: "Success",
		},
	})
}

func (handler *ticketHandler) GetStockTicketGroupByContinent(c *fiber.Ctx) error {
	tickets, err := handler.ticketUsecase.GetStockTicketGroupByContinent()
	if err != nil {
		handler.logger.Error("Error when getting stock ticket group by continent", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(model.ResponseWithoutData{
			Meta: model.Meta{
				Code:    fiber.StatusInternalServerError,
				Message: util.ERROR_BASE_MSG,
			},
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.Response{
		Data: tickets,
		Meta: model.Meta{
			Code:    fiber.StatusOK,
			Message: "Success",
		},
	})
}

func (handler *ticketHandler) GetTicketEventByTicketID(c *fiber.Ctx) error {
	ticketID := c.Params("ticket_id")
	ticketIDInt, _ := strconv.Atoi(ticketID)
	ticketEvent, err := handler.ticketUsecase.GetTicketEventByTicketID(ticketIDInt)
	if err != nil {
		handler.logger.Error("Error when getting ticket event by continent", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(model.ResponseWithoutData{
			Meta: model.Meta{
				Code:    fiber.StatusInternalServerError,
				Message: util.ERROR_BASE_MSG,
			},
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.Response{
		Data: ticketEvent,
		Meta: model.Meta{
			Code:    fiber.StatusOK,
			Message: "Success",
		},
	})
}
