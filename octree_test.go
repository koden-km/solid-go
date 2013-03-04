package solid

import (
	"testing"
)

func TestOctreeString(t *testing.T) {
	tree := new(Octree)
	if len(tree.String()) == 0 {
		t.Errorf("String() length is 0.")
	}
}

