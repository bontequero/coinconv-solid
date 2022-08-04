package convert

import (
	"convertor/entities"
)

type Convertor interface {
	Convert(entities.ConversionData) (entities.ConversionResult, error)
}
