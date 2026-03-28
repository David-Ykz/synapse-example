package main

import (
	"log"
	"os"
	"strconv"
	"time"
	"fmt"

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
		message := fmt.Sprintf("Hello World: %d", counter)
		producer.Produce([]byte(message))
		time.Sleep(100 * time.Millisecond)
		counter += 1
	}
}
