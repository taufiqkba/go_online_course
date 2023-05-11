package cloudinary

import "mime/multipart"

type FileUpload interface {
	Upload(image multipart.FileHeader) (*string, error)
	Delete(image multipart.FileHeader) (*string, error)
}

type ImageImpl struct {
}

func (fileUpload ImageImpl) Upload(image multipart.FileHeader) (*string, error) {
	//TODO implement me
	panic("implement me")
}

func (fileUpload ImageImpl) Delete(image multipart.FileHeader) (*string, error) {
	//TODO implement me
	panic("implement me")
}

func NewImage() FileUpload {
	return &ImageImpl{}
}
