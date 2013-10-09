package solid

const (
	WorldChunkRadius = 16
)

type World struct {
	// map of octree chunk tree roots
}

// use a channel for return?
// func (w *World) GetVoxelAtPosition(worldPos Vector3, currentDepth uint32, maxDepth uint32) Voxel {
// 	// ...
// }

// func (w *World) SetVoxelAtPosition(worldPos Vector3, currentDepth uint32, maxDepth uint32, voxel Voxel) {
// 	// ...
// }
