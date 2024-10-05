package images

import (
	"fmt"
	"os"
	"testing"
)

var imagesService = ImagesService{}

var img []byte

func TestMain(m *testing.M) {
	bytes, readErr := os.ReadFile("./mock/street.jpg")
	if readErr != nil {
		fmt.Println("failed to read image file")
		os.Exit(1)
	}

	img = bytes
	os.Exit(m.Run())
}

func TestImagesService_CheckIsValidImage(t *testing.T) {

	t.Run("valid image", func(t *testing.T) {
		valid := imagesService.CheckIsValidImage(img)
		if !valid {
			t.Fatalf("expected image to be valid")
		}
	})

	t.Run("invalid image", func(t *testing.T) {
		valid := imagesService.CheckIsValidImage([]byte{})
		if valid {
			t.Fatalf("expected image to be invalid")
		}
	})
}

func TestImagesService_Blur(t *testing.T) {

}
