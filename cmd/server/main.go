package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"agent/ai-agent-platform/internal/pkg/config"
	"agent/ai-agent-platform/internal/server"

	_ "agent/ai-agent-platform/docs"
)

// @title AI Agent Platform
// @version 1.0
// @description AI Agent Platform
// @contact.name API Support
// @contact.url https://github.com/AI-Agent-Platform/AI-Agent-Platform
// @contact.email support@AI-Agent-Platform.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @basePath /api/v1
// @schemes http
// @securityDefinitions.apikey AccessToken
// @name Authorization
// @in header
// @securityDefinitions.apikey RefreshToken
// @name Authorization
// @in header
func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func() {
		signals := make(chan os.Signal, 1)
		signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
		<-signals
		cancel()
	}()

	clientContainer := server.NewClientContainer(config)
	repositoryContainer := server.NewRepositoryContainer(clientContainer)
	serviceContainer := server.NewServiceContainer(clientContainer, repositoryContainer)
	controllerContainer := server.NewControllerContainer(clientContainer, serviceContainer)
	server := server.NewServer(clientContainer, controllerContainer, serviceContainer, repositoryContainer)
	server.Start(ctx)
}
