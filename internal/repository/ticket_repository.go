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
	GetTicketByContinent(continent string) ([]model.Ticket, error)
	UpdateStockCreateOrderTicket(ticketID, order int) error
	UpdateStockSuccessOrderTicket(ticketID, order int) error
	UpdateStockFailOrderTicket(ticketID, order int) error
	GetStockTicketGroupByContinent() ([]model.StockTicket, error)
	GetTicketEventByTicketID(ticketID int) (model.TicketEvent, error)
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

func (r *ticketRepository) GetTicketByContinent(continent string) ([]model.Ticket, error) {
	var tickets []model.Ticket
	query := `SELECT ticket_detail_id, type, price, continent_name, stock_ticket, stock, stock_ordered, country_name, 
		country_city, country_place, created_at, updated_at 
		FROM ticket_detail WHERE continent_name = ?`

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
	query := `UPDATE ticket_detail SET stock = stock + ?, stock_ordered = stock_ordered - ? WHERE ticket_detail_id = ?`
	_, err := r.DB.Exec(query, order, order, ticketID)
	if err != nil {
		r.logger.Error("Error when updating ticket_detail table", zap.Error(err))
		return err
	}
	return nil
}

func (r *ticketRepository) GetStockTicketGroupByContinent() ([]model.StockTicket, error) {
	var tickets []model.StockTicket
	query := `SELECT continent_name, SUM(stock) as stock FROM ticket_detail GROUP BY continent_name`

	rows, err := r.DB.Query(query)
	if err != nil {
		r.logger.Error("Error when querying ticket_detail table", zap.Error(err))
		return tickets, err
	}
	defer rows.Close()

	for rows.Next() {
		var ticket model.StockTicket
		err := rows.Scan(&ticket.Continent, &ticket.Stock)
		if err != nil {
			r.logger.Error("Error when scanning ticket_detail table", zap.Error(err))
			return tickets, err
		}
		tickets = append(tickets, ticket)
	}
	return tickets, nil
}

func (r *ticketRepository) GetTicketEventByTicketID(ticketID int) (model.TicketEvent, error) {
	var ticketEvent model.TicketEvent
	query := `SELECT td.ticket_detail_id, td.type, td.price, td.stock, td.continent_name, td.country_city, td.country_place, e.event_name, e.date, e.description
		from ticket_detail td
		left join event e on td.event_id = e.event_id
		WHERE td.ticket_detail_id = ?`

	err := r.DB.QueryRow(query, ticketID).Scan(&ticketEvent.TicketID, &ticketEvent.Type, &ticketEvent.Price, &ticketEvent.Stock,
		&ticketEvent.Continent, &ticketEvent.CountryCity, &ticketEvent.CountryPlace, &ticketEvent.EventName, &ticketEvent.Date, &ticketEvent.Description)
	if err != nil {
		r.logger.Error("Error when scanning ticket_event table", zap.Error(err))
		return ticketEvent, err
	}
	return ticketEvent, nil
}
