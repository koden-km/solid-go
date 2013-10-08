package solid

import (
	"image/RGBA"
)

type Material struct {
	color    RGBA
	matter   Matter
	density  float
	friction float
}
