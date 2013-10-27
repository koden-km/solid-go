package solid

// TODO:
// Probably just use functions on uint16's and uint32's instead of this custom type.

// Packed bitfield
type BitField32 uint32

// Get the masked bits of bitfield at offset.
func (x *BitField32) GetRaw(mask, offset BitField32) BitField32 {
	return (*x & mask) >> offset
}

// Get the masked bits of bitfield at offset.
func (x *BitField32) SetRaw(value, mask, offset BitField32) {
	// TODO: my bitfield knowledge is rusty, not sure if this is even close to right. Need tests for all this!
	shifted := (value & mask) << offset
	*x = *x &^ shifted
}

// Get the masked bits of bitfield at offset.
func (x *BitField32) GetUint(mask, offset uint32) uint32 {
	return uint32(x.GetRaw(BitField32(mask), BitField32(offset)))
}

// Get the masked bits of bitfield at offset.
// func (x *BitField32) GetInt(mask, offset uint32) int32 {
//     // TODO: deal with setting/unsetting sign?
//     return int32(x.GetRaw(mask, offset))
// }

// Set the masked bits of value into the bitfield at offset.
func (x *BitField32) SetUint(value, mask, offset uint32) {
	// shifted := (value & mask) << offset
	// *x = *x &^ shifted

	x.SetRaw(BitField32(value), BitField32(mask), BitField32(offset))
}

// Set the masked bits of value into the bitfield at offset.
// func (x *BitField32) SetInt(value, mask, offset uint32) {
//     // TODO: pack signed value into bitfield
//     x.SetRaw(uint32(value), mask, offset)
// }
