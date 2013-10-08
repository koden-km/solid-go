package solid

// Material:
// Should i keep a pointer to the actual material? ~64 bit pointer vs a ~16 bit index into material list (memory usage vs time to lookup).
// The material will be persisted as the numberic index into the material list.

// Attributes:
// - 5 bits = Surface normal type. 32 possible values and only need 26 options for really basic normals?
//            Something like this:
//            [straight] Up, Down, North, South, East, West,
//            [diagonal pitch] Up+North, Down+North, Down+South, Up+South,
//            [diagonal yaw] North+East, South+East, South+West, North+West,
//            [diagonal roll] Up+East, Down+East, Down+West, Up+West,
//            [diagonal pitch+yaw+roll 1] Up+North+East, Down+North+East, Down+South+West, Up+South+West.
//            [diagonal pitch+yaw+roll 2] Up+North+West, Down+North+West, Down+South+East, Up+South+East.
// - backed light colour and shadow?

import (
	"fmt"
)

const (
	AttributeNormal     = 0x1F
	AttributeBakedLight = 0x1FFFE0
)

type Voxel struct {
	material   *Material
	attributes uint
}

func (v *Voxel) String() string {
	return fmt.Sprintf("%#v", v)
}

func (v *Voxel) Normal() Vector3 {
	return Vector3{0, 1, 0}

	// Something like this?
	// switch v.attributes & AttributeNormal {
	// case 0:
	// 	return Vector3{0, 1, 0}
	// case 1:
	// 	return Vector3{0, -1, 0}
	// case 2:
	// 	return Vector3{0, 0, 1}
	// case 3:
	// 	return Vector3{0, 0, -1}
	// case 4:
	// 	return Vector3{1, 0, 0}
	// case 5:
	// 	return Vector3{-1, 0, 0}
	// case 6:
	// ...
	// }
}
