package proto

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

type UploadService struct {
	UnimplementedUploadServiceServer
}

func (s *UploadService) UploadImage(stream UploadService_UploadImageServer) error {
	fmt.Println("test")

	imageData := bytes.Buffer{}

	file, _ := os.Create("file.jpg")

	defer file.Close()

	for {

		fmt.Println("run")

		req, err := stream.Recv()

		if err == io.EOF {
			fmt.Println("no more data")
			break
		}

		if err != nil {
			return err
		}

		_, err = imageData.Write(req.GetChunkData())

		if err != nil {
			return err
		}
		fmt.Println("dadada")
	}

	fmt.Println("dada")

	file.Write(imageData.Bytes())

	return stream.SendAndClose(&UploadImageResponse{
		Uri:  "adadad",
		Size: 100,
	})
}
