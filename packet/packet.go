package packet

import "encoding/binary"

// ByteOrder is the byte order used for reading binary packet data.
// It is set to binary.LittleEndian.
var ByteOrder = binary.LittleEndian
