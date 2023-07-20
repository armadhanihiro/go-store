package controller

import (
	"gostore/internal/app/schema"
	"gostore/internal/pkg/handler"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ShoppingChartService interface {
	Create(req schema.CreateShoppingChart) error
	FindAll() ([]schema.DetailShoppingChart, error)
	FindByID(id int) (schema.DetailShoppingChart, error)
	Update(id int, req schema.UpdateShoppingChart) error
	Delete(id int) error
}

type ShoppingChartController struct {
	shoppingChartService ShoppingChartService
}

func NewShoppingChartController(shoppingChartService ShoppingChartService) *ShoppingChartController {
	return &ShoppingChartController{shoppingChartService}
}

func (c *ShoppingChartController) CreateShoppingChart(ctx *gin.Context) {
	req := schema.CreateShoppingChart{}

	if handler.BindAndCheck(ctx, &req) {
		return
	}

	if err := c.shoppingChartService.Create(req); err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusCreated, "success create shopping chart", nil)
}

func (c *ShoppingChartController) BrowseShoppingChart(ctx *gin.Context) {

	resp, err := c.shoppingChartService.FindAll()
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success get shopping chart", resp)
}

func (c *ShoppingChartController) DetailShoppingChart(ctx *gin.Context) {
	id, _ := ctx.Params.Get("id")
	idInt, _ := strconv.Atoi(id)

	shopping_chart, err := c.shoppingChartService.FindByID(idInt)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success get shopping chart", shopping_chart)
}

func (c *ShoppingChartController) UpdateShoppingChart(ctx *gin.Context) {
	id, _ := ctx.Params.Get("id")
	idInt, _ := strconv.Atoi(id)
	req := schema.UpdateShoppingChart{}

	if handler.BindAndCheck(ctx, &req) {
		return
	}

	if err := c.shoppingChartService.Update(idInt, req); err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success update shopping chart", nil)
}

func (c *ShoppingChartController) DeleteShoppingChart(ctx *gin.Context) {
	id, _ := ctx.Params.Get("id")
	idInt, _ := strconv.Atoi(id)

	if err := c.shoppingChartService.Delete(idInt); err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success delete shopping chart", nil)
}
