package handler

import (
	"github.com/SyamSolution/ticket-management-service/internal/model"
	"github.com/SyamSolution/ticket-management-service/mock"
	mock_config "github.com/SyamSolution/ticket-management-service/mock/config"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"net/http/httptest"
	"testing"
)

func TestGetAvailableTicketByContinent(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTicketUsecase := mock.NewMockTicketExecutor(ctrl)
	mockLogger := mock_config.NewMockLogger(ctrl)

	handler := NewTicketHandler(mockTicketUsecase, mockLogger)

	expectedTickets := []model.TicketResponse{{}}

	continent := "Asia"

	mockTicketUsecase.EXPECT().GetAvailableTicketByContinent(continent).Return(expectedTickets, nil)

	app := fiber.New()
	app.Get("/tickets/:continent", handler.GetAvailableTicketByContinent)

	req := httptest.NewRequest("GET", "/tickets/"+continent, nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestGetAvailableTicketByType(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTicketUsecase := mock.NewMockTicketExecutor(ctrl)
	mockLogger := mock_config.NewMockLogger(ctrl)

	handler := NewTicketHandler(mockTicketUsecase, mockLogger)

	expectedTickets := []model.TicketResponse{{}}

	ticketType := "Type1"

	mockTicketUsecase.EXPECT().GetAvailableTicketByType(ticketType).Return(expectedTickets, nil)

	app := fiber.New()
	app.Get("/tickets/type/:type", handler.GetAvailableTicketByType)

	req := httptest.NewRequest("GET", "/tickets/type/"+ticketType, nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestGetTicketByContinent(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTicketUsecase := mock.NewMockTicketExecutor(ctrl)
	mockLogger := mock_config.NewMockLogger(ctrl)

	handler := NewTicketHandler(mockTicketUsecase, mockLogger)

	expectedTickets := []model.TicketResponse{{}}

	continent := "Asia"

	mockTicketUsecase.EXPECT().GetTicketByContinent(continent).Return(expectedTickets, nil)

	app := fiber.New()
	app.Get("/tickets/:continent", handler.GetTicketByContinent)

	req := httptest.NewRequest("GET", "/tickets/"+continent, nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestGetStockTicketGroupByContinent(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTicketUsecase := mock.NewMockTicketExecutor(ctrl)
	mockLogger := mock_config.NewMockLogger(ctrl)

	handler := NewTicketHandler(mockTicketUsecase, mockLogger)

	expectedTickets := []model.StockTicket{{}}

	mockTicketUsecase.EXPECT().GetStockTicketGroupByContinent().Return(expectedTickets, nil)

	app := fiber.New()
	app.Get("/tickets/stock", handler.GetStockTicketGroupByContinent)

	req := httptest.NewRequest("GET", "/tickets/stock", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestGetTicketEventByTicketID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTicketUsecase := mock.NewMockTicketExecutor(ctrl)
	mockLogger := mock_config.NewMockLogger(ctrl)

	handler := NewTicketHandler(mockTicketUsecase, mockLogger)

	expectedTicketEvent := model.TicketEvent{} // Fill with expected data

	ticketID := "1"

	mockTicketUsecase.EXPECT().GetTicketEventByTicketID(1).Return(expectedTicketEvent, nil)

	app := fiber.New()
	app.Get("/tickets/event/:ticket_id", handler.GetTicketEventByTicketID)

	req := httptest.NewRequest("GET", "/tickets/event/"+ticketID, nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}
