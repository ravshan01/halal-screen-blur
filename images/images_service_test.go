package images

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"image"
	"os"
	"testing"
)

var imagesService = NewImagesService()

var imgBytes []byte
var img image.Image

func TestImagesService_CheckIsValidImage(t *testing.T) {
	setupImagesServiceTest()
	defer teardownImagesServiceTest()

	t.Run("valid image", func(t *testing.T) {
		assert.Equal(t, true, imagesService.CheckIsValidImage(imgBytes), "expected image to be valid")
	})

	t.Run("invalid image", func(t *testing.T) {
		assert.Equal(t, false, imagesService.CheckIsValidImage([]byte{}), "expected image to be invalid")
	})
}

func TestImagesService_Blur(t *testing.T) {
	setupImagesServiceTest()
	defer teardownImagesServiceTest()

	t.Run("should return InvalidBlurPercentageError if percentage < 0 or > 100 ", func(t *testing.T) {
		_, err := imagesService.Blur(img, -1)
		assert.NotNil(t, err, "expected error, got nil")
		assert.ErrorIs(t, err, InvalidBlurPercentageError, "expected error to be InvalidBlurPercentageError")

		_, err = imagesService.Blur(img, 101)
		assert.NotNil(t, err, "expected error, got nil")
		assert.ErrorIs(t, err, InvalidBlurPercentageError, "expected error to be InvalidBlurPercentageError")
	})

	// Check './mock/processed/mock-image__blurred.jpg' to see the blurred image
	t.Run("should return blurred image", func(t *testing.T) {
		blurredImgPointer, err := imagesService.Blur(img, 50)
		assert.Nil(t, err, "expected no error, got %v", err)

		blurredImg := *blurredImgPointer
		assert.NotNil(t, blurredImg, "expected blurred image to be returned")
		assert.Equal(t, img.Bounds().Dx(), blurredImg.Bounds().Dx(), "expected blurred image to have same width as original image")
		assert.Equal(t, img.Bounds().Dy(), blurredImg.Bounds().Dy(), "expected blurred image to have same height as original image")

		saveImage(blurredImg, "blurred")
	})
}

func TestImagesService_Crop(t *testing.T) {
	setupImagesServiceTest()
	defer teardownImagesServiceTest()

	// Check './mock/processed/mock-image__cropped.jpg' to see the cropped image
	t.Run("should return cropped image", func(t *testing.T) {
		x1, y1, width, height := 0, 0, 400, 200
		rect := image.Rect(x1, y1, x1+width, y1+height)

		croppedImgPointer, err := imagesService.Crop(img, rect)
		assert.Nil(t, err, "expected no error, got %v", err)

		croppedImg := *croppedImgPointer
		assert.NotNil(t, croppedImg, "expected cropped image to be returned")
		assert.Equal(t, width, croppedImg.Bounds().Dx(), "expected cropped image to have width %d, got width %d", width, croppedImg.Bounds().Dx())
		assert.Equal(t, height, croppedImg.Bounds().Dy(), "expected cropped image to have height %d, got height %d", height, croppedImg.Bounds().Dy())

		saveImage(croppedImg, "cropped")
	})
}

func TestImagesService_Paste(t *testing.T) {
	setupImagesServiceTest()
	defer teardownImagesServiceTest()

	// Check './mock/processed/mock-image__pasted.jpg' to see the pasted image
	t.Run("should return pasted image", func(t *testing.T) {
		cropRect := image.Rect(0, 0, 400, 400)
		croppedImgPointer, cropErr := imagesService.Crop(img, cropRect)
		croppedImg := *croppedImgPointer

		assert.NotNil(t, croppedImg, "expected cropped image to be returned")
		assert.Nil(t, cropErr, "expected no error, got %v", cropErr)

		pasteRect := image.Point{X: 200, Y: 200}
		pastedImgPointer, err := imagesService.Paste(img, croppedImg, pasteRect)
		assert.Nil(t, err, "expected no error, got %v", err)

		pastedImg := *pastedImgPointer
		assert.NotNil(t, pastedImg, "expected pasted image to be returned")
		assert.Equal(t, img.Bounds().Dx(), pastedImg.Bounds().Dx(), "expected pasted image to have same width as original image")
		assert.Equal(t, img.Bounds().Dy(), pastedImg.Bounds().Dy(), "expected pasted image to have same height as original image")

		saveImage(pastedImg, "pasted")
	})
}

func TestImagesService_BytesToImage(t *testing.T) {
	setupImagesServiceTest()
	defer teardownImagesServiceTest()

	t.Run("should return image", func(t *testing.T) {
		imgPointer, err := imagesService.BytesToImage(imgBytes)
		assert.Nil(t, err, "expected no error, got %v", err)
		assert.NotNil(t, imgPointer, "expected image to be returned")
	})

	t.Run("should return error if image is invalid", func(t *testing.T) {
		_, err := imagesService.BytesToImage([]byte{})
		assert.NotNil(t, err, "expected error, got nil")
	})
}

func TestImagesService_ImageToBytes(t *testing.T) {
	setupImagesServiceTest()
	defer teardownImagesServiceTest()

	t.Run("should return bytes", func(t *testing.T) {
		bytesPointer, err := imagesService.ImageToBytes(img)
		assert.Nil(t, err, "expected no error, got %v", err)
		assert.NotNil(t, bytesPointer, "expected bytes to be returned")
	})
}

func setupImagesServiceTest() {
	bytes, err := os.ReadFile("./mock/mock-image.jpg")
	if err != nil {
		fmt.Println("failed to read image file")
		os.Exit(1)
	}
	imgBytes = bytes

	origImg, toImgErr := imagesService.BytesToImage(bytes)
	if toImgErr != nil {
		fmt.Println("failed to convert bytes to image")
		os.Exit(1)
	}

	img = *origImg
}

func teardownImagesServiceTest() {
}

func saveImage(img image.Image, action string) {
	bytes, toBytesErr := imagesService.ImageToBytes(img)
	if toBytesErr != nil {
		fmt.Printf("failed to convert image to bytes for %s\n", action)
		return
	}

	err := os.WriteFile(fmt.Sprintf("./mock/processed/mock-image__%s.png", action), *bytes, 0644)
	if err != nil {
		fmt.Printf("failed to write %s image\n", action)
	}
}
