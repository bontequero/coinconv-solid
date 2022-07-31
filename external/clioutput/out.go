package clioutput

import (
	"fmt"

	"convertor/entities"
)

func PrintResult(r entities.ConversionResult) {
	fmt.Println(r.Result)
}
