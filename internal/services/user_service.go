package services

import (
	"context"
	"user-service/internal/db"
)

type UserService interface {
	FindUserByID(ctx context.Context, id int) (int, error)
}

type userService struct {
	queries *db.Queries
}

func NewUserService(queries *db.Queries) UserService {
	return &userService{
		queries: queries,
	}
}

func (us *userService) FindUserByID(ctx context.Context, id int) (int, error) {
	_, err := us.queries.CheckUserExists(ctx, "vy@gmail.com")
	if err != nil {
		return 0, err
	}
	return id, nil
}
