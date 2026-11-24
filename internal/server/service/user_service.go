package service

import (
	"agent/ai-agent-platform/internal/server/model"
	"agent/ai-agent-platform/internal/server/repository"
	"context"
)

type userService struct {
	userRepository repository.UserRepository
}

type UserService interface {
	GetUserByID(ctx context.Context, id string) (*model.User, error)
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{userRepository: userRepository}
}

func (s *userService) GetUserByID(ctx context.Context, id string) (*model.User, error) {
	return s.userRepository.GetUserByID(ctx, id)
}