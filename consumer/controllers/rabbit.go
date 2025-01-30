package controllers

import (
	"encoding/json"
	"gateway-router-consumer/utils"
	"log"

	"github.com/rabbitmq/amqp091-go"
)

var amqpURI = "amqp://guest:guest@localhost:5672/" // RabbitMQ URI

type TaskMessage struct {
	TaskID   int64       `json:"task_id"`
	TaskType string      `json:"task_type"`
	Data     interface{} `json:"data"`
	Action   string      `json:"action"`
}
type Response struct {
	TaskID   int64  `json:"task_id"`
	TaskType string `json:"task_type"`
	Status   string `json:"status"`
}

func StartConsumer() {
	conn, err := amqp091.Dial(amqpURI) // Updated method
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %s", err)
	}
	defer conn.Close()
	log.Println("Successfully connected to rabbitmq..")
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %s", err)
	}
	defer ch.Close()
	// Declare the queue
	queue, err := ch.QueueDeclare(
		"task_queue", // Queue name
		true,         // Durable
		false,        // Delete when unused
		false,        // Exclusive
		false,        // NoWait
		nil,          // Arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %s", err)
	}
	// Start consuming from the queue
	msgs, err := ch.Consume(
		queue.Name, // Queue name
		"",         // Consumer name
		true,       // AutoAck
		false,      // Exclusive
		false,      // NoLocal
		false,      // NoWait
		nil,        // Arguments
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %s", err)
	}
	log.Println("Waiting for messages....")
	// Consume messages
	for msg := range msgs {
		var task TaskMessage
		if err := json.Unmarshal(msg.Body, &task); err != nil {
			log.Printf("Failed to unmarshal task message: %v", err)
			continue
		}
		if err != nil {
			log.Printf("Failed to unmarshal message: %s", err)
			continue
		}
		go handleMessages(ch, task)
	}
}

func sendResponse(ch *amqp091.Channel, response Response) {
	body, _ := json.Marshal(response)

	ch.Publish(
		"",
		"response_queue",
		false,
		false,
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
}

func handleMessages(ch *amqp091.Channel, task TaskMessage) {
	// res := Response{
	// 	TaskID:   task.TaskID,
	// 	TaskType: task.TaskType,
	// }
	var status string
	log.Println(status)
	log.Println(task.Data)
	res := Response{
		TaskID:   task.TaskID,
		TaskType: task.TaskType,
	}
	switch task.Action {
	case "create":
		log.Println("create")
		if err := utils.Process(task.Data, task.Action); err != nil {
			log.Println(err)
			res.Status = "failed"
		} else {
			res.Status = "created"
		}
	case "delete":
		log.Println("delete")
		if err := utils.Process(task.Data, task.Action); err != nil {
			log.Println(err)
			res.Status = "failed"
		} else {
			res.Status = "deleted"
		}

	}
	log.Println(res)
	sendResponse(ch, res)
}
