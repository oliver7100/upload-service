package main

import (
	"net"

	imagestore "github.com/oliver7100/upload-service/internal/image-store"
	"github.com/oliver7100/upload-service/proto"
	"google.golang.org/grpc"
)

func main() {
	storage, err := imagestore.NewCloudinaryStorage(
		&imagestore.CloudinaryStorageConfig{
			Cloud:  "zanzanzan",
			Key:    "748773632958652",
			Secret: "a5puHSHwEyy12RtXBz44fPr104s",
		},
	)

	l, err := net.Listen("tcp", ":6000")

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
