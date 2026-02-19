package main

import (
	"context"
	"fmt"
	"time"

	"github.com/ashcycling/docker-bake-tst/internal/env"
	"github.com/ashcycling/docker-bake-tst/internal/qrclogs"
	glide "github.com/valkey-io/valkey-glide/go/v2"
	"github.com/valkey-io/valkey-glide/go/v2/config"
	"go.uber.org/zap"
)

var client *glide.Client
var err error

func main() {

	logger := qrclogs.CreateLogger()
	valkeyConnections := env.GetValkeyConnectionCredentials()

	config := config.NewClientConfiguration().
		WithAddress(&config.NodeAddress{Host: valkeyConnections.Host, Port: valkeyConnections.Port})

	for {
		client, err = glide.NewClient(config)
		if err == nil {
			res, err := client.Ping(context.Background())
			if err == nil && res == "PONG" {
				logger.Info("Successfully connected to Valkey", zap.String("response", res))
				break // Successfully connected
			}
		}
		// Log error and optionally add a retry limit or sleep
		logger.Error("Error connecting to Valkey", zap.Error(err))
		time.Sleep(10 * time.Second)
	}
	defer client.Close()

	for {

		set_name := "unique"

		set_messages, err := client.SMembers(context.Background(), set_name)
		if err != nil {
			logger.Error("Error getting set members", zap.Error(err))
		}
		logger.Info("Members in Unique Set is ", zap.String("unique Set members", fmt.Sprintf("%v", set_messages)))
		// for _, msg := range set_messages {
		// 	logger.Info("Pending messages in queue", zap.String("message", fmt.Sprintf("%v", msg)))
		// }

		pending_list_name := "pending_queue"
		list_messages_length, err := client.LLen(context.Background(), pending_list_name)
		if err != nil {
			logger.Error("Error getting set members", zap.Error(err))
		}
		logger.Info("Quantity of messages in pending List is", zap.String("message quantity", fmt.Sprintf("%v", list_messages_length)))
	}

}
