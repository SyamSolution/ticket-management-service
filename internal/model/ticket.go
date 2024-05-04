package model

import "time"

type Ticket struct {
	TicketID      int       `json:"ticket_id"`
	Type          string    `json:"type"`
	Price         int       `json:"price"`
	ContinentName string    `json:"continent_name"`
	StockTicket   int       `json:"stock_ticket"`
	Stock         int       `json:"stock"`
	StockOrdered  int       `json:"stock_ordered"`
	CountryName   string    `json:"country_name"`
	CountryCity   string    `json:"country_city"`
	CountryPlace  string    `json:"country_place"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type TicketResponse struct {
	TicketID      int    `json:"ticket_id"`
	Type          string `json:"type"`
	Price         int    `json:"price"`
	ContinentName string `json:"continent_name"`
	Stock         int    `json:"stock"`
	CountryName   string `json:"country_name"`
	CountryCity   string `json:"country_city"`
	CountryPlace  string `json:"country_place"`
}
