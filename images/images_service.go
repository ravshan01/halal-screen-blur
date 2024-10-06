package images

import (
	"bytes"
	"github.com/disintegration/imaging"
	_ "golang.org/x/image/bmp"
	_ "golang.org/x/image/tiff"
	_ "golang.org/x/image/webp"
	"image"
	"image/png"
)

type ImagesService struct {
}

func (s *ImagesService) CheckIsValidImage(imgBytes []byte) bool {
	_, err := s.BytesToImage(imgBytes)
	if err != nil {
		return false
	}

	return true
}

func (s *ImagesService) Blur(imgBytes []byte, percentage int32) (*[]byte, error) {
	img, imgErr := s.BytesToImage(imgBytes)
	if imgErr != nil {
		return nil, InvalidImageError
	}
	if percentage < 0 || percentage > 100 {
		return nil, InvalidBlurPercentageError
	}

	blurredImg := imaging.Blur(*img, float64(percentage)/2.5)
	blurredImgBytes, bytesErr := s.ImageToBytes(blurredImg)
	if bytesErr != nil {
		return nil, bytesErr
	}

	return blurredImgBytes, nil
}

func (s *ImagesService) Crop(imgBytes []byte, coords [4]int) (*[]byte, error) {
	img, imgErr := s.BytesToImage(imgBytes)
	if imgErr != nil {
		return nil, InvalidImageError
	}

	croppedImg := imaging.Crop(*img, image.Rect(coords[0], coords[1], coords[0]+coords[2], coords[1]+coords[3]))
	croppedImgBytes, bytesErr := s.ImageToBytes(croppedImg)
	if bytesErr != nil {
		return nil, bytesErr
	}

	return croppedImgBytes, nil
}

func (s *ImagesService) BytesToImage(imgBytes []byte) (*image.Image, error) {
	img, _, err := image.Decode(bytes.NewReader(imgBytes))
	if err != nil {
		return nil, err
	}

	return &img, nil
}

// ImageToBytes converts an image to bytes with png.Encode
func (s *ImagesService) ImageToBytes(img image.Image) (*[]byte, error) {
	buf := new(bytes.Buffer)
	err := png.Encode(buf, img)
	if err != nil {
		return nil, err
	}

	imgBytes := buf.Bytes()
	return &imgBytes, nil
}
