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
            Current:    0,
            Name:       categoryName,
            Target:     targetAllocation,
            Locked:     false, // Default value, it will be calculated later
            Investment: 0,     // Default value, it will be calculated later
        }

        calculatorInstance.AddCategory(category)
    }
}

func InputHoldings(calculatorInstance *calculator.InvestmentCalculator) {
    scanner := bufio.NewScanner(os.Stdin)

    categories := calculatorInstance.GetCategories()

    for i := range categories {
        fmt.Printf("Enter current value for %s: ", categories[i].Name)
        scanner.Scan()
        currentValueStr := strings.TrimSpace(scanner.Text())
        currentValueStr = strings.Replace(currentValueStr, ",", ".", -1)
        currentValue, err := strconv.ParseFloat(currentValueStr, 64)
        if err != nil {
            fmt.Println("Invalid input. Please enter a valid number.")
            continue
        }

        calculatorInstance.Categories[i].Current = currentValue
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

func ParseCategoryFromFlag(calculatorInstance *calculator.InvestmentCalculator, categoryFlag string) {
    categoryParts := strings.Split(categoryFlag, ";")

    if len(categoryParts) != 2 {
        fmt.Println("Invalid format for category:", categoryFlag)
        return
    }

    categoryName := categoryParts[0]
    targetAllocationStr := categoryParts[1]

    targetAllocation, err := strconv.ParseFloat(targetAllocationStr, 64)
    if err != nil {
        fmt.Println("Invalid target allocation for category:", categoryName)
        return
    }

    // Create a new category and add it to the calculator
    category := structs.Category{
        Name:       categoryName,
        Target:     targetAllocation,
        Current:    0, // Set default value for current investment
        Locked:     false,
        Investment: 0,
    }
    calculatorInstance.AddCategory(category)
}
