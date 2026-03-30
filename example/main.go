package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/David-Ykz/synapse/client"
	"github.com/David-Ykz/synapse/common"
)

var (
	producerConfig = common.Config{
		Host:        "localhost",
		Port:        8080,
		Namespace:   "example-input",
		ChannelSize: 128,
	}
)

func main() {
	brokerHost := os.Getenv("BROKER_HOST")
	if brokerHost != "" {
		producerConfig.Host = brokerHost
	}
	brokerPort := os.Getenv("BROKER_PORT")
	if brokerPort != "" {
		producerConfig.Port, _ = strconv.Atoi(brokerPort)
	}

	// initialize producer
	producer := client.NewProducer(producerConfig)
	if err := producer.Connect(); err != nil {
		log.Fatalf("Failed to connect producer: %v", err)
	}
	defer producer.Disconnect()

	counter := 0
	for {
		if counter < 1 {
			// message := fmt.Sprintf("What is %d squared", counter)
			message := "How do i get from toronto to san francisco. Use the function call"
			producer.Produce([]byte(message))
			// time.Sleep(100 * time.Millisecond)
			time.Sleep(1 * time.Second)
			counter += 1
		}
	}
}
