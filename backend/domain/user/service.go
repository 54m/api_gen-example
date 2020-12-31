package user

import (
	"context"

	"github.com/54m/api_gen-example/backend/domain/model"
	"golang.org/x/xerrors"
)

// Service - user service
type Service struct {
	repository Repository
}

// NewService - constructor
func NewService(repo Repository) *Service {
	return &Service{repository: repo}
}

// Get - user acquisition
func (s *Service) Get(ctx context.Context, id string) (*model.User, error) {
	user, err := s.repository.Get(ctx, id)
	if err != nil {
		return nil, xerrors.Errorf("failed to get: %w", err)
	}
	return user, nil
}

// Insert - user insertion
func (s *Service) Insert(ctx context.Context, user *model.User) error {
	if err := s.repository.Insert(ctx, user); err != nil {
		return xerrors.Errorf("failed to insertion: %w", err)
	}
	return nil
}

// Update - user update
func (s *Service) Update(ctx context.Context, user *model.User) error {
	if err := s.repository.Update(ctx, user); err != nil {
		return xerrors.Errorf("failed to update: %w", err)
	}
	return nil
}

// DeleteByID - user deletion by id
func (s *Service) DeleteByID(ctx context.Context, id string) error {
	if err := s.repository.DeleteByID(ctx, id); err != nil {
		return xerrors.Errorf("failed to insertion: %w", err)
	}
	return nil
}

// AgeIncrement - user age incrementation
func (s *Service) AgeIncrement(ctx context.Context, id string) (*model.User, error) {
	user, err := s.Get(ctx, id)
	if err != nil {
		return nil, xerrors.Errorf("failed to get: %w", err)
	}

	user.Age++

	if err := s.Update(ctx, user); err != nil {
		return nil, xerrors.Errorf("failed to age incrementation: %w", err)
	}
	return user, nil
}
