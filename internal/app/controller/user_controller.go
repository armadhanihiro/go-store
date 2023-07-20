package controller

import (
	"gostore/internal/app/schema"
	"gostore/internal/pkg/handler"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserIDFinder interface {
	DetailUser(id int) (schema.DetailUserResp, error)
}

type UserUpdater interface {
	Update(id int, req schema.UpdateUserReq) error
}

type UserFindAndUpdate interface {
	UserIDFinder
	UserUpdater
}

type UserController struct {
	userFindAndUpdate UserFindAndUpdate
}

func NewUserController(userFindAndUpdate UserFindAndUpdate) *UserController {
	return &UserController{userFindAndUpdate}
}

func (uc *UserController) FindByID(c *gin.Context) {
	userID, _ := strconv.Atoi(c.GetString("user_id"))

	resp, err := uc.userFindAndUpdate.DetailUser(userID)
	if err != nil {
		handler.ResponseError(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(c, http.StatusOK, "success get user detail", resp)

}

func (uc *UserController) UpdateByID(c *gin.Context) {
	req := schema.UpdateUserReq{}
	userID, _ := strconv.Atoi(c.GetString("user_id"))

	if handler.BindAndCheck(c, &req) {
		return
	}

	if err := uc.userFindAndUpdate.Update(userID, req); err != nil {
		handler.ResponseError(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(c, http.StatusOK, "success update user data", nil)
}
