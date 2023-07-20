package service

import (
	"errors"
	"gostore/internal/app/model"
	"gostore/internal/app/schema"
	"gostore/internal/pkg/reason"
)

type WishListRepository interface {
	Create(wish_list model.WishList) error
	FindAll(sort_by string, page int, limit int) ([]model.WishList, error)
	FindByID(id int) (model.WishList, error)
	Update(wish_list model.WishList) error
	DeleteByID(id int) error
}

type WishListService struct {
	userFindAndUpdate UserFindAndUpdate
	productRepository  ProductRepository
	wishListRepository  WishListRepository
}

func NewWishListService(userFindAndUpdate UserFindAndUpdate, productRepository ProductRepository, wishListRepository WishListRepository) *WishListService {
	return &WishListService{userFindAndUpdate, productRepository ,wishListRepository}
}

func (s *WishListService) Create(req schema.CreateWishList) error {
	existingProduct, err := s.productRepository.FindByID(req.ProductID)
	if err != nil || existingProduct.ID <= 0 {
		return errors.New("product not found")
	}

	existingUser, err := s.userFindAndUpdate.FindByID(req.UserID)
	if err != nil || existingUser.ID <= 0 {
		return errors.New("user not found")
	}

	insertData := model.WishList{
		UserID:         existingUser.ID,
		ProductID: 		existingProduct.ID,
	}

	if err := s.wishListRepository.Create(insertData); err != nil {
		return errors.New(reason.WishListFailedCreate)
	}

	return nil
}

func (s *WishListService) FindAll(req schema.BrowseWishList) ([]schema.DetailWishList, error) {
	var resp []schema.DetailWishList

	wish_list, err := s.wishListRepository.FindAll(req.SortBy, req.Page, req.Limit)
	if err != nil {
		return resp, errors.New(reason.WishListFailedBrowse)
	}

	for _, val := range wish_list {
		resData := schema.DetailWishList{
			ID:          val.ID,
			User: schema.DetailUserResp{
				ID:   val.User.ID,
				FullName: val.User.FullName,
				Email: val.User.Email,
				Address: val.User.Address,
				City: val.User.City,
				Province: val.User.Province,
				Country: val.User.Country,
				PostalCode: val.User.PostalCode,
			},
			Product: schema.DetailProductResp{
				ID: val.Product.ID,
				Name: val.Product.Name,
				Description: val.Product.Description,
				Sold: val.Product.Sold,
				Amount: val.Product.Amount,
				Views: val.Product.Views,
				Image: val.Product.Image,
			},
			CreatedAt: val.CreatedAt,
			UpdatedAt: val.UpdatedAt,
		}

		resp = append(resp, resData)
	}

	return resp, nil
}

func (s *WishListService) FindByID(id int) (schema.DetailWishList, error) {
	resp := schema.DetailWishList{}

	wish_list, err := s.wishListRepository.FindByID(id)
	if err != nil || wish_list.ID == 0 {
		return resp, errors.New(reason.WishListNotFound)
	}

	views, err := s.viewIncrement(wish_list.ID, wish_list.Views)
	if err != nil {
		return resp, errors.New("failed to update wishList views")
	}

	resp.ID = wish_list.ID
	resp.Views = views
	resp.User = schema.DetailUserResp{
		ID: wish_list.User.ID, 
		FullName: wish_list.User.FullName,
		Email: wish_list.User.Email,
		Address: wish_list.User.Address,
		City: wish_list.User.City,
		Province: wish_list.User.Province,
		Country: wish_list.User.Country,
		PostalCode: wish_list.User.PostalCode,
	}
	resp.Product = schema.DetailProductResp{
		ID: wish_list.Product.ID,
		Name: wish_list.Product.Name,
		Description: wish_list.Product.Description,
		Sold: wish_list.Product.Sold,
		Amount: wish_list.Product.Amount,
		Views: wish_list.Product.Views,
		Image: wish_list.Product.Image,
	}
	resp.CreatedAt = wish_list.CreatedAt
	resp.UpdatedAt = wish_list.UpdatedAt

	return resp, nil
}

func (s *WishListService) Update(id int, req schema.UpdateWishList) error {
	existingProduct, err := s.productRepository.FindByID(req.ProductID)
	if err != nil || existingProduct.ID <= 0 {
		return errors.New("product not found")
	}

	existingUser, err := s.userFindAndUpdate.FindByID(req.UserID)
	if err != nil || existingUser.ID <= 0 {
		return errors.New("user not found")
	}

	insertData := model.WishList{
		UserID:         existingUser.ID,
		ProductID: 		existingProduct.ID,
	}

	if err := s.wishListRepository.Update(insertData); err != nil {
		return errors.New(reason.WishListFailedUpdate)
	}

	return nil
}

func (s *WishListService) Delete(id int) error {
	if err := s.wishListRepository.DeleteByID(id); err != nil {
		return errors.New(reason.WishListFailedDelete)
	}

	return nil
}

func (s *WishListService) viewIncrement(id int, views int) (int, error) {
	insertData := model.WishList{
		ID:    id,
		Views: views + 1,
	}

	if err := s.wishListRepository.Update(insertData); err != nil {
		return views, errors.New(reason.WishListFailedUpdate)
	}

	return views + 1, nil
}
