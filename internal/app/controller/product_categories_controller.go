package controller

import (
	"fmt"
	"gostore/internal/app/schema"
	"gostore/internal/pkg/handler"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryService interface {
	Create(req schema.CreateProductCategory) error
	FindAll() ([]schema.DetailProductCategory, error)
	FindByID(id int) (schema.DetailProductCategory, error)
	Update(id int, req schema.UpdateProductCategory) error
	Delete(id int) error
}

type CategoryController struct {
	categoryService CategoryService
}

func NewCategoryController(categoryService CategoryService) *CategoryController {
	return &CategoryController{categoryService}
}

func (c *CategoryController) CreateCategory(ctx *gin.Context) {
	req := schema.CreateProductCategory{}

	if handler.BindAndCheck(ctx, &req) {
		return
	}

	if err := c.categoryService.Create(req); err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusCreated, "success create category", nil)
}

func (c *CategoryController) GetAllCategory(ctx *gin.Context) {
	resp, err := c.categoryService.FindAll()
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success get category", resp)
}

func (c *CategoryController) DetailCategory(ctx *gin.Context) {
	id, _ := ctx.Params.Get("id")
	idInt, _ := strconv.Atoi(id)

	product, err := c.categoryService.FindByID(idInt)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success get product", product)
}

func (c *CategoryController) UpdateCategory(ctx *gin.Context) {
	id, _ := ctx.Params.Get("id")
	idInt, _ := strconv.Atoi(id)
	req := schema.UpdateProductCategory{}

	if handler.BindAndCheck(ctx, &req) {
		return
	}

	if err := c.categoryService.Update(idInt, req); err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success update category", nil)
}

func (c *CategoryController) DeleteCategory(ctx *gin.Context) {
	id, _ := ctx.Params.Get("id")
	idInt, _ := strconv.Atoi(id)

	fmt.Println("ID", idInt)

	if err := c.categoryService.Delete(idInt); err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success delete category", nil)
}
