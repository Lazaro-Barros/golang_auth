package application

import (
	"github.com/severusTI/auth_golang/internal/application/dtos"
	"github.com/severusTI/auth_golang/internal/application/service"
	"github.com/severusTI/auth_golang/internal/interfaces/persistance/repositories"
	"github.com/severusTI/auth_golang/pkg/ops"
	"github.com/severusTI/auth_golang/pkg/password"
)

type ILoginApplication interface {
	Login(reqLogin *dtos.ReqLogin) (resLogin *dtos.ResLogin, err error)
}

type LoginApplication struct {
	userRepository repositories.IUserRepository
	tokenService   service.ITokenService
}

func NewLoginApplication(userRepository repositories.IUserRepository, tokenService service.ITokenService) ILoginApplication {
	return &LoginApplication{
		userRepository: userRepository,
		tokenService:   tokenService,
	}
}

func (obj *LoginApplication) Login(reqLogin *dtos.ReqLogin) (resLogin *dtos.ResLogin, err error) {
	user, err := obj.userRepository.GetUserByEmail(&reqLogin.Email)
	if err != nil {
		return nil, ops.Err(err)
	}

	if !password.ComparePassword(reqLogin.Password, user.HashedPassword()) {
		return nil, ops.NewErro("Invalid Password")
	}

	token, err := obj.tokenService.GenerateToken(user.Id())
	if err != nil {
		return nil, ops.Err(err)
	}

	return &dtos.ResLogin{Token: token}, nil
}
