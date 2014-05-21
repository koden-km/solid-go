package solid

import (
	"fmt"
	//"image/Color"
)

type LeafPos uint32

const (
	LeafPosUpperNorthEast LeafPos = iota
	LeafPosUpperNorthWest
	LeafPosUpperSouthEast
	LeafPosUpperSouthWest
	LeafPosLowerNorthEast
	LeafPosLowerNorthWest
	LeafPosLowerSouthEast
	LeafPosLowerSouthWest
)

// Octree data structure.
type Octree struct {
	leaves *[8]Octree
	voxel *Voxel
	// averageColor color.RGBA    // Average color of child voxel materials. TODO: Not sure if this should be stored always (faster, more memory), or just looked up when needed (slower, less memory)
}

// Create new Octree.
func NewOctree() *Octree {
	return &Octree{}
}

func (t *Octree) hasLeaves() bool {
	return t.leaves == nil
}

func (t *Octree) createLeaves() {
	if t.leaves == nil {
		t.leaves = new([8]Octree)
	}
}

// String representation for debug.
func (t *Octree) String() string {
	return fmt.Sprintf("%#v", t)
}

func (t *Octree) updateAverageColor() {
	// t.averageColor = blend t.leaves[...]
	// loop and do each in a go routine (concurrency)
}

func (t *Octree) SetAtIndex(index LeafPos, voxel *Voxel) {
	t.createLeaves()
	t.leaves[index].voxel = voxel
	go t.updateAverageColor()
}

func (t *Octree) indexFromPosition(myPos Vector3, atPos Vector3) LeafPos {
	return LeafPosUpperNorthEast // TODO: Calc index using position offset to get direction
}

// Position is passed down during traversal to save storage memory, it doesnt need to exist at every node.
func (t *Octree) SetAtPosition(myPos Vector3, atPos Vector3, currentDepth uint32, maxDepth uint32, voxel *Voxel) {
	var index = t.indexFromPosition(myPos, atPos)
	if currentDepth != maxDepth || myPos == atPos {
		t.SetAtIndex(index, voxel)
		return
	}

	// TODO: Finish vector lib
	// var leafPos = atPos.Subtract(myPos)  // TODO: Not sure if correct. trying to get the local offset or something...
	// t.leaves[index].SetAtPosition(leafPos, atPos, currentDepth + 1, maxDepth, voxel)
}

// TODO: Return on a channel?
func (t *Octree) GetAtPosition(myPos Vector3, atPos Vector3, currentDepth uint32, maxDepth uint32, voxel *Voxel) {
	// ...
}
