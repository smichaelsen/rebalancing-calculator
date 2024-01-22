package io

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/smichaelsen/rebalancing-calculator/calculator"
	"github.com/smichaelsen/rebalancing-calculator/structs"
	"os"
	"strconv"
	"strings"
)

func RenderTable(calculatorInstance *calculator.InvestmentCalculator, printableResults []structs.Category) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Category", "Target Allocation (%)", "Amount Before", "Amount Added", "Total Investment After", "Achieved Allocation (%)", "Investment Percentage (%)"})

	overallTotalBefore := 0.0
	overallTotalInvestment := 0.0

	for _, category := range printableResults {
		totalInvestmentAfter := category.Current + category.Investment
		table.Append([]string{
			category.Name,
			formatFloat(category.Target, 2),
			formatFloat(category.Current, 2),
			formatFloat(category.Investment, 2),
			formatFloat(totalInvestmentAfter, 2),
			formatFloat(category.AchievedAllocation, 2),
			formatFloat((category.Investment/calculatorInstance.AmountToInvest)*100, 2),
		})
		overallTotalBefore += category.Current
		overallTotalInvestment += totalInvestmentAfter
	}
	table.SetFooter([]string{
		fmt.Sprintf("%d categories", len(printableResults)),
		formatFloat(100.0, 2),
		formatFloat(overallTotalBefore, 2),
		formatFloat(calculatorInstance.AmountToInvest, 2),
		formatFloat(overallTotalInvestment, 2),
		formatFloat(100.0, 2),
		formatFloat(100.0, 2),
	})
	table.Render()
}

func formatFloat(num float64, precision int) string {
	// First, format the number with the dot as the decimal separator
	str := strconv.FormatFloat(num, 'f', precision, 64)

	// Split it into the integer and decimal parts
	parts := strings.Split(str, ".")

	// Reverse the integer part for easier processing
	reversedIntPart := reverse(parts[0])

	// Insert dots every three characters
	reversedIntPartWithDots := ""
	for i, ch := range reversedIntPart {
		if i > 0 && i%3 == 0 {
			reversedIntPartWithDots += "."
		}
		reversedIntPartWithDots += string(ch)
	}

	// Reverse the integer part back to its original order
	intPartWithDots := reverse(reversedIntPartWithDots)

	// Combine the integer part with the decimal part using ',' as the decimal separator
	result := intPartWithDots
	if precision > 0 {
		result += "," + parts[1]
	}

	return result
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
