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
			ID:           id,
			Index:        index,
			Contents:     section[0:3968],
			CheckSum:     checkSum,
			SectionIndex: ix,
		}
	}
	return sections
}

type Section struct {
	ID           int
	Index        uint32
	Contents     []byte
	CheckSum     uint16
	SectionIndex int
}

func CalculateChecksum(data []byte) uint16 {
	var checksum uint32 = 0
	for i := 0; i < len(data); i += 4 {
		if i+4 <= len(data) {
			word := binary.LittleEndian.Uint32(data[i : i+4])
			checksum += word
		}
	}
	checksum = (checksum & 0xFFFF) + (checksum >> 16)
	return uint16(checksum & 0xFFFF)
}
