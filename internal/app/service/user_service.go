package service

import (
	"errors"
	"gostore/internal/app/model"
	"gostore/internal/app/schema"
	"gostore/internal/pkg/hashed"
	"gostore/internal/pkg/reason"
)

type UserIDFinder interface {
	FindByID(id int) (model.User, error)
}

type UserUpdater interface {
	Update(user model.User) error
}

type UserFindAndUpdate interface {
	UserIDFinder
	UserUpdater
}

type UserService struct {
	userFindAndUpdate UserFindAndUpdate
}

func NewUserService(userFindAndUpdate UserFindAndUpdate) *UserService {
	return &UserService{userFindAndUpdate}
}

func (s *UserService) DetailUser(id int) (schema.DetailUserResp, error) {
	resp := schema.DetailUserResp{}

	existingUser, err := s.userFindAndUpdate.FindByID(id)
	if err != nil {
		return resp, errors.New(reason.UserNotFound)
	}

	resp.ID = existingUser.ID
	resp.FullName = existingUser.FullName
	resp.Email = existingUser.Email
	resp.Address = existingUser.Address
	resp.City = existingUser.City
	resp.Province = existingUser.Province
	resp.Country = existingUser.Country
	resp.PostalCode = existingUser.PostalCode
	resp.CreatedAt = existingUser.CreatedAt
	resp.UpdatedAt = existingUser.UpdatedAt
	resp.DeletedAt = existingUser.DeletedAt

	return resp, nil
}

func (s *UserService) Update(id int, req schema.UpdateUserReq) error {
	hashedPassword, _ := hashed.HashedPassword(req.Password)

	user := model.User{
		ID:             id,
		FullName:       req.FullName,
		Email:          req.Email,
		HashedPassword: hashedPassword,
		Address:        req.Address,
		City:           req.City,
		Province:       req.Province,
		Country:        req.Country,
		PostalCode:     req.PostalCode,
	}
	if err := s.userFindAndUpdate.Update(user); err != nil {
		return errors.New(reason.FailedUpdateUser)
	}

	return nil
}
