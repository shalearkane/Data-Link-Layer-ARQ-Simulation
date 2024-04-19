package transmitter

import (
	"encoding/binary"

	"github.com/snksoft/crc"
)

func AddCRC(s *[]byte) []byte {
	hash := crc.NewHash(crc.CCITT)
	hash.Reset() // Reset the hash instance for each chunk
	hash.Update(*s)
	crc := hash.CRC()
	i := binary.LittleEndian.AppendUint64(*s, crc)

	return i
}
