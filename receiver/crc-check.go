package receiver

import (
	"encoding/binary"

	"github.com/snksoft/crc"
)

func CheckCRC(b []byte) bool {
	array_len := len(b)
	if array_len != 72 {
		return false
	}

	crc_expected := b[(array_len - 8):]

	hash := crc.NewHash(crc.CCITT)
	hash.Reset() // Reset the hash instance for each chunk
	hash.Update(b[:(array_len - 8)])
	crc_actual := hash.CRC()

	return crc_actual == binary.LittleEndian.Uint64(crc_expected)
}
