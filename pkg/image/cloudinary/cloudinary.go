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
	Delete(file string) (*string, error)
}

type ImageImpl struct {
}

func (fileUploadImpl *ImageImpl) Delete(fileName string) (*string, error) {
	cld, err := cloudinary.NewFromURL("cloudinary://" + os.Getenv("CLOUDINARY_APIKEY") + ":" + os.Getenv("CLOUDINARY_SECRET") + "@" + os.Getenv("CLOUDINARY_NAME"))
	if err != nil {
		return nil, err
	}

	fileName = utils.GetFileName(fileName)
	resp, err := cld.Upload.Destroy(ctx, uploader.DestroyParams{PublicID: fileName})
	if err != nil {
		return nil, err
	}
	return &resp.Result, nil
}

func (fileUploadImpl *ImageImpl) Upload(fileName multipart.FileHeader) (*string, error) {
	//Upload Image
	cld, err := cloudinary.NewFromURL("cloudinary://" + os.Getenv("CLOUDINARY_APIKEY") + ":" + os.Getenv("CLOUDINARY_SECRET") + "@" + os.Getenv("CLOUDINARY_NAME"))

	if err != nil {
		return nil, err
	}

	var ctx = context.Background()
	binary, err := fileName.Open()
	defer binary.Close()

	if err != nil {
		return nil, err
	}

	if binary != nil {
		uploadResult, err := cld.Upload.Upload(
			ctx,
			binary,
			uploader.UploadParams{
				PublicID: uuid.New().String(),
			},
		)
		if err != nil {
			return nil, err
		}
		return &uploadResult.SecureURL, nil
	}

	return nil, errors.New("error")

}

func NewImage() *ImageImpl {
	return &ImageImpl{}
}
