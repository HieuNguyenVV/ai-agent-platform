package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	clientContainer     *ClientContainer
	controllerContainer *ControllerContainer
	serviceContainer    *ServiceContainer
	repositoryContainer *RepositoryContainer
}

func NewServer(
	clientContainer *ClientContainer,
	controllerContainer *ControllerContainer,
	serviceContainer *ServiceContainer,
	repositoryContainer *RepositoryContainer,
) *Server {
	engine := clientContainer.httpClient
	if clientContainer.config.App.Debug {
		engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	{
		engine.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "OK"})
		})

		engine.GET("/ready", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "OK"})
		})
	}

	return &Server{
		clientContainer:     clientContainer,
		controllerContainer: controllerContainer,
		serviceContainer:    serviceContainer,
		repositoryContainer: repositoryContainer,
	}
}

func (s *Server) Start(ctx context.Context) {
	server := &http.Server{
		Addr:              fmt.Sprintf(":%s", s.clientContainer.config.App.Port),
		Handler:           s.clientContainer.httpClient,
		ReadHeaderTimeout: 10 * time.Second,
	}

	go func() {
		s.clientContainer.logger.Infof("Starting server on port %s", s.clientContainer.config.App.Port)
		s.clientContainer.logger.Infof("Swagger UI: http://localhost:%s/swagger/index.html", s.clientContainer.config.App.Port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed 	{
			s.clientContainer.logger.WithError(err).Error("failed to start server")
		}
	}()

	<-ctx.Done()

	s.clientContainer.logger.Info("Shutting down server")
	shutdownCtx, shutdownCancel:= context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownCancel()

	if err := server.Shutdown(shutdownCtx); err!=nil{
		s.clientContainer.logger.WithError(err).Error("failed to shutdown server")
	}

	s.clientContainer.postgres.Shutdown()

	if err := s.clientContainer.redis.Shutdown(); err != nil {
		s.clientContainer.logger.WithError(err).Error("failed to shutdown redis")
	}
	s.clientContainer.logger.Info("server shutdown completed")
}
