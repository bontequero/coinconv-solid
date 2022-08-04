package convert

import (
	"fmt"

	"convertor/entities"
)


type convertorWithCommission struct {
	convertor Convertor
	commision entities.Commission
}

func NewWithCommision(cv Convertor, commission entities.Commission) convertorWithCommission {
	return convertorWithCommission{
		convertor: cv,
		commision: commission,
	}
}

func (cm *convertorWithCommission) Convert(cd entities.ConversionData) (entities.ConversionResult, error) {
	result, err := cm.convertor.Convert(cd)
	if err != nil {
		return entities.ConversionResult{}, fmt.Errorf("cant convert: %w", err)
	}

	var resultWithCommission entities.ConversionResult
	resultWithCommission.Result = result.Result + (result.Result / 100.0) * float64(cm.commision)
	return resultWithCommission, nil
}

var _ Convertor = (*convertorWithCommission)(nil)
