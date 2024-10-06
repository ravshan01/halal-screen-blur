package images

import "image"

type IIMagesService interface {
	CheckIsValidImage(image []byte) bool
	// Blur percentage - 0 to 100, otherwise return InvalidBlurPercentageError
	Blur(image image.Image, percentage int32) (*image.Image, error)
	Crop(image image.Image, coords image.Rectangle) (*image.Image, error)

	BytesToImage(image []byte) (*image.Image, error)
	ImageToBytes(image image.Image) (*[]byte, error)
}
