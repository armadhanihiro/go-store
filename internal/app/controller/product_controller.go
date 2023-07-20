package controller

import (
	"gostore/internal/app/schema"
	"gostore/internal/pkg/handler"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductService interface {
	Create(req schema.CreateProductReq) error
	FindAll(req schema.BrowsProductReq) ([]schema.DetailProductResp, error)
	FindByID(id int) (schema.DetailProductResp, error)
	Update(id int, req schema.UpdateProductReq) error
	Delete(id int) error
}

type ProductController struct {
	productService ProductService
}

func NewProductController(productService ProductService) *ProductController {
	return &ProductController{productService}
}

func (c *ProductController) CreateCategory(ctx *gin.Context) {
	req := schema.CreateProductReq{}

	if handler.BindAndCheck(ctx, &req) {
		return
	}

	if err := c.productService.Create(req); err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusCreated, "success create product", nil)
}

func (c *ProductController) BrowseProduct(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.Query("page"))
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	if limit == 0 {
		limit = 10
	}
	req := schema.BrowsProductReq{
		Name:     ctx.Query("name"),
		Category: ctx.Query("category"),
		SortBy:   ctx.Query("sort_by"),
		Page:     page,
		Limit:    limit,
	}

	resp, err := c.productService.FindAll(req)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success get products", resp)
}

func (c *ProductController) DetailProduct(ctx *gin.Context) {
	id, _ := ctx.Params.Get("id")
	idInt, _ := strconv.Atoi(id)

	product, err := c.productService.FindByID(idInt)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success get product", product)
}

func (c *ProductController) UpdateProduct(ctx *gin.Context) {
	id, _ := ctx.Params.Get("id")
	idInt, _ := strconv.Atoi(id)
	req := schema.UpdateProductReq{}

	if handler.BindAndCheck(ctx, &req) {
		return
	}

	if err := c.productService.Update(idInt, req); err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success update product", nil)
}

func (c *ProductController) DeleteProduct(ctx *gin.Context) {
	id, _ := ctx.Params.Get("id")
	idInt, _ := strconv.Atoi(id)

	if err := c.productService.Delete(idInt); err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success delete product", nil)
}
