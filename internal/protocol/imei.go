package protocol

func IMEIToBcd(imei string) ([]byte, error) {
	if len(imei) == 15 {
		imei += "0"
	}

	bcd := make([]byte, 0, 8)

	for i := 0; i < len(imei); i += 2 {
		b := (imei[i]-'0')<<4 | (imei[i+1] - '0')
		bcd = append(bcd, b)
	}

	return bcd, nil
}

func CheckIMEI(imei string) bool {
	return true //
}
