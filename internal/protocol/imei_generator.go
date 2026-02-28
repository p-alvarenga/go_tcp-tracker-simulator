package protocol

import (
	"fmt"
	"gt06_sim/internal/device"
)

type ImeiGenerator struct {
	tac    string
	serial int
}

func NewImeiGenerator(tac string, start int) *ImeiGenerator {
	return &ImeiGenerator{
		tac:    tac,
		serial: start,
	}
}

func (ig *ImeiGenerator) Next() device.Imei {
	ig.serial++

	serialStr := fmt.Sprintf("%06d", ig.serial) // 16 == "000016", 12345 == "012345"
	base := ig.tac + serialStr

	checkDigit := luhnCheckDigit(base)

	return device.Imei(base + checkDigit)
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
