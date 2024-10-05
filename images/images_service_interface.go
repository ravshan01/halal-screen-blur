package images

import "halal-screen-blur/proto"

type IIMagesService interface {
	CheckIsValidImage(image []byte) (bool, error)
	Blur(image []byte, percentage int32) ([]byte, error)
	Crop(image []byte, coords proto.Detection_Coords) ([]byte, error)
}
