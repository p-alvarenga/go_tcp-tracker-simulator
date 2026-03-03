package protocol

import (
	"encoding/binary"
)

func CalculateCrc(data []byte) uint16 {
	var crc uint16 = 0xffff

	for _, b := range data {
		crc ^= uint16(b) << 8

		for range 8 {
			if crc&0x8000 != 0 {
				crc = (crc << 1) ^ 0x1021
			} else {
				crc <<= 1
			}
		}
	}

	return crc
}

func ValidateCrc(data []byte) bool {
	crc := binary.BigEndian.Uint16(data[6:8])

	return crc == CalculateCrc(data[2:6])
}
