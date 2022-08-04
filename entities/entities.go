package entities

type ConversionData struct {
	Amount float64
	From   Coin
	To     Coin
}

type ConversionResult struct {
	Result float64
}

type Coin string

type Commission float64
