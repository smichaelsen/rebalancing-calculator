package main

import (
	"fmt"
	"github.com/smichaelsen/rebalancing-calculator/calculator"
	"github.com/smichaelsen/rebalancing-calculator/io"
)

func main() {
	calculatorInstance := calculator.NewInvestmentCalculator()

    fmt.Println("Rebalancing Calculator")
    fmt.Println("")

	io.InputCategories(calculatorInstance)
	if len(calculatorInstance.Categories) < 2 {
		fmt.Println("Enter at least 2 categories to start the calculation.")
		return
	}

	io.InputAmountToInvest(calculatorInstance)
	if calculatorInstance.AmountToInvest < 0.01 {
		fmt.Println("Enter at least an amount of 0.01 to start the calculation.")
		return
	}

    fmt.Println("")
	io.RenderTable(calculatorInstance, calculatorInstance.CalculateAllocation())
}
