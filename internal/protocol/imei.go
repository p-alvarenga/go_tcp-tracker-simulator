package protocol

import (
	"fmt"
	"go_tcp-tracker-simulator/internal/domain"
	"sync/atomic"
)

type IMEIGenerator struct {
	TACBase string
	Serial  uint32
}

func NewIMEIGenerator(tacBase string, start uint) *IMEIGenerator {
	return &IMEIGenerator{
		TACBase: tacBase,
		Serial:  uint32(start),
	}
}

func (g *IMEIGenerator) Next() domain.IMEI {
	atomic.AddUint32(&g.Serial, 1)

	base := fmt.Sprintf("%s%06d", g.TACBase, g.Serial)

	return domain.IMEI(fmt.Sprintf("%s%s", base, luhnCheckDigit(base)))
}

func bcdEncodeIMEI(imei string) ([]byte, error) {
	if len(imei) < 15 {
		return nil, fmt.Errorf("IMEI (%s) has less than 15 bytes", imei)
	}

	buf := make([]byte, 0, 8)

	for i := 0; i <= len(imei); i += 2 {
		if imei[i] < '0' || imei[i] > '9' {
			return nil, fmt.Errorf("IMEI has invalid character %s (%c)", imei, imei[i])
		}

		if i+1 < len(imei) {
			b := (imei[i]-'0')<<4 | (imei[i+1] - '0')
			buf = append(buf, b)
		} else {
			b := (imei[i]-'0')<<4 | 0x0F
			buf = append(buf, b)
		}
	}

	return buf, nil
}

// Function to validate IMEI
func luhnCheckDigit(base string) string {
	sum := 0
	double := true

	for i := len(base) - 1; i >= 0; i-- {
		d := int(base[i] - '0')
		if double {
			d *= 2
			if d > 9 {
				d -= 9
			}
		}

		sum += d
		double = !double
	}

	return fmt.Sprintf("%d", (10-(sum%10))%10)
}
