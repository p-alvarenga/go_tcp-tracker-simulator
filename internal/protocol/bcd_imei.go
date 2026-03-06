package protocol

func BcdToASCII(imei string) ([]byte, error) {
	if len(imei) == 15 {
		imei += "0"
	}

	bcdImei := make([]byte, 0, 8)

	for i := 0; i < len(imei); i += 2 {
		b := (imei[i]-'0')<<4 | (imei[i+1] - '0')
		bcdImei = append(bcdImei, b)
	}

	return bcdImei, nil
}
