package util

const (
	USD = "USD"
	EUR = "EUR"
	RFW = "RFW"
	CAD = "CAD"
)

func IsSupportedCurrency(currency string) bool{
	switch(currency){
	case USD,EUR,RFW,CAD:
		return true
	}
	return false
}
