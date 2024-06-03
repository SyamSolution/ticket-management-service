package usecase

import (
	"github.com/SyamSolution/ticket-management-service/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
	"testing"
)

type MockTicketPersister struct {
	mock.Mock
}

func (m *MockTicketPersister) GetAvailableTicketByContinent(continent string) ([]model.Ticket, error) {
	args := m.Called(continent)
	return args.Get(0).([]model.Ticket), args.Error(1)
}

func (m *MockTicketPersister) GetTicketByContinent(continent string) ([]model.Ticket, error) {
	args := m.Called(continent)
	return args.Get(0).([]model.Ticket), args.Error(1)
}

func (m *MockTicketPersister) UpdateStockCreateOrderTicket(ticketID, order int) error {
	args := m.Called(ticketID, order)
	return args.Error(0)
}

func (m *MockTicketPersister) UpdateStockSuccessOrderTicket(ticketID, order int) error {
	args := m.Called(ticketID, order)
	return args.Error(0)
}

func (m *MockTicketPersister) UpdateStockFailOrderTicket(ticketID, order int) error {
	args := m.Called(ticketID, order)
	return args.Error(0)
}

func (m *MockTicketPersister) GetStockTicketGroupByContinent() ([]model.StockTicket, error) {
	args := m.Called()
	return args.Get(0).([]model.StockTicket), args.Error(1)
}

func (m *MockTicketPersister) GetTicketEventByTicketID(ticketID int) (model.TicketEvent, error) {
	args := m.Called(ticketID)
	return args.Get(0).(model.TicketEvent), args.Error(1)
}

func (m *MockTicketPersister) GetAvailableTicketByType(ticketType string) ([]model.Ticket, error) {
	args := m.Called(ticketType)
	return args.Get(0).([]model.Ticket), args.Error(1)
}

func (m *MockTicketPersister) GetTicketByID(ticketID int) (model.Ticket, error) {
	args := m.Called(ticketID)
	return args.Get(0).(model.Ticket), args.Error(1)

}

func TestTicketUsecase(t *testing.T) {
	mockRepo := new(MockTicketPersister)
	mockLogger := zap.NewNop()
	ticketUsecase := NewTicketUsecase(mockRepo, mockLogger)

	// Define your mock data here
	mockTickets := []model.Ticket{
		{
			TicketID:      1,
			Type:          "Type1",
			Price:         100,
			ContinentName: "Asia",
			Stock:         10,
			CountryName:   "Country1",
			CountryCity:   "City1",
			CountryPlace:  "Place1",
		},
	}

	// Set up the mock responses
	mockRepo.On("GetAvailableTicketByContinent", "Asia").Return(mockTickets, nil)
	mockRepo.On("GetTicketByContinent", "Asia").Return(mockTickets, nil)
	mockRepo.On("GetStockTicketGroupByContinent").Return([]model.StockTicket{}, nil)
	mockRepo.On("GetTicketEventByTicketID", 1).Return(model.TicketEvent{}, nil)
	mockRepo.On("GetAvailableTicketByType", "Type1").Return(mockTickets, nil)

	// Call the methods and assert the results
	tickets, err := ticketUsecase.GetAvailableTicketByContinent("Asia")
	assert.NoError(t, err)
	assert.NotNil(t, tickets)

	tickets, err = ticketUsecase.GetTicketByContinent("Asia")
	assert.NoError(t, err)
	assert.NotNil(t, tickets)

	stockTickets, err := ticketUsecase.GetStockTicketGroupByContinent()
	assert.NoError(t, err)
	assert.NotNil(t, stockTickets)

	ticketEvent, err := ticketUsecase.GetTicketEventByTicketID(1)
	assert.NoError(t, err)
	assert.NotNil(t, ticketEvent)

	tickets, err = ticketUsecase.GetAvailableTicketByType("Type1")
	assert.NoError(t, err)
	assert.NotNil(t, tickets)

	// Assert that the mock expectations were met
	mockRepo.AssertExpectations(t)
}
