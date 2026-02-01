package dbConnection

import (
	"context"
	"log"

	"github.com/addixit1/fiber-boilerplate/internal/config"
	"github.com/addixit1/fiber-boilerplate/internal/utils"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/event"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongo() {

	// ANSI Color codes for terminal output
	const (
		colorReset  = "\033[0m"
		colorCyan   = "\033[36m"
		colorGreen  = "\033[32m"
		colorYellow = "\033[33m"
		colorRed    = "\033[31m"
		colorBlue   = "\033[34m"
	)

	// Debug log before connection
	if config.Config.DebugStatus == "true" {
		log.Println(colorCyan + "[MongoDB Debug] Attempting to connect to database: " + config.Config.MongoDbName + colorReset)
		log.Println(colorCyan + "[MongoDB Debug] Connection URI: " + config.Config.MongoURI + colorReset)
	}

	// Prepare client options
	clientOpts := options.Client().ApplyURI(config.Config.MongoURI)

	// Enable query logging if debug mode is on
	if config.Config.DebugStatus == "true" {
		monitor := &event.CommandMonitor{
			Started: func(_ context.Context, evt *event.CommandStartedEvent) {
				log.Printf(colorYellow+"[MongoDB Query] Command: %s | DB: %s"+colorReset+"\n", evt.CommandName, evt.DatabaseName)
				log.Printf(colorBlue+"[MongoDB Query] Full Query: %v"+colorReset+"\n", evt.Command)
			},
			Succeeded: func(_ context.Context, evt *event.CommandSucceededEvent) {
				log.Printf(colorGreen+"[MongoDB Query] ✓ %s completed in %v"+colorReset+"\n",
					evt.CommandName, evt.Duration)
			},
			Failed: func(_ context.Context, evt *event.CommandFailedEvent) {
				log.Printf(colorRed+"[MongoDB Query] ✗ %s failed: %v"+colorReset+"\n",
					evt.CommandName, evt.Failure)
			},
		}
		clientOpts.SetMonitor(monitor)
	}

	err := mgm.SetDefaultConfig(
		nil,
		config.Config.MongoDbName,
		clientOpts,
	)

	if err != nil {
		log.Fatalf("Failed to setup MGM: %v", err)
	}

	// Debug log after successful connection
	if config.Config.DebugStatus == "true" {
		log.Println(colorGreen + "[MongoDB Debug] Successfully connected to database" + colorReset)
		log.Println(colorGreen + "[MongoDB Debug] Database name: " + config.Config.MongoDbName + colorReset)
	}

	debugStatus := ""
	if config.Config.DebugStatus == "true" {
		debugStatus = " (Debug: ON)"
	}

	utils.LogDatabase("MongoDB connected to " + config.Config.MongoDbName + debugStatus)
}
