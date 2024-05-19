package repository

import (
	"database/sql"
	"github.com/SyamSolution/ticket-management-service/config"
	"github.com/SyamSolution/ticket-management-service/internal/model"
	"go.uber.org/zap"
)

type ticketRepository struct {
	DB     *sql.DB
	logger config.Logger
}

type TicketPersister interface {
	GetAvailableTicketByContinent(continent string) ([]model.Ticket, error)
	GetAvailableTicketByType(ticketType string) ([]model.Ticket, error)
	GetTicketByID(ticketID int) (model.Ticket, error)
	UpdateStockCreateOrderTicket(ticketID, order int) error
	UpdateStockSuccessOrderTicket(ticketID, order int) error
	UpdateStockFailOrderTicket(ticketID, order int) error
}

func NewTicketRepository(DB *sql.DB, logger config.Logger) TicketPersister {
	return &ticketRepository{DB: DB, logger: logger}
}

func (r *ticketRepository) GetAvailableTicketByContinent(continent string) ([]model.Ticket, error) {
	var tickets []model.Ticket
	query := `SELECT ticket_detail_id, type, price, continent_name, stock_ticket, stock, stock_ordered, country_name, 
       	country_city, country_place, created_at, updated_at 
		FROM ticket_detail WHERE continent_name = ? AND stock > 0`

	rows, err := r.DB.Query(query, continent)
	if err != nil {
		r.logger.Error("Error when querying ticket_detail table", zap.Error(err))
		return tickets, err
	}
	defer rows.Close()

	for rows.Next() {
		var ticket model.Ticket
		err := rows.Scan(&ticket.TicketID, &ticket.Type, &ticket.Price, &ticket.ContinentName, &ticket.StockTicket, &ticket.Stock,
			&ticket.StockOrdered, &ticket.CountryName, &ticket.CountryCity, &ticket.CountryPlace, &ticket.CreatedAt, &ticket.UpdatedAt)
		if err != nil {
			r.logger.Error("Error when scanning ticket_detail table", zap.Error(err))
			return tickets, err
		}
		tickets = append(tickets, ticket)
	}
	return tickets, nil
}

func (r *ticketRepository) GetAvailableTicketByType(ticketType string) ([]model.Ticket, error) {
	var tickets []model.Ticket
	query := `SELECT ticket_detail_id, type, price, continent_name, stock_ticket, stock, stock_ordered, country_name, 
	   	country_city, country_place, created_at, updated_at 
		FROM ticket_detail WHERE type = ? AND stock > 0`

	rows, err := r.DB.Query(query, ticketType)
	if err != nil {
		r.logger.Error("Error when querying ticket_detail table", zap.Error(err))
		return tickets, err
	}
	defer rows.Close()

	for rows.Next() {
		var ticket model.Ticket
		err := rows.Scan(&ticket.TicketID, &ticket.Type, &ticket.Price, &ticket.ContinentName, &ticket.StockTicket, &ticket.Stock,
			&ticket.StockOrdered, &ticket.CountryName, &ticket.CountryCity, &ticket.CountryPlace, &ticket.CreatedAt, &ticket.UpdatedAt)
		if err != nil {
			r.logger.Error("Error when scanning ticket_detail table", zap.Error(err))
			return tickets, err
		}
		tickets = append(tickets, ticket)
	}
	return tickets, nil
}

func (r *ticketRepository) GetTicketByID(ticketID int) (model.Ticket, error) {
	var ticket model.Ticket
	query := `SELECT ticket_detail_id, type, price, continent_name, stock_ticket, stock, stock_ordered, country_name, 
		country_city, country_place, created_at, updated_at 
		FROM ticket_detail WHERE ticket_detail_id = ?`

	err := r.DB.QueryRow(query, ticketID).Scan(&ticket.TicketID, &ticket.Type, &ticket.Price, &ticket.ContinentName, &ticket.StockTicket, &ticket.Stock,
		&ticket.StockOrdered, &ticket.CountryName, &ticket.CountryCity, &ticket.CountryPlace, &ticket.CreatedAt, &ticket.UpdatedAt)
	if err != nil {
		r.logger.Error("Error when scanning ticket_detail table", zap.Error(err))
		return ticket, err
	}
	return ticket, nil
}

func (r *ticketRepository) UpdateStockCreateOrderTicket(ticketID, order int) error {
	query := `UPDATE ticket_detail SET stock = stock - ?, stock_ordered = stock_ordered + ? WHERE ticket_detail_id = ?`
	_, err := r.DB.Exec(query, order, order, ticketID)
	if err != nil {
		r.logger.Error("Error when updating ticket_detail table", zap.Error(err))
		return err
	}
	return nil
}

func (r *ticketRepository) UpdateStockSuccessOrderTicket(ticketID, order int) error {
	query := `UPDATE ticket_detail SET stock_ticket = stock_ticket - ? WHERE ticket_detail_id = ?`
	_, err := r.DB.Exec(query, order, ticketID)
	if err != nil {
		r.logger.Error("Error when updating ticket_detail table", zap.Error(err))
		return err
	}
	return nil
}

func (r *ticketRepository) UpdateStockFailOrderTicket(ticketID, order int) error {
	query := `UPDATE ticket_detail SET stock = stock + ?, stock_order = stock_order - ? WHERE ticket_detail_id = ?`
	_, err := r.DB.Exec(query, order, order, ticketID)
	if err != nil {
		r.logger.Error("Error when updating ticket_detail table", zap.Error(err))
		return err
	}
	return nil
}
