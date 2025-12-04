package services

type AuthService interface {
	Login() (int, error)
}

type authService struct{}

func NewAuthService() AuthService {
	return &authService{}
}

func (as *authService) Login() (int, error) {
	return 1, nil
}
