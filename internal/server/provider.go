package server

import (
	"agent/ai-agent-platform/internal/pkg/cache"
	"agent/ai-agent-platform/internal/pkg/config"
	"agent/ai-agent-platform/internal/pkg/database"
	"agent/ai-agent-platform/internal/pkg/log"
	"agent/ai-agent-platform/internal/server/controller"
	"agent/ai-agent-platform/internal/server/repository"
	"agent/ai-agent-platform/internal/server/service"

	"github.com/gin-gonic/gin"
)

type ClientContainer struct {
	config     *config.Config
	logger     log.Logger
	httpClient *gin.Engine
	postgres   *database.Postgres
	redis      *cache.Redis
}

type ControllerContainer struct {
	userController controller.UserController
}

type ServiceContainer struct {
	userService service.UserService
}

type RepositoryContainer struct {
	userRepository repository.UserRepository
}

func NewClientContainer(cfg *config.Config) *ClientContainer {
	logger := log.New(&log.Config{
		LogLevel: cfg.Log.Level,
	})
	postgres, err := database.NewPostgres(cfg.Postgres)
	if err != nil {
		logger.Fatalf("NewClientContainer: failed to create postgres: %v", err)
	}
	redis, err := cache.NewRedis(cfg.Redis)
	if err != nil {
		logger.Fatalf("NewClientContainer: failed to create redis: %v", err)
	}
	return &ClientContainer{
		config:     cfg,
		logger:     logger,
		httpClient: NewHttpServer(cfg),
		postgres:   postgres,
		redis:      redis,
	}
}

func NewControllerContainer(clientContainer *ClientContainer, serviceContainer *ServiceContainer) *ControllerContainer {
	return &ControllerContainer{
		userController: controller.NewUserController(serviceContainer.userService),
	}
}

func NewServiceContainer(clientContainer *ClientContainer, repositoryContainer *RepositoryContainer) *ServiceContainer {
	return &ServiceContainer{
		userService: service.NewUserService(repositoryContainer.userRepository),
	}
}

func NewRepositoryContainer(clientContainer *ClientContainer) *RepositoryContainer {
	return &RepositoryContainer{
		userRepository: repository.NewUserRepository(clientContainer.postgres),
	}
}

func NewHttpServer(cfg *config.Config) *gin.Engine {
	var engine *gin.Engine
	if cfg.App.Debug {
		gin.SetMode(gin.DebugMode)
		engine = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		engine = gin.New()
	}
	return engine
}
