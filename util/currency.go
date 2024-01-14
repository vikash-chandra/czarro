package util

const (
	INR = "INR"
)

func IsSupportCurrency(currency string) bool {
	switch currency {
	case INR:
		return true
	}
	return false
}
