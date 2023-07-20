package controller

import (
	"gostore/internal/app/schema"
	"gostore/internal/pkg/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Inserter interface {
	Insert(req *schema.SignUpReq) error
}

type SignUpController struct {
	inserter Inserter
}

func NewSignUpController(inserter Inserter) *SignUpController {
	return &SignUpController{inserter}
}

func (sc *SignUpController) Insert(c *gin.Context) {
	req := schema.SignUpReq{}

	if handler.BindAndCheck(c, &req) {
		return
	}

	if err := sc.inserter.Insert(&req); err != nil {
		handler.ResponseError(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(c, http.StatusOK, "success sign up", nil)
}
