package service

import (
	"errors"
	"gostore/internal/app/model"
	"gostore/internal/app/schema"
	"gostore/internal/pkg/reason"
	"mime/multipart"
)

type ImageUploader interface {
	UploadImage(productID int, input *multipart.FileHeader) (imageURL string, err error)
}

type CategoryRepository interface {
	FindByID(id int) (model.ProductCategory, error)
	FindByName(name string) (model.ProductCategory, error)
}

type ProductRepository interface {
	Create(product model.Product) error
	FindAll(name string, sort_by string, page int, limit int) ([]model.Product, error)
	FindByID(id int) (model.Product, error)
	Update(product model.Product) error
	UpdateImageURL(id int, imageURL string) error
	DeleteByID(id int) error
}

type ProductService struct {
	imageUploader      ImageUploader
	categoryRepository CategoryRepository
	productRepository  ProductRepository
}

func NewProductService(imageUploader ImageUploader, categoryRepository CategoryRepository, productRepository ProductRepository) *ProductService {
	return &ProductService{imageUploader, categoryRepository, productRepository}
}

func (s *ProductService) Create(req schema.CreateProductReq) error {
	existingCategory, err := s.categoryRepository.FindByID(req.ProductCategoryID)
	if err != nil || existingCategory.ID <= 0 {
		return errors.New("category not found")
	}

	imageUrl, err := s.imageUploader.UploadImage(1, req.Image)
	if err != nil {
		return errors.New(reason.ProductFailedCreate)
	}

	inserData := model.Product{
		Description:       req.Description,
		Name:              req.Name,
		ProductCategoryID: existingCategory.ID,
		Amount:            req.Amount,
		Image:             imageUrl,
	}

	if err := s.productRepository.Create(inserData); err != nil {
		return errors.New(reason.ProductFailedCreate)
	}

	return nil
}

func (s *ProductService) FindAll(req schema.BrowsProductReq) ([]schema.DetailProductResp, error) {
	var resp []schema.DetailProductResp

	products, err := s.productRepository.FindAll(req.Name, req.SortBy, req.Page, req.Limit)
	if err != nil {
		return resp, errors.New(reason.ProductFailedBrowse)
	}

	for _, val := range products {
		resData := schema.DetailProductResp{
			ID:          val.ID,
			Name:        val.Name,
			Description: val.Description,
			Sold:        val.Sold,
			Amount:      val.Amount,
			Views:       val.Views,
			Image:       val.Image,
			Category: schema.DetailProductCategory{
				ID:   val.ProductCategory.ID,
				Name: val.ProductCategory.Name,
			},
			CreatedAt: val.CreatedAt,
			UpdatedAt: val.UpdatedAt,
		}

		resp = append(resp, resData)
	}

	return resp, nil
}

func (s *ProductService) FindByID(id int) (schema.DetailProductResp, error) {
	resp := schema.DetailProductResp{}

	product, err := s.productRepository.FindByID(id)
	if err != nil || product.ID == 0 {
		return resp, errors.New(reason.ProductNotFound)
	}

	views, err := s.viewIncrement(product.ID, product.Views)
	if err != nil {
		return resp, errors.New("failed to update product views")
	}

	resp.ID = product.ID
	resp.Name = product.Name
	resp.Description = product.Description
	resp.Sold = product.Sold
	resp.Amount = product.Amount
	resp.Views = views
	resp.Image = product.Image
	resp.Category = schema.DetailProductCategory{ID: product.ProductCategory.ID, Name: product.ProductCategory.Name}
	resp.CreatedAt = product.CreatedAt
	resp.UpdatedAt = product.UpdatedAt

	return resp, nil
}

func (s *ProductService) Update(id int, req schema.UpdateProductReq) error {
	existingCategory, err := s.categoryRepository.FindByID(req.ProductCategoryID)
	if err != nil || existingCategory.ID <= 0 {
		return errors.New("category not found")
	}

	imageUrl, err := s.imageUploader.UploadImage(id, req.Image)
	if err != nil {
		return errors.New(reason.ProductFailedUpdate)
	}

	insertData := model.Product{
		ID:                id,
		Description:       req.Description,
		Name:              req.Name,
		ProductCategoryID: existingCategory.ID,
		Amount:            req.Amount,
		Image:             imageUrl,
	}

	if err := s.productRepository.Update(insertData); err != nil {
		return errors.New(reason.ProductFailedUpdate)
	}

	return nil
}

func (s *ProductService) Delete(id int) error {
	if err := s.productRepository.DeleteByID(id); err != nil {
		return errors.New(reason.ProductFailedDelete)
	}

	return nil
}

func (s *ProductService) viewIncrement(id int, views int) (int, error) {
	insertData := model.Product{
		ID:    id,
		Views: views + 1,
	}

	if err := s.productRepository.Update(insertData); err != nil {
		return views, errors.New(reason.ProductFailedUpdate)
	}

	return views + 1, nil
}
