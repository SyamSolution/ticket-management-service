package repository

import (
	"github.com/DATA-DOG/go-sqlmock"
	mock_config "github.com/SyamSolution/ticket-management-service/mock/config"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"regexp"
	"testing"
	"time"
)

func TestGetAvailableTicketByContinent(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"ticket_detail_id", "type", "price", "continent_name", "stock_ticket", "stock", "stock_ordered", "country_name", "country_city", "country_place", "created_at", "updated_at"}).
		AddRow(1, "Type1", 100, "Continent1", 10, 10, 0, "Country1", "City1", "Place1", time.Now(), time.Now())

	mock.ExpectQuery("^SELECT (.+) FROM ticket_detail WHERE continent_name = \\? AND stock > 0$").WithArgs("Continent1").WillReturnRows(rows)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger := mock_config.NewMockLogger(ctrl)
	repo := NewTicketRepository(db, logger)

	tickets, err := repo.GetAvailableTicketByContinent("Continent1")
	if err != nil {
		t.Errorf("error was not expected while getting available tickets: %s", err)
	}

	assert.Len(t, tickets, 1)
	assert.Equal(t, "Type1", tickets[0].Type)
	assert.Equal(t, "Continent1", tickets[0].ContinentName)
}

func TestGetAvailableTicketByType(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"ticket_detail_id", "type", "price", "continent_name", "stock_ticket", "stock", "stock_ordered", "country_name", "country_city", "country_place", "created_at", "updated_at"}).
		AddRow(1, "Type1", 100, "Continent1", 10, 10, 0, "Country1", "City1", "Place1", time.Now(), time.Now())

	mock.ExpectQuery("^SELECT (.+) FROM ticket_detail WHERE type = \\? AND stock > 0$").WithArgs("Type1").WillReturnRows(rows)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger := mock_config.NewMockLogger(ctrl)
	repo := NewTicketRepository(db, logger)

	tickets, err := repo.GetAvailableTicketByType("Type1")
	if err != nil {
		t.Errorf("error was not expected while getting available tickets: %s", err)
	}

	assert.Len(t, tickets, 1)
	assert.Equal(t, "Type1", tickets[0].Type)
	assert.Equal(t, "Continent1", tickets[0].ContinentName)
}

func TestGetTicketByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"ticket_detail_id", "type", "price", "continent_name", "stock_ticket", "stock", "stock_ordered", "country_name", "country_city", "country_place", "created_at", "updated_at"}).
		AddRow(1, "Type1", 100, "Continent1", 10, 10, 0, "Country1", "City1", "Place1", time.Now(), time.Now())

	mock.ExpectQuery("^SELECT (.+) FROM ticket_detail WHERE ticket_detail_id = \\?$").WithArgs(1).WillReturnRows(rows)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger := mock_config.NewMockLogger(ctrl)
	repo := NewTicketRepository(db, logger)

	ticket, err := repo.GetTicketByID(1)
	if err != nil {
		t.Errorf("error was not expected while getting ticket by ID: %s", err)
	}

	assert.Equal(t, "Type1", ticket.Type)
	assert.Equal(t, "Continent1", ticket.ContinentName)
}

func TestGetTicketByContinent(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"ticket_detail_id", "type", "price", "continent_name", "stock_ticket", "stock", "stock_ordered", "country_name", "country_city", "country_place", "created_at", "updated_at"}).
		AddRow(1, "Type1", 100, "Continent1", 10, 10, 0, "Country1", "City1", "Place1", time.Now(), time.Now())

	mock.ExpectQuery("^SELECT (.+) FROM ticket_detail WHERE continent_name = \\?$").WithArgs("Continent1").WillReturnRows(rows)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger := mock_config.NewMockLogger(ctrl)
	repo := NewTicketRepository(db, logger)

	tickets, err := repo.GetTicketByContinent("Continent1")
	if err != nil {
		t.Errorf("error was not expected while getting tickets by continent: %s", err)
	}

	assert.Len(t, tickets, 1)
	assert.Equal(t, "Type1", tickets[0].Type)
	assert.Equal(t, "Continent1", tickets[0].ContinentName)
}

func TestUpdateStockCreateOrderTicket(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec(regexp.QuoteMeta("UPDATE ticket_detail SET stock = stock - ?, stock_ordered = stock_ordered + ? WHERE ticket_detail_id = ?")).
		WithArgs(10, 10, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger := mock_config.NewMockLogger(ctrl)
	repo := NewTicketRepository(db, logger)

	err = repo.UpdateStockCreateOrderTicket(1, 10)
	assert.NoError(t, err)
}

func TestUpdateStockSuccessOrderTicket(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec(regexp.QuoteMeta("UPDATE ticket_detail SET stock_ticket = stock_ticket - ? WHERE ticket_detail_id = ?")).
		WithArgs(10, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger := mock_config.NewMockLogger(ctrl)
	repo := NewTicketRepository(db, logger)

	err = repo.UpdateStockSuccessOrderTicket(1, 10)
	assert.NoError(t, err)
}

func TestUpdateStockFailOrderTicket(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec(regexp.QuoteMeta("UPDATE ticket_detail SET stock = stock + ?, stock_ordered = stock_ordered - ? WHERE ticket_detail_id = ?")).
		WithArgs(10, 10, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger := mock_config.NewMockLogger(ctrl)
	repo := NewTicketRepository(db, logger)

	err = repo.UpdateStockFailOrderTicket(1, 10)
	assert.NoError(t, err)
}

func TestGetStockTicketGroupByContinent(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"continent_name", "stock"}).
		AddRow("Continent1", 10)

	mock.ExpectQuery("^SELECT continent_name, SUM\\(stock\\) as stock FROM ticket_detail GROUP BY continent_name$").WillReturnRows(rows)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger := mock_config.NewMockLogger(ctrl)
	repo := NewTicketRepository(db, logger)

	tickets, err := repo.GetStockTicketGroupByContinent()
	if err != nil {
		t.Errorf("error was not expected while getting stock tickets by continent: %s", err)
	}

	assert.Len(t, tickets, 1)
	assert.Equal(t, "Continent1", tickets[0].Continent)
	assert.Equal(t, 10, tickets[0].Stock)
}

func TestGetTicketEventByTicketID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"ticket_detail_id", "type", "price", "stock", "continent_name", "country_city", "country_place", "event_name", "date", "description"}).
		AddRow(1, "Type1", 100, 10, "Continent1", "City1", "Place1", "Event1", time.Now(), "Description1")

	mock.ExpectQuery("^SELECT td.ticket_detail_id, td.type, td.price, td.stock, td.continent_name, td.country_city, td.country_place, e.event_name, e.date, e.description from ticket_detail td left join event e on td.event_id = e.event_id WHERE td.ticket_detail_id = \\?$").
		WithArgs(1).
		WillReturnRows(rows)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger := mock_config.NewMockLogger(ctrl)
	repo := NewTicketRepository(db, logger)

	ticketEvent, err := repo.GetTicketEventByTicketID(1)
	if err != nil {
		t.Errorf("error was not expected while getting ticket event by ticket id: %s", err)
	}

	assert.Equal(t, "Type1", ticketEvent.Type)
	assert.Equal(t, "Continent1", ticketEvent.Continent)
	assert.Equal(t, "City1", ticketEvent.CountryCity)
	assert.Equal(t, "Place1", ticketEvent.CountryPlace)
	assert.Equal(t, "Event1", ticketEvent.EventName)
	assert.Equal(t, "Description1", ticketEvent.Description)
}
