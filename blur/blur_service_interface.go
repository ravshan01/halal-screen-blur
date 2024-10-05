package blur

import (
	"halal-screen-blur/proto"
)

type IBlurService interface {
	BlurImages(req *proto.BlurImagesRequest) (*proto.BlurImageResponse, error)
}
