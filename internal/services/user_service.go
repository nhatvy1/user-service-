package services

type UserService interface {
	FindUserByID(id int) (int, error)
}

type userService struct {
}

func NewUserService() UserService {
	return &userService{}
}

func (us *userService) FindUserByID(id int) (int, error) {
	return id, nil
}
