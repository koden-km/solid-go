package solid

type Matter uint32

const (
	MatterVacuum Matter = iota
	MatterPlasma
	MatterGas
	MatterLiquid
	MatterSolid
)
