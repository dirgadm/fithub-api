package service

import (
	"context"
	"time"

	"github.com/dirgadm/fithub-api/internal/dto"
	"github.com/dirgadm/fithub-api/internal/model"
	"github.com/dirgadm/fithub-api/pkg/ehttp"
	"github.com/dirgadm/fithub-api/pkg/jwt"
	"github.com/dirgadm/fithub-api/pkg/log"
	"golang.org/x/crypto/bcrypt"
)

type IAuthService interface {
	Login(ctx context.Context, req dto.LoginRequest) (res dto.AuthResponse, err error)
	Register(ctx context.Context, req dto.RegisterRequest) (res dto.AuthResponse, err error)
}

type AuthService struct {
	opt SOption
}

func NewAuthService(opt SOption) IAuthService {
	return &AuthService{
		opt: opt,
	}
}

func (s *AuthService) Login(ctx context.Context, req dto.LoginRequest) (res dto.AuthResponse, err error) {
	var user model.User
	user, err = s.opt.Repository.User.GetByEmail(ctx, req.Email)
	if err != nil {
		s.opt.Common.Logger.AddMessage(log.ErrorLevel, err).Print()
		err = ehttp.ErrorOutput("email", "The email is not exists")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		s.opt.Common.Logger.AddMessage(log.ErrorLevel, err).Print()
		err = ehttp.ErrorOutput("password", "The password is not match")
		return
	}

	jwtInit := jwt.NewJWT([]byte(s.opt.Common.Config.Jwt.Key))
	expiredAt := time.Now().Add(time.Hour * 1)
	claim := jwt.UserClaim{
		UserId:    user.Id,
		ExpiresAt: expiredAt.Unix(),
	}

	token, err := jwtInit.Create(claim)
	if err != nil {
		s.opt.Common.Logger.AddMessage(log.ErrorLevel, err).Print()
		return
	}

	res.Token = token
	res.ExpiredAt = expiredAt
	res.User = dto.UserResponse{
		Id:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return
}

func (s *AuthService) Register(ctx context.Context, req dto.RegisterRequest) (res dto.AuthResponse, err error) {
	var userExist model.User
	userExist, _ = s.opt.Repository.User.GetByEmail(ctx, req.Email)
	if userExist.Id != 0 {
		s.opt.Common.Logger.AddMessage(log.ErrorLevel, err).Print()
		err = ehttp.ErrorOutput("email", "The email is already exists")
		return
	}

	var bytes []byte
	bytes, err = bcrypt.GenerateFromPassword([]byte(req.Password), 14)
	if err != nil {
		s.opt.Common.Logger.AddMessage(log.ErrorLevel, err).Print()
		return
	}
	passwordHash := string(bytes)

	user := &model.User{
		Email:     req.Email,
		Password:  passwordHash,
		Name:      req.Name,
		Phone:     req.Phone,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err = s.opt.Repository.User.Create(ctx, user)
	if err != nil {
		s.opt.Common.Logger.AddMessage(log.ErrorLevel, err).Print()
		return
	}

	jwtInit := jwt.NewJWT([]byte(s.opt.Common.Config.Jwt.Key))
	expiredAt := time.Now().Add(time.Hour * 1)
	claim := jwt.UserClaim{
		UserId:    user.Id,
		ExpiresAt: expiredAt.Unix(),
	}

	token, err := jwtInit.Create(claim)
	if err != nil {
		s.opt.Common.Logger.AddMessage(log.ErrorLevel, err).Print()
		return
	}

	res.Token = token
	res.ExpiredAt = expiredAt
	res.User = dto.UserResponse{
		Id:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return
}
