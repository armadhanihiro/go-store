package service

import (
	"errors"
	"gostore/internal/app/model"
	"gostore/internal/app/schema"
	"gostore/internal/pkg/reason"
)

type ProductCategory interface {
	Create(productCategory model.ProductCategory) error
	FindAll() ([]model.ProductCategory, error)
	FindByID(id int) (model.ProductCategory, error)
	Update(productCategory model.ProductCategory) error
	DeleteByID(id int) error
}

type CategoryService struct {
	productCategory ProductCategory
}

func NewProductCategoryService(productCategory ProductCategory) *CategoryService {
	return &CategoryService{productCategory}
}

func (s *CategoryService) Create(req schema.CreateProductCategory) error {
	inserData := model.ProductCategory{
		Name: req.Name,
	}

	if err := s.productCategory.Create(inserData); err != nil {
		return errors.New(reason.CategoryFailedCreate)
	}

	return nil
}

func (s *CategoryService) FindAll() ([]schema.DetailProductCategory, error) {
	var resp []schema.DetailProductCategory

	products, err := s.productCategory.FindAll()
	if err != nil {
		return resp, errors.New(reason.CategoryFailedBrowse)
	}

	for _, val := range products {
		resData := schema.DetailProductCategory{
			ID:          val.ID,
			Name:        val.Name,
			CreatedAt: val.CreatedAt,
			UpdatedAt: val.UpdatedAt,
		}

		resp = append(resp, resData)
	}

	return resp, nil
}

func (s *CategoryService) FindByID(id int) (schema.DetailProductCategory, error) {
	resp := schema.DetailProductCategory{}

	product, err := s.productCategory.FindByID(id)
	if err != nil {
		return resp, errors.New(reason.CategoryFailedGetDetail)
	}

	resp.ID = product.ID
	resp.Name = product.Name
	resp.CreatedAt = product.CreatedAt
	resp.UpdatedAt = product.UpdatedAt

	return resp, nil
}

func (s *CategoryService) Update(id int, req schema.UpdateProductCategory) error {

	inserData := model.ProductCategory{
		ID : id,
		Name: req.Name,
	}

	if err := s.productCategory.Update(inserData); err != nil {
		return errors.New(reason.CategoryFailedUpdate)
	}

	return nil
}

func (s *CategoryService) Delete(id int) error {
	if err := s.productCategory.DeleteByID(id); err != nil {
		return errors.New(reason.CategoryFailedDelete)
	}

	return nil
}