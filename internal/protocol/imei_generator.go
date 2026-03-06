package protocol

import (
	"fmt"

	"github.com/p-alvarenga/go_tcp-tracker-simulator/internal/domain/device"
)

type IMEIGenerator struct {
	TAC    string
	Serial int
}

func NewIMEIGenerator(tac string, start int) *IMEIGenerator {
	return &IMEIGenerator{
		TAC:    tac,
		Serial: start,
	}
}

func (g *IMEIGenerator) Next() device.IMEI {
	g.Serial++

	serialStr := fmt.Sprintf("%06d", g.Serial) // 16 == "000016", 12345 == "012345"
	base := g.TAC + serialStr

	checkDigit := luhnCheckDigit(base)

	return device.IMEI(base + checkDigit)
}

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
