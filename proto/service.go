package proto

import (
	"bytes"
	"io"
)

type UploadService struct {
	UnimplementedUploadServiceServer
}

func (s *UploadService) UploadImage(stream UploadService_UploadImageServer) error {
	imageData := bytes.Buffer{}

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		_, err = imageData.Write(req.GetChunkData())

		if err != nil {
			return err
		}
	}

	return stream.SendAndClose(&UploadImageResponse{
		Uri:  "adadad",
		Size: 100,
	})
}
