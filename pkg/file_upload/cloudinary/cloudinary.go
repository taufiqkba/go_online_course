package cloudinary

import (
	"context"
	"errors"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/google/uuid"
	"go_online_course/pkg/utils"
	"mime/multipart"
	"os"
)

type FileUpload interface {
	Upload(file multipart.FileHeader) (*string, error)
	Delete(fileName string) (*string, error)
}

type FileUploadImpl struct {
}

func (fileUploadImpl *FileUploadImpl) Upload(file multipart.FileHeader) (*string, error) {
	//connect to cloudinary
	cld, err := cloudinary.NewFromURL("cloudinary://" + os.Getenv("CLOUDINARY_APIKEY") + ":" + os.Getenv("CLOUDINARY_SECRET") + "@" + os.Getenv("CLOUDINARY_CLOUD_NAME"))

	if err != nil {
		return nil, err
	}

	//	using context to avoid deadlock
	var ctx = context.Background()
	binary, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer binary.Close()

	if binary != nil {
		uploadResult, err := cld.Upload.Upload(
			ctx,
			binary,
			uploader.UploadParams{PublicID: uuid.New().String()},
		)
		if err != nil {
			return nil, err
		}
		return &uploadResult.SecureURL, nil
	}
	return nil, errors.New("file upload format does not match")
}

func (fileUploadImpl *FileUploadImpl) Delete(fileName string) (*string, error) {
	//connect to cloudinary
	cld, err := cloudinary.NewFromURL("cloudinary://" + os.Getenv("CLOUDINARY_APIKEY") + ":" + os.Getenv("CLOUDINARY_SECRET") + "@" + os.Getenv("CLOUDINARY_CLOUD_NAME"))

	if err != nil {
		return nil, err
	}
	var ctx = context.Background()
	fileName = utils.GetFileName(fileName)
	resp, err := cld.Upload.Destroy(ctx, uploader.DestroyParams{PublicID: fileName})
	if err != nil {
		return nil, err
	}
	return &resp.Result, nil
}

func NewFileUpload() FileUpload {
	return &FileUploadImpl{}
}
