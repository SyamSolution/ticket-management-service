package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/SyamSolution/ticket-management-service/internal/model"
	"github.com/SyamSolution/ticket-management-service/internal/usecase"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

const workerCount = 5
func processTicketSuccessUpdate(message types.Message) {
    fmt.Printf("Update Ticket Success Queue - Message ID: %s\n", *message.MessageId)
    fmt.Printf("Update Ticket Success Queue - Message Body: %s\n", *message.Body)
}

func processDeadLetterMessage(message types.Message) {
    fmt.Printf("Dead Letter Queue - Message ID: %s\n", *message.MessageId)
    fmt.Printf("Dead Letter Queue - Message Body: %s\n", *message.Body)
}

func workerDeadLetter(client *sqs.Client, queueURL string, wg *sync.WaitGroup, ticketUsecase usecase.TicketExecutor, status string) {
    defer wg.Done()
    for {
        result, err := client.ReceiveMessage(context.TODO(), &sqs.ReceiveMessageInput{
            QueueUrl:            &queueURL,
            MaxNumberOfMessages: 10,
            WaitTimeSeconds:     10,
        })
        if err != nil {
            log.Printf("failed to receive messages from Dead Letter queue, %v", err)
            continue
        }

        if len(result.Messages) == 0 {
            time.Sleep(1 * time.Second)
            continue
        }

        for _, message := range result.Messages {

            var msg model.MessageOrderTicket
            err := json.Unmarshal([]byte(*message.Body), &msg)
            if err != nil {
                fmt.Println("Error unmarshalling message", err)
            }else{
                log.Printf("consume DLQ %s ticket", status)
                if err := ticketUsecase.UpdateStockTicket(msg.TicketID, msg.Order, status); err != nil {
					log.Printf("Error when update status %s ticket: %s\n", status, err)
				} else {
					log.Printf("%s order ticket with ticketID: %d and order: %d success\n", status, msg.TicketID, msg.Order)
				}
            }
  
            processDeadLetterMessage(message)

            _, err = client.DeleteMessage(context.TODO(), &sqs.DeleteMessageInput{
                QueueUrl:      &queueURL,
                ReceiptHandle: message.ReceiptHandle,
            })
            if err != nil {
                log.Printf("failed to delete message from Dead Letter queue, %v", err)
            }
        }
    }
}

func workerTicket(client *sqs.Client, queueURL string, wg *sync.WaitGroup, ticketUsecase usecase.TicketExecutor, status string) {
    defer wg.Done()
    for {
        result, err := client.ReceiveMessage(context.TODO(), &sqs.ReceiveMessageInput{
            QueueUrl:            &queueURL,
            MaxNumberOfMessages: 10,
            WaitTimeSeconds:     10,
        })
        if err != nil {
            log.Printf("failed to receive messages from SendEmailPdf queue, %v", err)
            continue
        }

        if len(result.Messages) == 0 {
            time.Sleep(1 * time.Second)
            continue
        }

        for _, message := range result.Messages {
            var msg model.MessageOrderTicket
            err := json.Unmarshal([]byte(*message.Body), &msg)
            if err != nil {
                fmt.Println("Error unmarshalling message", err)
            }else{
                log.Printf("consume %s ticket", status)
                if err := ticketUsecase.UpdateStockTicket(msg.TicketID, msg.Order, status); err != nil {
                    log.Printf("Error when update status %s ticket: %s\n", status, err)
                } else {
					log.Printf("%s order ticket with ticketID: %d and order: %d success\n", status, msg.TicketID, msg.Order)
				}
            }
            processTicketSuccessUpdate(message)

            _, err = client.DeleteMessage(context.TODO(), &sqs.DeleteMessageInput{
                QueueUrl:      &queueURL,
                ReceiptHandle: message.ReceiptHandle,
            })
            if err != nil {
                log.Printf("failed to delete message from SendEmailPdf queue, %v", err)
            }
        }
    }
}

func StartConsumer(ticketUsecase usecase.TicketExecutor) {
    cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("ap-southeast-1"))
    if err != nil {
        log.Fatalf("unable to load SDK config, %v", err)
    }

    client := sqs.NewFromConfig(cfg)

    var wg sync.WaitGroup

    for i := 0; i < workerCount; i++ {
        wg.Add(1)
        go workerTicket(client, os.Getenv("SQS_TICKET_SUCCESS_URL"), &wg, ticketUsecase, "success")
    }

    for i := 0; i < workerCount; i++ {
        wg.Add(1)
        go workerTicket(client, os.Getenv("SQS_TICKET_FAILED_URL"), &wg, ticketUsecase, "failed")
   }

    for i := 0; i < workerCount; i++ {
        wg.Add(1)
        go workerTicket(client, os.Getenv("SQS_TICKET_URL"), &wg, ticketUsecase, "create")
    }

    for i := 0; i < workerCount; i++ {
        wg.Add(1)
        go workerDeadLetter(client, os.Getenv("SQS_TICKET_DLQ_URL"), &wg, ticketUsecase, "create")
    }

    for i := 0; i < workerCount; i++ {
        wg.Add(1)
        go workerDeadLetter(client, os.Getenv("SQS_TICKET_FAILED_DLQ_URL"), &wg, ticketUsecase,  "failed")
    }

    for i := 0; i < workerCount; i++ {
        wg.Add(1)
        go workerDeadLetter(client, os.Getenv("SQS_TICKET_SUCCESS_DLQ_URL"), &wg, ticketUsecase,  "success")
    }

    wg.Wait()
}
