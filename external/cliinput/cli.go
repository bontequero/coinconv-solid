package cliinput

import (
	"fmt"
	"os"
	"strconv"
)

type QueryInput struct {
	Amount float64
	From   string
	To     string
}

func GetConversionInput() (*QueryInput, error) {
	convertMaxArgs := 3

	args := os.Args[1:]
	if len(args) < convertMaxArgs {
		return nil, fmt.Errorf("insufficient number of arguments: want=%v, got=%v", convertMaxArgs, len(args))
	}

	var amount float64
	var err error
	if amount, err = strconv.ParseFloat(args[0], 32); err != nil {
		return nil, fmt.Errorf("cant parse amount: %v", args[0])
	}

	from := args[1]
	to := args[2]

	return &QueryInput{Amount: amount, From: from, To: to}, nil
}
