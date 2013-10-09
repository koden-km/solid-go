package solid

import (
	"image/Color"
)

type Material struct {
	name     string
	color    color.RGBA
	matter   Matter
	density  float32
	friction float32
}
