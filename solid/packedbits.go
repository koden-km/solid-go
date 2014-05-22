package solid

// TODO:
// Possible bit packing lib: https://github.com/joshlf13/gopack
// Move all this into its own (packedbits) package?
// Check that these bit field operations are correct (write tests!)
// Will need little and big endian versions of these?

// Get a packed value from the bitfield at the masked offset.
func GetUInt16(bitfield, mask, offset uint16) uint16 {
	// TODO
	return 0
	// value := bitfield & (mask << offset)
	// return value >> offset
}

// Return the bitfield with the new value packed in at the masked offset.
func SetUint16(value, bitfield, mask, offset uint16) uint16 {
	// TODO
	return 0
	// bitfield &= ~(mask << offset)
	// return bitfield | (value << offset)
}

// Get a packed value from the bitfield at the masked offset.
func GetUInt32(bitfield, mask, offset uint32) uint32 {
	// TODO
	return 0
	// value := bitfield & (mask << offset)
	// return value >> offset
}

// Return the bitfield with the new value packed in at the masked offset.
func SetUint32(value, bitfield, mask, offset uint32) uint32 {
	// TODO
	return 0
	// bitfield &= ~(mask << offset)
	// return bitfield | (value << offset)
}
