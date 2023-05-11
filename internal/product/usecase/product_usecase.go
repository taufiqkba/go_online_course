package usecase

import (
	"go_online_course/internal/product/dto"
	entity3 "go_online_course/internal/product/entity"
	"go_online_course/internal/product/repository"
	"go_online_course/pkg/file_upload/cloudinary"
)

type ProductUseCase interface {
	FindAll(offset int, limit int) []entity3.Product
	FindById(id int) (*entity3.Product, error)
	Create(dto dto.ProductRequestBody) (*entity3.Product, error)
	Update(id int, dto dto.ProductRequestBody) (*entity3.Product, error)
	Delete(id int) error
}

type ProductUseCaseImpl struct {
	repository repository.ProductRepository
	fileUpload cloudinary.FileUpload
}

func (usecase *ProductUseCaseImpl) FindAll(offset int, limit int) []entity3.Product {
	return usecase.repository.FindAll(offset, limit)
}

func (usecase *ProductUseCaseImpl) FindById(id int) (*entity3.Product, error) {
	return usecase.repository.FindById(id)
}

func (usecase *ProductUseCaseImpl) Create(dto dto.ProductRequestBody) (*entity3.Product, error) {
	dataProduct := entity3.Product{
		ProductCategoryID: dto.ProductCategoryID,
		Title:             dto.Title,
		Description:       dto.Description,
		Price:             dto.Price,
		CreatedByID:       dto.CreatedBy,
	}

	//	upload images
	if dto.Image != nil {
		image, err := usecase.fileUpload.Upload(*dto.Image)
		if err != nil {
			return nil, err
		}
		if image != nil {
			dataProduct.Image = image
		}
	}

	//	upload video
	if dto.Video != nil {
		video, err := usecase.fileUpload.Upload(*dto.Video)
		if err != nil {
			return nil, err
		}
		if video != nil {
			dataProduct.Video = video
		}
	}

	//	call repository to save data into database
	product, err := usecase.repository.Create(dataProduct)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (usecase *ProductUseCaseImpl) Update(id int, dto dto.ProductRequestBody) (*entity3.Product, error) {
	//	find data product based from id
	product, err := usecase.repository.FindById(id)
	if err != nil {
		return nil, err
	}
	product.Title = dto.Title
	product.Description = dto.Description
	product.Price = dto.Price
	product.UpdatedByID = &dto.UpdatedBy

	//	if there is an update file image
	if dto.Image != nil {
		image, err := usecase.fileUpload.Upload(*dto.Image)
		if err != nil {
			return nil, err
		}

		if product.Image != nil {
			//	delete image
			_, err := usecase.fileUpload.Delete(*product.Image)
			if err != nil {
				return nil, err
			}
		}
		product.Image = image
	}

	//if there is an update file video
	if dto.Video != nil {
		video, err := usecase.fileUpload.Upload(*dto.Video)
		if err != nil {
			return nil, err
		}

		if product.Video != nil {
			//	delete video
			_, err := usecase.fileUpload.Delete(*product.Video)
			if err != nil {
				return nil, err
			}
		}
		product.Video = video
	}

	updateProduct, err := usecase.repository.Update(*product)
	if err != nil {
		return nil, err
	}
	return updateProduct, nil
}

func (usecase *ProductUseCaseImpl) Delete(id int) error {
	// find data based from id
	product, err := usecase.repository.FindById(id)
	if err != nil {
		return err
	}

	err = usecase.repository.Delete(*product)
	if err != nil {
		return err
	}
	return nil
}

func NewProductUseCase(repository repository.ProductRepository, fileUpload cloudinary.FileUpload) ProductUseCase {
	return &ProductUseCaseImpl{repository, fileUpload}
}
