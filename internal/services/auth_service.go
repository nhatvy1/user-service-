package services

import (
	"context"
	"user-service/internal/db"
	"user-service/internal/types/response"
	"user-service/internal/utils"
	"user-service/internal/vo"

	"github.com/jackc/pgx/v5/pgtype"
)

type AuthService interface {
	Login(ctx context.Context, loginBody *vo.LoginRequest) (response.UserLoginResponse, error)
	Register(ctx context.Context, registerBody *vo.RegisterRequest) (response.UserLoginResponse, error)
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
		return response.UserLoginResponse{}, utils.NewAppError(401, "email or password incorrect")
	}

	if isValidPassword := utils.ComparePassword(loginBody.Password, user.Password.String); !isValidPassword {
		return response.UserLoginResponse{}, utils.NewAppError(401, "email or password incorrect")
	}

	accessToken, err := utils.GenerateJWTToken(int(user.ID))
	if err != nil {
		return response.UserLoginResponse{}, utils.NewAppError(500, "please try again later")
	}

	return response.UserLoginResponse{AccessToken: accessToken}, nil
}

func (as *authService) Register(ctx context.Context, registerBody *vo.RegisterRequest) (response.UserLoginResponse, error) {
	userExists, err := as.queries.CheckUserExists(ctx, registerBody.Email)
	if err != nil {
		return response.UserLoginResponse{}, utils.NewAppError(500, "please try again later")
	}

	if userExists {
		return response.UserLoginResponse{}, utils.NewAppError(409, "user with this email already exists")
	}

	hashedPassword, err := utils.HashPassword(registerBody.Password)
	if err != nil {
		return response.UserLoginResponse{}, utils.NewAppError(500, "please try again later")
	}

	createUserParams := db.UserRegisterParams{
		Email:     registerBody.Email,
		Password:  utils.NullableText(&hashedPassword),
		AuthType:  0,
		Firstname: registerBody.FirstName,
		Lastname:  registerBody.LastName,
		Enabled:   pgtype.Bool{Bool: false, Valid: true},
		Secret:    pgtype.Text{String: "", Valid: false},
	}

	userId, err := as.queries.UserRegister(ctx, createUserParams)
	if err != nil {
		return response.UserLoginResponse{}, utils.NewAppError(500, "please try again later")
	}

	accessToken, err := utils.GenerateJWTToken(int(userId))
	if err != nil {
		return response.UserLoginResponse{}, utils.NewAppError(500, "please try again later")
	}

	return response.UserLoginResponse{AccessToken: accessToken}, nil
}
