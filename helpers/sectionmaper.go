package helpers

import "encoding/binary"

func CreateSectionsMap(savedata []byte) map[int]Section {
	sections := make(map[int]Section)
	for i := 0; i < 14; i++ {
		ix := i * 4096
		section := savedata[ix : ix+4096]
		footer := section[4084:]
		id := int(binary.LittleEndian.Uint16(footer[0:2]))
		index := binary.LittleEndian.Uint32(footer[4:8])
		checkSum := binary.LittleEndian.Uint16(footer[2:4])
		sections[id] = Section{
			ID:       id,
			Index:    index,
			Contents: section[0:3968],
			CheckSum: checkSum,
		}
	}
	return sections
}

type Section struct {
	ID       int
	Index    uint32
	Contents []byte
	CheckSum uint16
}
