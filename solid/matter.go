package solid

// Matter is used in Material definitions.

type Matter uint32

const (
	MatterVacuum Matter = 0
	MatterGas           = 1
	MatterLiquid        = 2
	MatterSolid         = 3
	MatterPlasma        = 4  // Lookup Plasma again. If its a cross between the other types, then maybe bit masks might work better here.
)
