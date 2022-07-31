package convert

import (
	"convertor/entities"
)

type PriceConvertor interface {
	ConvertPrice(entities.ConversionData) (entities.ConversionResult, error)
}

type priceConvertor struct {
	PriceConvertor
}

func New(c PriceConvertor) priceConvertor {
	return priceConvertor{c}
}

func (c priceConvertor) Convert(cd entities.ConversionData) (entities.ConversionResult, error) {
	return c.ConvertPrice(cd)
}
