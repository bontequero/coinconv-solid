package main

import (
	"log"
	"os"

	"convertor/entities"
	"convertor/external/cliinput"
	"convertor/external/clioutput"
	"convertor/external/coinmarketcap"
	"convertor/usecase/convert"
)

func main() {
	input, err := cliinput.GetConversionInput()
	if err != nil {
		log.Printf("bad input: %v", err)
		os.Exit(1)
	}
	toConvert := entities.ConversionData{
		Amount: input.Amount,
		From:   entities.Coin(input.From),
		To:     entities.Coin(input.To),
	}

	coinApi := coinmarketcap.New()
	convertor := convert.New(coinApi)

	result, err := convertor.Convert(toConvert)
	if err != nil {
		log.Printf("cant convert: %v", err)
		os.Exit(1)
	}

	clioutput.PrintResult(result)
}
