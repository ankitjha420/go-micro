package main

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"listener/event"
	"log"
	"math"
	"os"
	"time"
)

func main() {
	// connect to rabbit ->
	rabbitCon, err := connect()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer func(rabbitCon *amqp.Connection) {
		_ = rabbitCon.Close()
	}(rabbitCon)

	// listen for messages ->
	log.Println("listening and consuming rabbit messages")

	// create consumer ->
	consumer, err := event.NewConsumer(rabbitCon)
	if err != nil {
		panic(err)
	}

	// watch and consume events ->
	err = consumer.Listen([]string{"log.INFO", "log.WARNING", "log.ERROR"})
	if err != nil {
		log.Println(err)
	}
}

func connect() (*amqp.Connection, error) {
	var counts int64
	var backoff = time.Second * 1
	var connection *amqp.Connection

	for {
		c, err := amqp.Dial("amqp://guest:guest@rabbitmq")
		if err != nil {
			fmt.Println("rabbit not ready")
			counts++
		} else {
			log.Println("connected to rabbit")
			connection = c
			break
		}

		if counts > 5 {
			fmt.Println(err)
			return nil, err
		}

		backoff = time.Duration(math.Pow(float64(counts), 2)) * time.Second
		log.Printf("backing off for %d seconds", backoff)
		time.Sleep(backoff)
		continue
	}

	return connection, nil
}
