package controller

import (
	"gostore/internal/app/schema"
	"gostore/internal/pkg/handler"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type WishListService interface {
	Create(req schema.CreateWishList) error
	FindAll(req schema.BrowseWishList) ([]schema.DetailWishList, error)
	FindByID(id int) (schema.DetailWishList, error)
	Update(id int, req schema.UpdateWishList) error
	Delete(id int) error
}

type WishListController struct {
	wishListService WishListService
}

func NewWishListController(wishListService WishListService) *WishListController {
	return &WishListController{wishListService}
}

func (c *WishListController) CreateWishList(ctx *gin.Context) {
	req := schema.CreateWishList{}

	if handler.BindAndCheck(ctx, &req) {
		return
	}

	if err := c.wishListService.Create(req); err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusCreated, "success create wishlist", nil)
}

func (c *WishListController) BrowseWishList(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.Query("page"))
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	if limit == 0 {
		limit = 10
	}
	req := schema.BrowseWishList{
		Product: ctx.Query("product"),
		User: ctx.Query("user"),
		SortBy:   ctx.Query("sort_by"),
		Page:     page,
		Limit:    limit,
	}

	resp, err := c.wishListService.FindAll(req)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success get wishlist", resp)
}

func (c *WishListController) DetailWishList(ctx *gin.Context) {
	id, _ := ctx.Params.Get("id")
	idInt, _ := strconv.Atoi(id)

	wish_list, err := c.wishListService.FindByID(idInt)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success get wishlist", wish_list)
}

func (c *WishListController) UpdateWishList(ctx *gin.Context) {
	id, _ := ctx.Params.Get("id")
	idInt, _ := strconv.Atoi(id)
	req := schema.UpdateWishList{}

	if handler.BindAndCheck(ctx, &req) {
		return
	}

	if err := c.wishListService.Update(idInt, req); err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success update wishList", nil)
}

func (c *WishListController) DeleteWishList(ctx *gin.Context) {
	id, _ := ctx.Params.Get("id")
	idInt, _ := strconv.Atoi(id)

	if err := c.wishListService.Delete(idInt); err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success delete wishList", nil)
}
