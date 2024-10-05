package images

import "errors"

var InvalidImageError = errors.New("invalid image")
var InvalidBlurPercentageError = errors.New("invalid percentage, must be between 0 and 100")
