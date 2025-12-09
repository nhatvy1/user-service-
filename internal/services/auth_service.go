package services

import (
	"context"
	"fmt"
	"user-service/internal/db"
	"user-service/internal/types/response"
	"user-service/internal/utils"
	"user-service/internal/vo"
)

type AuthService interface {
	Login(ctx context.Context, loginBody *vo.LoginRequest) (response.UserLoginResponse, error)
}

type authService struct {
	queries *db.Queries
}

func NewAuthService(queries *db.Queries) AuthService {
	return &authService{
		queries: queries,
	}
}

func (as *authService) Login(ctx context.Context, loginBody *vo.LoginRequest) (response.UserLoginResponse, error) {
	user, err := as.queries.GetUserLoginInfo(ctx, loginBody.Email)

	if err != nil {
		return response.UserLoginResponse{}, fmt.Errorf("email or password incorrect")
	}

	if isValidPassword := utils.ComparePassword(loginBody.Password, user.Password.String); !isValidPassword {
		return response.UserLoginResponse{}, fmt.Errorf("email or password incorrect")
	}

	accessToken, err := utils.GenerateJWTToken(int(user.ID))
	if err != nil {
		return response.UserLoginResponse{}, fmt.Errorf("failed to generate access token: %w", err)
	}

	return response.UserLoginResponse{UserId: int(user.ID), AccessToken: accessToken}, nil
}
