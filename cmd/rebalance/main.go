package main

import (
	"bufio"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"os"
	"strconv"
	"strings"
)

func main() {
	calculator := NewInvestmentCalculator()

	inputCategories(calculator)
	if len(calculator.Categories) < 2 {
		fmt.Println("Enter at least 2 categories to start the calculation.")
		return
	}

	inputAmountToInvest(calculator)
	if calculator.AmountToInvest < 0.01 {
		fmt.Println("Enter at least an amount of 0.01 to start the calculation.")
		return
	}

	printableResults := calculator.CalculateAllocation()

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Category", "Target Allocation (%)", "Amount Before", "Amount Added", "Total Investment After", "Achieved Allocation (%)", "Investment Percentage (%)"})

	overallTotalBefore := 0.0
	overallTotalInvestment := 0.0

	for _, category := range printableResults {
		fmt.Printf("category target allocation: %v", category.Target)
		totalInvestmentAfter := category.Current + category.Investment
		table.Append([]string{
			category.Name,
			formatFloat(category.Target, 2),
			formatFloat(category.Current, 2),
			formatFloat(category.Investment, 2),
			formatFloat(totalInvestmentAfter, 2),
			formatFloat(category.AchievedAllocation, 2),
			formatFloat((category.Investment/calculator.AmountToInvest)*100, 2),
		})
		overallTotalBefore += category.Current
		overallTotalInvestment += totalInvestmentAfter
	}
	table.SetFooter([]string{
		fmt.Sprintf("%d categories", len(printableResults)),
		formatFloat(100.0, 2),
		formatFloat(overallTotalBefore, 2),
		formatFloat(calculator.AmountToInvest, 2),
		formatFloat(overallTotalInvestment, 2),
		formatFloat(100.0, 2),
		formatFloat(100.0, 2),
	})
	table.Render()

}

func inputCategories(calculator *InvestmentCalculator) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter category name (leave empty to finish): ")
		scanner.Scan()
		categoryName := strings.TrimSpace(scanner.Text())

		// Break the loop if the category name is empty
		if categoryName == "" {
			break
		}

		fmt.Printf("Enter amount already invested in %s: ", categoryName)
		scanner.Scan()
		currentInvestmentStr := strings.TrimSpace(scanner.Text())
		currentInvestmentStr = strings.Replace(currentInvestmentStr, ",", ".", -1)
		currentInvestment, err := strconv.ParseFloat(currentInvestmentStr, 64)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid number.")
			continue
		}

		fmt.Printf("Enter target allocation in %% for %s: ", categoryName)
		scanner.Scan()
		targetAllocationStr := strings.TrimSpace(scanner.Text())
		targetAllocationStr = strings.Replace(targetAllocationStr, ",", ".", -1)
		targetAllocation, err := strconv.ParseFloat(targetAllocationStr, 64)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid number.")
			continue
		}

		var category = Category{
			Current:    currentInvestment,
			Name:       categoryName,
			Target:     targetAllocation,
			Locked:     false, // Default value, it will be calculated later
			Investment: 0,     // Default value, it will be calculated later
		}

		calculator.AddCategory(category)
	}
}

func inputAmountToInvest(calculator *InvestmentCalculator) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter amount to invest: ")
	scanner.Scan()
	amountToInvestString := strings.TrimSpace(scanner.Text())
	amountToInvestString = strings.Replace(amountToInvestString, ",", ".", -1)
	amountToInvestFloat, _ := strconv.ParseFloat(amountToInvestString, 64)
	calculator.SetAmountToInvest(amountToInvestFloat)
}

type InvestmentCalculator struct {
	Categories     []Category
	AmountToInvest float64
}

func NewInvestmentCalculator() *InvestmentCalculator {
	return &InvestmentCalculator{
		Categories:     []Category{},
		AmountToInvest: 0,
	}
}

// AddCategory adds a new category to the calculator
func (ic *InvestmentCalculator) AddCategory(category Category) {
	ic.Categories = append(ic.Categories, category)
}

// SetAmountToInvest sets the amount to invest
func (ic *InvestmentCalculator) SetAmountToInvest(amount float64) {
	ic.AmountToInvest = amount
}

func (ic *InvestmentCalculator) CalculateAllocation() []Category {
	totalCurrentInvestment := 0.0
	adjustedTotalTargetPercentage := 0.0
	lockedInvestment := 0.0

	// Calculate the total current investment
	for _, category := range ic.Categories {
		totalCurrentInvestment += category.Current
	}

	totalFutureInvestment := totalCurrentInvestment + ic.AmountToInvest

	// First pass: Determine which categories are locked and calculate adjusted total target percentage
	for i, category := range ic.Categories {
		targetInvestment := totalFutureInvestment * (category.Target / 100)
		if category.Current >= targetInvestment {
			ic.Categories[i].Locked = true
			lockedInvestment += category.Current
		} else {
			adjustedTotalTargetPercentage += category.Target
		}
	}

	adjustedRemainingInvestment := totalFutureInvestment - lockedInvestment

	// Second pass: Calculate the amount to be added for each category
	for i, category := range ic.Categories {
		if category.Locked {
			continue
		}
		adjustedTarget := category.Target / adjustedTotalTargetPercentage
		ic.Categories[i].Investment = adjustedRemainingInvestment*adjustedTarget - category.Current
	}

	// set achieved allocation
	for i, category := range ic.Categories {
		ic.Categories[i].AchievedAllocation = ((category.Current + category.Investment) / totalFutureInvestment) * 100
	}

	return ic.Categories
}

type Category struct {
	Name               string
	Current            float64
	Target             float64
	Locked             bool
	Investment         float64
	AchievedAllocation float64
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
