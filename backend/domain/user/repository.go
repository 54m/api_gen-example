package user

import (
	"context"

	"github.com/54m/api_gen-example/backend/domain/model"
)

// Repository - User repository
type Repository interface {
	Get(ctx context.Context, id string) (*model.User, error)
	Insert(ctx context.Context, user *model.User) error
	Update(ctx context.Context, user *model.User) error
	DeleteByID(ctx context.Context, id string) error
}
