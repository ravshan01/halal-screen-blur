package images

type IIMagesService interface {
	CheckIsValidImage(image []byte) (bool, error)
	// Blur percentage - 0 to 100, otherwise return InvalidBlurPercentageError
	Blur(image []byte, percentage int32) ([]byte, error)
	// Crop coords - [x1, y1, width, height]
	Crop(image []byte, coords [4]int) ([]byte, error)
}
