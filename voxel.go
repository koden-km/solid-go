package solid

// this could probably do the job of Material?

import (
    "fmt"
)

type Voxel struct {
	mat *Material
	// surface normal?
}

func (v *Voxel) String() string {
	return fmt.Sprintf("%#v", v)
}
