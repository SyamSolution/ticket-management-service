package consumer

import (
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
	"github.com/SyamSolution/ticket-management-service/helper"
	"github.com/SyamSolution/ticket-management-service/internal/model"
	"github.com/SyamSolution/ticket-management-service/internal/usecase"
	"log"
	"os"
)

func Consumer(master sarama.Consumer, doneCh chan struct{}, ticketUsecase usecase.TicketExecutor) {
	consumer, errors := helper.Consume(master, []string{"order-ticket", "success-order-ticket", "failed-order-ticket"})

	signals := make(chan os.Signal, 1)
	for {
		select {
		case msg := <-consumer:
			switch msg.Topic {
			case "order-ticket":
				var message model.MessageOrderTicket
				if err := json.Unmarshal(msg.Value, &message); err != nil {
					log.Printf("Error unmarshalling message: %s\n", err)
				}

				if err := ticketUsecase.UpdateStockTicket(message.TicketID, message.Order, "create"); err != nil {
					log.Printf("Error when ordering ticket: %s\n", err)
				} else {
					log.Printf("Order ticket with ticketID: %d and order: %d success\n", message.TicketID, message.Order)
				}
			case "success-order-ticket":
				var message model.MessageOrderTicket
				if err := json.Unmarshal(msg.Value, &message); err != nil {
					log.Printf("Error unmarshalling message: %s\n", err)
				}

				if err := ticketUsecase.UpdateStockTicket(message.TicketID, message.Order, "success"); err != nil {
					log.Printf("Error when success ordering ticket: %s\n", err)
				} else {
					log.Printf("Success order ticket with ticketID: %d and order: %d success\n", message.TicketID, message.Order)
				}
			case "failed-order-ticket":
				var message model.MessageOrderTicket
				if err := json.Unmarshal(msg.Value, &message); err != nil {
					log.Printf("Error unmarshalling message: %s\n", err)
				}

				if err := ticketUsecase.UpdateStockTicket(message.TicketID, message.Order, "failed"); err != nil {
					log.Printf("Error when failed ordering ticket: %s\n", err)
				} else {
					log.Printf("Failed order ticket with ticketID: %d and order: %d success\n", message.TicketID, message.Order)
				}
			}
		case consumerError := <-errors:
			fmt.Println("Received consumer error", (consumerError).Error())
		case <-signals:
			fmt.Println("Interrupt is detected")
			doneCh <- struct{}{}
		}
	}
}
