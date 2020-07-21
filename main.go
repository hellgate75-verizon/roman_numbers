package main

import (
	"flag"
	"fmt"
	"github.com/hellgate75/roman_numbers/model"
	"os"
)

var in int64
var max int64
var fg *flag.FlagSet

func init() {
	model.Init()
	fg = flag.NewFlagSet("roman-numbers", flag.ContinueOnError)
	fg.Int64Var(&in, "in", 0, fmt.Sprintf("Number for the conversion [1...%v]", model.NumberConversionLimit))
	fg.Int64Var(&max, "max", 0, fmt.Sprint("Change limit to a new number"))
}

func main() {
	// Parse input arguments
	err := 	fg.Parse(os.Args[1:])
	if err != nil {
		fmt.Println("Error: Required arguments missing or wrong format")
		fmt.Println("roman-number -in <number>")
		fg.PrintDefaults()
		os.Exit(1)
	}
	if max > 0 {
		fmt.Printf("Using upper limit %v instead of %v\n", max, model.NumberConversionLimit)
		model.NumberConversionLimit = max
	}
	// Check Pre-Conditions
	state := model.CheckNumber(in)
	if ! state {
		fmt.Println("Error: Number of of bounds")
		fmt.Println("roman-number -in <number>")
		fg.PrintDefaults()
		os.Exit(2)
	}
	out := model.Convert(in)
	fmt.Printf("Roman number for %v is %s\n", in, out)
}