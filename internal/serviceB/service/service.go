package service

import (
	"context"

	uuid "github.com/satori/go.uuid"

	"microservices-boilerplate/internal/constants"
	"microservices-boilerplate/internal/pkg/log"
	"microservices-boilerplate/internal/serviceB/domain"
	"microservices-boilerplate/internal/serviceB/repository"
)

type Service interface {
	GetAll(ctx context.Context) ([]*domain.ItemB, error)
	GetOneByID(ctx context.Context, id string) (*domain.ItemB, error)
	Create(ctx context.Context, item *domain.ItemB) (*domain.ItemB, error)
	Update(ctx context.Context, id string, item *domain.ItemB) error
	Delete(ctx context.Context, id string) error
}

func New(config *Config) Service {
	return &service{
		config: config,
	}
}

type Config struct {
	Log        log.Logger
	Repository repository.Repository
}

type service struct {
	config *Config
}

func (s *service) GetAll(ctx context.Context) ([]*domain.ItemB, error) {
	resp, err := s.config.Repository.GetAll(ctx)
	if err != nil {
		s.config.Log.Error("failed to get all item a", err)
		return nil, err
	}

	return resp, nil
}

func (s *service) GetOneByID(ctx context.Context, id string) (*domain.ItemB, error) {
	itemID, err := uuid.FromString(id)
	if err != nil {
		s.config.Log.Error("failed to parse id to UUID", err)
		return nil, constants.ErrCreatingUUIDFromString
	}

	resp, err := s.config.Repository.GetByID(ctx, itemID)
	if err != nil {
		s.config.Log.Error("failed to get item with id", itemID, err)
		return nil, err
	}

	return resp, nil
}

func (s *service) Create(ctx context.Context, item *domain.ItemB) (*domain.ItemB, error) {
	resp, err := s.config.Repository.Insert(ctx, item)
	if err != nil {
		s.config.Log.Error("failed to create item", item, err)
		return nil, err
	}

	return resp, nil
}

func (s *service) Update(ctx context.Context, id string, item *domain.ItemB) error {
	itemID, err := uuid.FromString(id)
	if err != nil {
		s.config.Log.Error("failed to parse id to UUID", err)
		return constants.ErrCreatingUUIDFromString
	}

	if err = s.config.Repository.Update(ctx, itemID, item); err != nil {
		s.config.Log.Error("failed to update item", itemID, item, err)
		return err
	}

	return nil
}

func (s *service) Delete(ctx context.Context, id string) error {
	itemID, err := uuid.FromString(id)
	if err != nil {
		s.config.Log.Error("failed to parse id to UUID", err)
		return constants.ErrCreatingUUIDFromString
	}

	if err = s.config.Repository.Remove(ctx, itemID); err != nil {
		s.config.Log.Error("failed to delete item", itemID, err)
		return err
	}

	return nil
}
