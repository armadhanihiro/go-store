package service

import (
	"errors"
	"gostore/internal/app/model"
	"gostore/internal/app/schema"
	"gostore/internal/pkg/hashed"
	"gostore/internal/pkg/reason"
)

type UserCreator interface {
	Create(user model.User) error
}

type UserGetter interface {
	FindByEmail(email string) (model.User, error)
}

type UserGetterAndCreator interface {
	UserGetter
	UserCreator
}

type SignUpService struct {
	userGetterAndCreator UserGetterAndCreator
}

func NewSignUpService(userGetterAndCreator UserGetterAndCreator) *SignUpService {
	return &SignUpService{userGetterAndCreator}
}

func (s *SignUpService) Insert(req *schema.SignUpReq) error {
	existingUser, _ := s.userGetterAndCreator.FindByEmail(req.Email)

	if existingUser.ID > 0 {
		return errors.New(reason.UserAlreadyExist)
	}

	hashedPassword, _ := hashed.HashedPassword(req.Password)
	insertData := model.User{
		FullName:       req.FullName,
		Email:          req.Email,
		HashedPassword: hashedPassword,
	}

	if err := s.userGetterAndCreator.Create(insertData); err != nil {
		return errors.New(reason.SignUpFailed)
	}

	return nil
}
