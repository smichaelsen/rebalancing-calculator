package main

import (
    "flag"
    "fmt"
    "github.com/smichaelsen/rebalancing-calculator/calculator"
    "github.com/smichaelsen/rebalancing-calculator/io"
)

func main() {
    calculatorInstance := calculator.NewInvestmentCalculator()

    fmt.Println("Rebalancing Calculator")
    fmt.Println("")

    var categoriesFlags []string
    flag.Var((*flagSlice)(&categoriesFlags), "category", "Specify investment categories and target allocations")
    flag.Parse()

    if len(categoriesFlags) > 0 {
        for _, category := range categoriesFlags {
            io.ParseCategoryFromFlag(calculatorInstance, category)
        }
    } else {
        io.InputCategories(calculatorInstance)
    }
    if len(calculatorInstance.Categories) < 2 {
        fmt.Println("Enter at least 2 categories to start the calculation.")
        return
    }

    io.InputHoldings(calculatorInstance)

    io.InputAmountToInvest(calculatorInstance)
    if calculatorInstance.AmountToInvest < 0.01 {
        fmt.Println("Enter at least an amount of 0.01 to start the calculation.")
        return
    }

    fmt.Println("")
    io.RenderTable(calculatorInstance, calculatorInstance.CalculateAllocation())
}

// Custom flag type to support multiple occurrences of the same flag
type flagSlice []string

func (f *flagSlice) String() string {
    return fmt.Sprintf("%s", *f)
}

func (f *flagSlice) Set(value string) error {
    *f = append(*f, value)
    return nil
}
