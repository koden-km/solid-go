package solid

// Materials will be created and stored in an indexed lookup table.

import (
	"image/Color"
)

type Material struct {
	name     string
	color    color.RGBA  // TODO: Is this "diffuse" in normal lighting terms?
	matter   Matter
	density  float32
	friction float32
}
