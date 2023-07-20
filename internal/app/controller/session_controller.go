package controller

import (
	"gostore/internal/app/schema"
	"gostore/internal/pkg/handler"
	"gostore/internal/pkg/reason"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SessionService interface {
	SignIn(req *schema.SignInReq) (schema.SignInResp, error)
	Refresh(req *schema.RefreshTokenReq) (schema.RefreshTokenResp, error)
	SignOut(userID int) error
}

type RefreshTokenVerifier interface {
	VerifyRefreshToken(tokenString string) (string, error)
}

type SessionController struct {
	sessionService SessionService
	tokenCreator   RefreshTokenVerifier
}

func NewSessionController(sessionService SessionService, tokenCreator RefreshTokenVerifier) *SessionController {
	return &SessionController{sessionService, tokenCreator}
}

func (sc *SessionController) SignIn(c *gin.Context) {
	req := schema.SignInReq{}

	if handler.BindAndCheck(c, &req) {
		return
	}

	resp, err := sc.sessionService.SignIn(&req)
	if err != nil {
		handler.ResponseError(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(c, http.StatusOK, "success sign in", resp)
}

func (c *SessionController) Refresh(ctx *gin.Context) {
	refreshToken := ctx.GetHeader("refresh_token")
	if refreshToken == "" {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, reason.FailedRefreshToken)
	}

	sub, err := c.tokenCreator.VerifyRefreshToken(refreshToken)

	if err != nil {
		handler.ResponseError(ctx, http.StatusUnauthorized, reason.FailedRefreshToken)
		return
	}

	intSub, _ := strconv.Atoi(sub)
	req := &schema.RefreshTokenReq{}
	req.RefreshToken = refreshToken
	req.UserID = intSub

	res, err := c.sessionService.Refresh(req)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, reason.FailedRefreshToken)
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success refresh token", res)
}

func (c *SessionController) SignOut(ctx *gin.Context) {
	userID, _ := strconv.Atoi(ctx.GetString("user_id"))
	if err := c.sessionService.SignOut(userID); err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success sign out", nil)
}
