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

func (s *ImagesService) Blur(img image.Image, percentage int32) (*image.Image, error) {
	if percentage < 0 || percentage > 100 {
		return nil, InvalidBlurPercentageError
	}

	blurredImgNRGBA := imaging.Blur(img, float64(percentage)/2.5)
	var blurredImg image.Image = blurredImgNRGBA

	return &blurredImg, nil
}

func (s *ImagesService) Crop(img image.Image, coords image.Rectangle) (*image.Image, error) {
	croppedImgNRGBA := imaging.Crop(img, coords)
	var croppedImg image.Image = croppedImgNRGBA

	return &croppedImg, nil
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
