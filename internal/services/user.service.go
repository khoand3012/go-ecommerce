package services

import "github.com/khoand3012/go-ecommerce/internal/repositories"

type UserService struct {
	userRepo *repositories.UserRepo
}

func NewUserService() *UserService {
	return &UserService{
		userRepo: repositories.NewUserRepo(),
	}
}

func (us *UserService) GetUserInfoService() string {
	return us.userRepo.GetUserInfo()
}
