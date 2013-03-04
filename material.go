package solid

import (
	"image/RGBA"
)

type Material struct {
	matter Matter
	color RGBA
	friction float
}
