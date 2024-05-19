package usecase

import (
	"github.com/SyamSolution/ticket-management-service/config"
	"github.com/SyamSolution/ticket-management-service/internal/model"
	"github.com/SyamSolution/ticket-management-service/internal/repository"
	"go.uber.org/zap"
)

type ticketUsecase struct {
	ticketRepo repository.TicketPersister
	logger     config.Logger
}

type TicketExecutor interface {
	GetAvailableTicketByContinent(continent string) ([]model.TicketResponse, error)
	GetAvailableTicketByType(ticketType string) ([]model.TicketResponse, error)
	UpdateStockTicket(ticketID, order int, typeStock string) error
}

func NewTicketUsecase(ticketRepo repository.TicketPersister, logger config.Logger) TicketExecutor {
	return &ticketUsecase{ticketRepo: ticketRepo, logger: logger}
}

func (uc *ticketUsecase) GetAvailableTicketByContinent(continent string) ([]model.TicketResponse, error) {
	var TicketsResponse []model.TicketResponse

	tickets, err := uc.ticketRepo.GetAvailableTicketByContinent(continent)
	if err != nil {
		uc.logger.Error("Error when getting available ticket by continent", zap.Error(err))
		return TicketsResponse, err
	}

	for _, ticket := range tickets {
		TicketsResponse = append(TicketsResponse, model.TicketResponse{
			TicketID:      ticket.TicketID,
			Type:          ticket.Type,
			Price:         ticket.Price,
			ContinentName: ticket.ContinentName,
			Stock:         ticket.Stock,
			CountryName:   ticket.CountryName,
			CountryCity:   ticket.CountryCity,
			CountryPlace:  ticket.CountryPlace,
		})
	}
	return TicketsResponse, nil
}

func (uc *ticketUsecase) GetAvailableTicketByType(ticketType string) ([]model.TicketResponse, error) {
	var TicketsResponse []model.TicketResponse

	tickets, err := uc.ticketRepo.GetAvailableTicketByType(ticketType)
	if err != nil {
		uc.logger.Error("Error when getting available ticket by type", zap.Error(err))
		return TicketsResponse, err
	}

	for _, ticket := range tickets {
		TicketsResponse = append(TicketsResponse, model.TicketResponse{
			TicketID:      ticket.TicketID,
			Type:          ticket.Type,
			Price:         ticket.Price,
			ContinentName: ticket.ContinentName,
			Stock:         ticket.Stock,
			CountryName:   ticket.CountryName,
			CountryCity:   ticket.CountryCity,
			CountryPlace:  ticket.CountryPlace,
		})
	}
	return TicketsResponse, nil
}

func (uc *ticketUsecase) UpdateStockTicket(ticketID, order int, typeStock string) error {
	switch typeStock {
	case "create":
		if err := uc.ticketRepo.UpdateStockCreateOrderTicket(ticketID, order); err != nil {
			uc.logger.Error("Error when updating stock ticket", zap.Error(err))
			return err
		}
		return nil
	case "success":
		if err := uc.ticketRepo.UpdateStockSuccessOrderTicket(ticketID, order); err != nil {
			uc.logger.Error("Error when updating stock ticket", zap.Error(err))
			return err
		}
		return nil
	case "failed":
		if err := uc.ticketRepo.UpdateStockFailOrderTicket(ticketID, order); err != nil {
			uc.logger.Error("Error when updating stock ticket", zap.Error(err))
			return err
		}
		return nil
	}

	return nil
}
