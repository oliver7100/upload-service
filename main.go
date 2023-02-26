package main

import (
	"net"
	"os"

	"github.com/joho/godotenv"
	imagestore "github.com/oliver7100/upload-service/internal/image-store"
	"github.com/oliver7100/upload-service/proto"
	"google.golang.org/grpc"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}
}

func main() {
	storage, err := imagestore.NewCloudinaryStorage(
		&imagestore.CloudinaryStorageConfig{
			Cloud:  os.Getenv("CLOUDINARY_CLOUD_NAME"),
			Key:    os.Getenv("CLOUDINARY_CLOUD_KEY"),
			Secret: os.Getenv("CLOUDINARY_CLOUD_SECRET"),
		},
	)

	l, err := net.Listen(
		"tcp",
		os.Getenv("PORT"),
	)

	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()

	proto.RegisterUploadServiceServer(
		s,
		&proto.UploadService{
			Store: storage,
		},
	)

	panic(s.Serve(l))
}
