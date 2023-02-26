package imagestore

import (
	"context"
	"fmt"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/google/uuid"
)

type IStorage interface {
	SaveImage(item string) (*SaveImageResult, error)
}

type SaveImageResult struct {
	Url string
}

type CloudinaryStorageConfig struct {
	Cloud  string
	Key    string
	Secret string
}

type CloudinaryStorage struct {
	client *cloudinary.Cloudinary
}

type Base64EncodedImage string

func (s *CloudinaryStorage) SaveImage(item string) (*SaveImageResult, error) {
	fmt.Println(item)

	response, err := s.client.Upload.Upload(context.Background(), item, uploader.UploadParams{
		PublicID: uuid.NewString(),
	})

	fmt.Println(response.Error.Message)

	fmt.Println(response.AssetID)

	if err != nil {
		return nil, err
	}

	return &SaveImageResult{
		Url: response.SecureURL,
	}, nil
}

func NewCloudinaryStorage(cfg *CloudinaryStorageConfig) (IStorage, error) {

	cloudinaryClient, err := cloudinary.NewFromParams(
		cfg.Cloud,
		cfg.Key,
		cfg.Secret,
	)

	if err != nil {
		return nil, err
	}

	return &CloudinaryStorage{
		client: cloudinaryClient,
	}, nil
}
