package service

import (
	"errors"
	"fmt"
	"gostore/internal/app/model"
	"gostore/internal/app/schema"
	"gostore/internal/pkg/reason"
	"time"

	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type EmailAndIDUserGetter interface {
	FindByEmail(email string) (model.User, error)
	FindByID(id int) (model.User, error)
}

type AuthRepository interface {
	Create(auth model.Auth) error
	DeleteAllByUserID(userID int) error
	Find(userID int, refreshToken string) (model.Auth, error)
}

type JWTTokenCreator interface {
	CreateAccessToken(userID int) (token string, expiredAt time.Time, err error)
	CreateRefreshToken(userID int) (token string, expiredAt time.Time, err error)
}

type SessionService struct {
	userRepo     EmailAndIDUserGetter
	authRepo     AuthRepository
	tokenCreator JWTTokenCreator
}

func NewSessionService(userRepo EmailAndIDUserGetter, authRepo AuthRepository, tokenCreator JWTTokenCreator) *SessionService {
	return &SessionService{userRepo, authRepo, tokenCreator}
}

func (s *SessionService) SignIn(req *schema.SignInReq) (schema.SignInResp, error) {
	res := schema.SignInResp{}

	// find existing user by Email
	existingUser, _ := s.userRepo.FindByEmail(req.Email)
	if existingUser.ID <= 0 {
		return res, errors.New(reason.UserNotFound)
	}

	// verify password
	isVerified := s.verifyPassword(existingUser.HashedPassword, req.Password)
	if !isVerified {
		return res, errors.New(reason.FailedSignIn)
	}

	// create JWT access token
	accessToken, _, err := s.tokenCreator.CreateAccessToken(existingUser.ID)
	if err != nil {
		log.Error("error Login - access token creation : %w", err)
		return res, errors.New(reason.FailedSignIn)
	}

	// create refresh token
	refreshToken, expiredAt, err := s.tokenCreator.CreateRefreshToken(existingUser.ID)
	if err != nil {
		log.Error("error Login - refresh token creation : %w", err)
		return res, errors.New(reason.FailedSignIn)
	}

	res.AccessToken = accessToken
	res.RefreshToken = refreshToken

	// save refresh toke to database
	authPayload := model.Auth{
		UserID:    existingUser.ID,
		Token:     refreshToken,
		AuthType:  "refresh_token",
		ExpiredAt: expiredAt,
	}
	if err := s.authRepo.Create(authPayload); err != nil {
		log.Error("error Login - refresh token saving : %w", err)
		return res, errors.New(reason.FailedSignOut)
	}

	return res, nil
}

func (s *SessionService) SignOut(userID int) error {
	if err := s.authRepo.DeleteAllByUserID(userID); err != nil {
		log.Error("error SignOut - delete refresh token : %w", err)
		return errors.New(reason.FailedSignOut)
	}

	return nil
}

func (s *SessionService) Refresh(req *schema.RefreshTokenReq) (schema.RefreshTokenResp, error) {
	res := schema.RefreshTokenResp{}

	// find existing user by ID
	existingUser, _ := s.userRepo.FindByID(req.UserID)
	if existingUser.ID <= 0 {
		return res, errors.New(reason.FailedRefreshToken)
	}

	// find existing refresh token
	auth, err := s.authRepo.Find(existingUser.ID, req.RefreshToken)
	if err != nil || auth.ID < 0 {
		log.Error(fmt.Errorf("error SessionService - refresh : %w", err))
		return res, errors.New(reason.FailedRefreshToken)
	}

	// create JWT access token
	accessToken, _, err := s.tokenCreator.CreateAccessToken(existingUser.ID)
	if err != nil {
		log.Error("error Login - access token creation : %w", err)
		return res, errors.New(reason.FailedSignIn)
	}

	res.AccessToken = accessToken
	return res, nil
}

func (s *SessionService) verifyPassword(hashPwd string, plainPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPwd), []byte(plainPwd))
	return err == nil
}
