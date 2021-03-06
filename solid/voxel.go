package solid

// Voxels are the Octree leaf payload.

// Material:
// Should i keep a pointer to the actual material? ~64 bit pointer vs a ~16 bit index into material list (memory usage vs time to lookup).
// The material will be persisted as the numberic index into the material list.

// Normals could be calculated by looking at adjacent voxels, but then nearly all surfaces would be flat.

// Attributes:
// Possible bit packing lib: https://github.com/joshlf13/gopack
// - 5 bits = Surface normal type. 32 possible values and only need 26 options for really basic normals?
//            Something like this:
//            [straight] Up, Down, North, South, East, West,
//            [diagonal pitch] Up+North, Down+North, Down+South, Up+South,
//            [diagonal yaw] North+East, South+East, South+West, North+West,
//            [diagonal roll] Up+East, Down+East, Down+West, Up+West,
//            [diagonal pitch+yaw+roll 1] Up+North+East, Down+North+East, Down+South+West, Up+South+West.
//            [diagonal pitch+yaw+roll 2] Up+North+West, Down+North+West, Down+South+East, Up+South+East.
// - baked light colour and shadow?
// - 6 bits = Attached/Bonded voxels/atoms. Bonded with: Up, Down, North, South, East, West.
//            How larger objects are made from small atoms with varing materials.
//            Actually might just make this a regular uint8 var on the type. The other 2 bits could be used for something though. Also aim for memory alignment?
//            Or maybe i should make this a ~5 bit index lookup like surface normal. Because the bonding should also support diagonal angles?
//            A bonding strength (for breaking appart) might be needed also. But that could be based on the material type itself?

import (
	"fmt"
)

type Voxel struct {
	material   uint16       // Index into material lookup table.
	attributes uint16       // Bitfield flags for this voxel. Using bitfields to pack data and keep memory use as low as possible.
	//attached   uint8		// Bits signalling the bonded attachment to nearby voxel/atoms in each direction (U,D,N,S,E,W).
}

// Attribute bitfield packing info.
// NOTE: If there are not many attributes needed, i may just use a simple type in the Voxel.
const (
	attributeMaskNormal     = 0x1F			// 0000 0000 0001 1111
	//attributeMaskBakedLight = 0xFFFE0		// 0000 0000 0000 0000

	attributeOffsetNormal     = 0
	//attributeOffsetBakedLight = 5
)

//var SurfaceNormalTable []Vector3 = {
//	Vector3{0, 1, 0},		// Up
//	Vector3{0, -1, 0},		// Down
//	// TODO: add the rest...
//}

func (v *Voxel) String() string {
	return fmt.Sprintf("%#v", v)
}

func (v *Voxel) MaterialIndex() uint16 {
	return v.Material;
}

func (v *Voxel) Material() *Material {
	// return MaterialTable[v.material]
	return nil;
}

func (v *Voxel) Attributes() uint16 {
	return v.attributes
}

func (v *Voxel) Normal() Vector3 {
	return Vector3{0, 1, 0}

	// return SurfaceNormalTable[GetUint32(v.attributes, attributeMaskNormal, attributeOffsetNormal)]
}
