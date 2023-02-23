package main

import (
	"net"

	"github.com/oliver7100/upload-service/proto"
	"google.golang.org/grpc"
)

func main() {
	/* cloudinaryClient, err := cloudinary.NewFromParams("zanzanzan", "748773632958652", "a5puHSHwEyy12RtXBz44fPr104s")

	if err != nil {
		panic(err)
	}

	cloudinaryClient.Upload.Upload() */

	l, err := net.Listen("tcp", ":6000")

	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()

	proto.RegisterUploadServiceServer(
		s,
		&proto.UploadService{},
	)

	panic(s.Serve(l))
}
