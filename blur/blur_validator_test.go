package blur

import (
	"fmt"
	"halal-screen-blur/config"
	"halal-screen-blur/images"
	"halal-screen-blur/proto"
	"image"
	"os"
	"testing"
)

var blurValidator *BlurValidator

func TestBlurValidatorIsDefined(t *testing.T) {
	setupBlurValidatorTest()
	defer teardownBlurValidatorTest()

	if blurValidator == nil {
		t.Fatalf("expected blurValidator to be defined")
	}
}

func TestBlurValidator_ValidateImagesCount(t *testing.T) {
	setupBlurServiceServerTest()
	defer teardownBlurServiceServerTest()

	t.Run("should return BadRequest error if no images provided", func(t *testing.T) {
		res := blurValidator.ValidateImagesCount([]*proto.ImageForBlur{})
		assertErrorResponse(t, res, proto.BlurErrorCode_BadRequest)
	})

	t.Run("should return MaxImagesExceeded error if too many images provided", func(t *testing.T) {
		res := blurValidator.ValidateImagesCount(make([]*proto.ImageForBlur, config.GetConfig().MaxImagesPerRequest+1))
		assertErrorResponse(t, res, proto.BlurErrorCode_MaxImagesExceeded)
	})

	t.Run("should return success if images count is valid", func(t *testing.T) {
		res := blurValidator.ValidateImagesCount(make([]*proto.ImageForBlur, 1))
		assertSuccess(t, res)
	})
}

func TestBlurValidator_ValidateImagesExist(t *testing.T) {
	setupBlurServiceServerTest()
	defer teardownBlurServiceServerTest()

	t.Run("should return BadRequest error if no images provided", func(t *testing.T) {
		res := blurValidator.ValidateImagesExist([]*proto.ImageForBlur{})
		assertErrorResponse(t, res, proto.BlurErrorCode_BadRequest)
	})

	t.Run("should return success if images exist", func(t *testing.T) {
		res := blurValidator.ValidateImagesExist(make([]*proto.ImageForBlur, 1))
		assertSuccess(t, res)
	})
}

func TestBlurValidator_ValidateImagesLimits(t *testing.T) {
	setupBlurServiceServerTest()
	defer teardownBlurServiceServerTest()

	t.Run("should return MaxImagesExceeded error if too many images provided", func(t *testing.T) {
		res := blurValidator.ValidateImagesLimits(make([]*proto.ImageForBlur, config.GetConfig().MaxImagesPerRequest+1))
		assertErrorResponse(t, res, proto.BlurErrorCode_MaxImagesExceeded)
	})

	t.Run("should return success if images count is valid", func(t *testing.T) {
		res := blurValidator.ValidateImagesLimits(make([]*proto.ImageForBlur, config.GetConfig().MaxImagesPerRequest))
		assertSuccess(t, res)
	})
}

func TestBlurValidator_ValidateImage(t *testing.T) {
	setupBlurServiceServerTest()
	defer teardownBlurServiceServerTest()

	t.Run("should return MaxSizeExceeded error if image size exceeds limit", func(t *testing.T) {
		res := blurValidator.ValidateImage(&proto.ImageForBlur{
			Content: make([]byte, config.GetConfig().MaxImageSize+1),
		})
		assertErrorImage(t, res, proto.BlurredImage_MaxSizeExceeded)
	})

	t.Run("should return InvalidImage error if image content is invalid", func(t *testing.T) {
		res := blurValidator.ValidateImage(&proto.ImageForBlur{
			Content: []byte("invalid image content"),
		})
		assertErrorImage(t, res, proto.BlurredImage_InvalidImage)
	})

	t.Run("should return success if image is valid", func(t *testing.T) {
		res := blurValidator.ValidateImage(&proto.ImageForBlur{
			Content: generateMockImage(),
		})
		assertSuccess(t, res)
	})
}

func TestBlurValidator_ValidateImageSize(t *testing.T) {
	setupBlurServiceServerTest()
	defer teardownBlurServiceServerTest()

	t.Run("should return MaxSizeExceeded error if image size exceeds limit", func(t *testing.T) {
		res := blurValidator.ValidateImageSize(&proto.ImageForBlur{
			Content: make([]byte, config.GetConfig().MaxImageSize+1),
		})
		assertErrorImage(t, res, proto.BlurredImage_MaxSizeExceeded)
	})

	t.Run("should return success if image size is valid", func(t *testing.T) {
		res := blurValidator.ValidateImageSize(&proto.ImageForBlur{
			Content: make([]byte, config.GetConfig().MaxImageSize),
		})
		assertSuccess(t, res)
	})
}

func TestBlurValidator_ValidateImageContent(t *testing.T) {
	setupBlurServiceServerTest()
	defer teardownBlurServiceServerTest()

	t.Run("should return InvalidImage error if image content is invalid", func(t *testing.T) {
		res := blurValidator.ValidateImageContent(&proto.ImageForBlur{
			Content: []byte("invalid image content"),
		})
		assertErrorImage(t, res, proto.BlurredImage_InvalidImage)
	})

	t.Run("should return success if image content is valid", func(t *testing.T) {
		res := blurValidator.ValidateImageContent(&proto.ImageForBlur{
			Content: generateMockImage(),
		})
		assertSuccess(t, res)
	})
}

func assertSuccess[T any](t *testing.T, res BlurValidationResult[T]) {
	if !res.Success {
		t.Fatalf("expected success, got error")
	}
	if res.ResWithErr != nil {
		t.Fatalf("expected nil error, got %v", res.ResWithErr)
	}
}
func assertErrorResponse(t *testing.T, res BlurValidationResult[*proto.BlurImagesResponse], expectedErrorCode proto.BlurErrorCode) {
	if res.Success {
		t.Fatalf("expected error, got success")
	}
	if res.ResWithErr.Error.Code != expectedErrorCode {
		t.Fatalf("expected %v error, got %v", expectedErrorCode, res.ResWithErr.Error.Code)
	}
}
func assertErrorImage(t *testing.T, res BlurValidationResult[*proto.BlurredImage], expectedErrorCode proto.BlurredImage_ErrorCode) {
	if res.Success {
		t.Fatalf("expected error, got success")
	}
	if res.ResWithErr.Error.Code != expectedErrorCode {
		t.Fatalf("expected %v error, got %v", expectedErrorCode, res.ResWithErr.Error.Code)
	}
}

func setupBlurValidatorTest() {
	_, err := config.LoadConfig()
	if err != nil {
		fmt.Println("failed to load config, got:", err)
		os.Exit(1)
	}

	blurValidator = NewBlurValidator()
}

func teardownBlurValidatorTest() {
}

func generateMockImage() []byte {
	width := 100
	height := 100

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			img.Set(x, y, image.Black)
		}
	}

	imagesService := images.NewImagesService()
	imgBytes, err := imagesService.ImageToBytes(img)
	if err != nil {
		fmt.Println("failed to convert image to bytes")
		os.Exit(1)
	}

	return *imgBytes
}
