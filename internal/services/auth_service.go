package services

import (
	"context"
	"fmt"
	"user-service/internal/db"
	"user-service/internal/types/response"
	"user-service/internal/utils"
	"user-service/internal/vo"
	"user-service/pkg/cache"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type AuthService interface {
	Login(ctx context.Context, loginBody *vo.LoginRequest) (response.UserLoginResponse, error)
	Register(ctx context.Context, registerBody *vo.RegisterRequest) (response.UserRegisterResponse, error)
	VerifyOTP(ctx context.Context, otpBody *vo.RegisterVerifyOTP) (response.UserVerifyOTPResponse, error)
}

type authService struct {
	queries *db.Queries
	cache   cache.Cache
}

func NewAuthService(queries *db.Queries, cache cache.Cache) AuthService {
	return &authService{
		queries: queries,
		cache:   cache,
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

	isVerified := user.Verify.Bool
	if !isVerified {
		return response.UserLoginResponse{}, utils.NewAppError(401, "please verify your account first")
	}

	accessToken, err := utils.GenerateJWTToken(int(user.ID))
	if err != nil {
		return response.UserLoginResponse{}, utils.NewAppError(500, "please try again later")
	}

	return response.UserLoginResponse{AccessToken: accessToken}, nil
}

func (as *authService) Register(ctx context.Context, registerBody *vo.RegisterRequest) (response.UserRegisterResponse, error) {
	userExists, err := as.queries.CheckUserExists(ctx, registerBody.Email)
	if err != nil {
		return response.UserRegisterResponse{}, utils.NewAppError(500, "please try again later")
	}

	if userExists {
		return response.UserRegisterResponse{}, utils.NewAppError(409, "user with this email already exists")
	}

	hashedPassword, err := utils.HashPassword(registerBody.Password)
	if err != nil {
		return response.UserRegisterResponse{}, utils.NewAppError(500, "please try again later")
	}

	createUserParams := db.UserRegisterParams{
		Email:     registerBody.Email,
		Password:  utils.NullableText(&hashedPassword),
		Firstname: utils.OptionalText(&registerBody.FirstName),
		Lastname:  utils.OptionalText(&registerBody.LastName),
		Verify:    pgtype.Bool{Bool: false, Valid: true},
	}

	userId, err := as.queries.UserRegister(ctx, createUserParams)
	if err != nil {
		return response.UserRegisterResponse{}, utils.NewAppError(500, "please try again later")
	}

	otp, err := utils.GenerateNumericOTP(6)
	if err != nil {
		return response.UserRegisterResponse{}, utils.NewAppError(500, "please try again later")
	}

	verificationId := uuid.New()
	key := fmt.Sprintf("otp:%s", verificationId.String())
	jsonData := response.OTPData{
		UserId: userId,
		OTP:    otp,
	}
	err = as.cache.Set(ctx, key, jsonData, 3600)
	if err != nil {
		return response.UserRegisterResponse{}, utils.NewAppError(500, "please try again later")
	}

	return response.UserRegisterResponse{VerificationId: verificationId.String()}, nil
}

func (as *authService) VerifyOTP(ctx context.Context, otpBody *vo.RegisterVerifyOTP) (response.UserVerifyOTPResponse, error) {

	key := fmt.Sprintf("otp:%s", otpBody.VerificationId)
	var cacheOTP response.OTPData
	err := as.cache.Get(ctx, key, &cacheOTP)

	if err != nil {
		return response.UserVerifyOTPResponse{}, utils.NewAppError(400, "invalid verification id or otp")
	}

	if cacheOTP.OTP == "" {
		return response.UserVerifyOTPResponse{}, utils.NewAppError(400, "OTP incorrect")
	}

	if cacheOTP.OTP != otpBody.OTP {
		return response.UserVerifyOTPResponse{}, utils.NewAppError(400, "OTP incorrect")
	}

	// delete key after successful verification (mandatory)
	_ = as.cache.Clear(ctx, key)

	err = as.queries.UpdateUserVerifiedStatus(ctx, cacheOTP.UserId)
	if err != nil {
		return response.UserVerifyOTPResponse{}, utils.NewAppError(500, "please try again later")
	}

	accessToken, err := utils.GenerateJWTToken(int(cacheOTP.UserId))
	if err != nil {
		return response.UserVerifyOTPResponse{}, err
	}

	res := response.UserVerifyOTPResponse{
		AccessToken: accessToken,
	}

	return res, nil
}
