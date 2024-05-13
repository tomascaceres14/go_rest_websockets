package repository

import (
	"context"

	"github.com/tomascaceres14/go_rest_websockets/models"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUserById(ctx context.Context, id int64) (*models.User, error)
	Close() error
}

var implementation UserRepository

func SetRepository(repository UserRepository) {
	implementation = repository
}

func CreateUser(ctx context.Context, user *models.User) error {
	return implementation.CreateUser(ctx, user)
}

func GetUserById(ctx context.Context, id int64) (*models.User, error) {
	return implementation.GetUserById(ctx, id)
}

func Close() error {
	return implementation.Close()
}
