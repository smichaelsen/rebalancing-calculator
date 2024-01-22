package io

import (
	"bufio"
	"fmt"
	"github.com/smichaelsen/rebalancing-calculator/calculator"
	"github.com/smichaelsen/rebalancing-calculator/structs"
	"os"
	"strconv"
	"strings"
)

func InputCategories(calculatorInstance *calculator.InvestmentCalculator) {
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

		var category = structs.Category{
			Current:    currentInvestment,
			Name:       categoryName,
			Target:     targetAllocation,
			Locked:     false, // Default value, it will be calculated later
			Investment: 0,     // Default value, it will be calculated later
		}

		calculatorInstance.AddCategory(category)
	}
}

func InputAmountToInvest(calculatorInstance *calculator.InvestmentCalculator) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter amount to invest: ")
	scanner.Scan()
	amountToInvestString := strings.TrimSpace(scanner.Text())
	amountToInvestString = strings.Replace(amountToInvestString, ",", ".", -1)
	amountToInvestFloat, _ := strconv.ParseFloat(amountToInvestString, 64)
	calculatorInstance.SetAmountToInvest(amountToInvestFloat)
}
