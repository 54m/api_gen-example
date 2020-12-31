package userimpl

import (
	"context"

	"github.com/54m/api_gen-example/backend/domain/model"
)

// UserImpl - user implements
type UserImpl struct{}

// NewUserImpl - constructor
func NewUserImpl() *UserImpl {
	return new(UserImpl)
}

// Get - user acquisition
func (u *UserImpl) Get(_ context.Context, id string) (*model.User, error) {
	// FIXME: implement me
	return &model.User{
		ID:     id,
		Age:    100,
		Name:   "54m",
		Gender: model.GenderMale,
	}, nil
}

// Insert - user insertion
func (u *UserImpl) Insert(_ context.Context, user *model.User) error {
	// FIXME: implement me
	user.ID = "xxx"
	return nil
}

// Update - user update
func (u *UserImpl) Update(_ context.Context, _ *model.User) error {
	// FIXME: implement me
	return nil
}

// DeleteByID - user deletion by id
func (u *UserImpl) DeleteByID(_ context.Context, _ string) error {
	// FIXME: implement me
	return nil
}
