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

type Voxel struct {
	material   *Material
	attributes BitField32
}

const (
	attributeMaskNormal     = 0x1F
	attributeMaskBakedLight = 0x1FFFE0

	attributeOffsetNormal     = 0
	attributeOffsetBakedLight = 5
)

//var SurfaceNormalTable []Vector3 = {
//	Vector3{0, 1, 0},		// Up
//	Vector3{0, -1, 0},		// Down
//	// TODO: add the rest...
//}

func (v *Voxel) String() string {
	return fmt.Sprintf("%#v", v)
}

// Moving this to bitfield32.go
// func (v *Voxel) extractAttributeUint(mask, offset uint32) uint32 {
// 	return (v.attributes & mask) >> offset
// }

// Moving this to bitfield32.go
// func (v *Voxel) extractAttributeInt(mask, offset uint32) int32 {
// 	return int32(v.extractAttributeUint(mask, offset))
// }

func (v *Voxel) Normal() Vector3 {
	return Vector3{0, 1, 0}

	// return SurfaceNormalTable[v.attributes.GetUint(attributeMaskNormal, attributeOffsetNormal)]
}
