package service

import (
	"errors"
	"gostore/internal/app/model"
	"gostore/internal/app/schema"
	"gostore/internal/pkg/reason"
)

type ShoppingChartRepository interface {
	Create(shopping_chart model.ShoppingChart) error
	FindAll() ([]model.ShoppingChart, error)
	FindByID(id int) (model.ShoppingChart, error)
	Update(shopping_chart model.ShoppingChart) error
	DeleteByID(id int) error
}

type ShoppingChartService struct {
	userFindAndUpdate UserFindAndUpdate
	productRepository  ProductRepository
	shoppingChartRepository  ShoppingChartRepository
}

func NewShoppingChartService(userFindAndUpdate UserFindAndUpdate, productRepository ProductRepository, shoppingChartRepository  ShoppingChartRepository) *ShoppingChartService {
	return &ShoppingChartService{userFindAndUpdate, productRepository ,shoppingChartRepository}
}

func (s *ShoppingChartService) Create(req schema.CreateShoppingChart) error {
	existingProduct, err := s.productRepository.FindByID(req.ProductID)
	if err != nil || existingProduct.ID <= 0 {
		return errors.New("product not found")
	}

	existingUser, err := s.userFindAndUpdate.FindByID(req.UserID)
	if err != nil || existingUser.ID <= 0 {
		return errors.New("user not found")
	}

	insertData := model.ShoppingChart{
		UserID:         existingUser.ID,
		ProductID: 		existingProduct.ID,
		Quantity: 		req.Quantity,
	}

	if err := s.shoppingChartRepository.Create(insertData); err != nil {
		return errors.New(reason.ShoppingChartFailedCreate)
	}

	return nil
}

func (s *ShoppingChartService) FindAll() ([]schema.DetailShoppingChart, error) {
	var resp []schema.DetailShoppingChart

	shopping_chart, err := s.shoppingChartRepository.FindAll()
	if err != nil {
		return resp, errors.New(reason.ShoppingChartFailedBrowse)
	}

	for _, val := range shopping_chart {
		resData := schema.DetailShoppingChart{
			ID:			val.ID,
			Quantity: 	val.Quantity,
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

func (s *ShoppingChartService) FindByID(id int) (schema.DetailShoppingChart, error) {
	resp := schema.DetailShoppingChart{}

	shopping_chart, err := s.shoppingChartRepository.FindByID(id)
	if err != nil || shopping_chart.ID == 0 {
		return resp, errors.New(reason.ShoppingChartNotFound)
	}

	resp.ID = shopping_chart.ID
	resp.Quantity = shopping_chart.Quantity
	resp.User = schema.DetailUserResp{
		ID: shopping_chart.User.ID, 
		FullName: shopping_chart.User.FullName,
		Email: shopping_chart.User.Email,
		Address: shopping_chart.User.Address,
		City: shopping_chart.User.City,
		Province: shopping_chart.User.Province,
		Country: shopping_chart.User.Country,
		PostalCode: shopping_chart.User.PostalCode,
	}
	resp.Product = schema.DetailProductResp{
		ID: shopping_chart.Product.ID,
		Name: shopping_chart.Product.Name,
		Description: shopping_chart.Product.Description,
		Sold: shopping_chart.Product.Sold,
		Amount: shopping_chart.Product.Amount,
		Views: shopping_chart.Product.Views,
		Image: shopping_chart.Product.Image,
	}
	resp.CreatedAt = shopping_chart.CreatedAt
	resp.UpdatedAt = shopping_chart.UpdatedAt

	return resp, nil
}

func (s *ShoppingChartService) Update(id int, req schema.UpdateShoppingChart) error {
	existingProduct, err := s.productRepository.FindByID(req.ProductID)
	if err != nil || existingProduct.ID <= 0 {
		return errors.New("product not found")
	}

	existingUser, err := s.userFindAndUpdate.FindByID(req.UserID)
	if err != nil || existingUser.ID <= 0 {
		return errors.New("user not found")
	}

	insertData := model.ShoppingChart{
		UserID:         existingUser.ID,
		ProductID: 		existingProduct.ID,
		Quantity: 		req.Quantity,
	}

	if err := s.shoppingChartRepository.Update(insertData); err != nil {
		return errors.New(reason.ShoppingChartFailedUpdate)
	}

	return nil
}

func (s *ShoppingChartService) Delete(id int) error {
	if err := s.shoppingChartRepository.DeleteByID(id); err != nil {
		return errors.New(reason.ShoppingChartFailedDelete)
	}

	return nil
}