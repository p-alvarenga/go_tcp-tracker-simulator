package protocol

func CalculateCRC(slice []byte) uint16 {
	var crc uint16 = 0xffff

	for _, b := range slice {
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

func ValidateCRC(crc uint16, slice []byte) bool {
	return crc == CalculateCRC(slice)
}
