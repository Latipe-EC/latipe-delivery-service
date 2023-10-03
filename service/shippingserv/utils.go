package shippingserv

import "time"

const (
	Suburb    = 1.5
	InnerCity = 2
)

func CalculateShippingCodes(src string, dest string, cost int) (int, time.Time) {
	now := time.Now()
	if src != dest {
		return int(Suburb * float64(cost)), now.Add(4 * 24 * time.Hour)
	}
	return int(InnerCity * float64(cost)), now.Add(7 * 24 * time.Hour)
}
