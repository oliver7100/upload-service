package proto

import (
	"bytes"
	"io"
	"os"

	imagestore "github.com/oliver7100/upload-service/internal/image-store"
)

type UploadService struct {
	UnimplementedUploadServiceServer
	Store imagestore.IStorage
}

func (s *UploadService) UploadImage(stream UploadService_UploadImageServer) error {
	buffer := bytes.Buffer{}

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		_, err = buffer.Write(req.GetChunkData())

		if err != nil {
			return err
		}
	}

	file, err := os.CreateTemp("", "fmf-*.jpeg")

	defer os.Remove(file.Name())

	if err != nil {
		return err
	}

	_, err = file.Write(buffer.Bytes())

	if err != nil {
		return err
	}

	res, err := s.Store.SaveImage(file.Name())

	if err != nil {
		return err
	}

	return stream.SendAndClose(&UploadImageResponse{
		Uri: res.Url,
	})
}
