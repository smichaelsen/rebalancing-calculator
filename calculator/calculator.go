package calculator

import (
    "math"
    "github.com/smichaelsen/rebalancing-calculator/structs"
)

type InvestmentCalculator struct {
    Categories     []structs.Category
    AmountToInvest float64
}

func NewInvestmentCalculator() *InvestmentCalculator {
    return &InvestmentCalculator{
        Categories:     []structs.Category{},
        AmountToInvest: 0,
    }
}

func (ic *InvestmentCalculator) AddCategory(category structs.Category) {
    ic.Categories = append(ic.Categories, category)
}

func (ic *InvestmentCalculator) GetCategories() []structs.Category {
    return ic.Categories
}

// SetAmountToInvest sets the amount to invest
func (ic *InvestmentCalculator) SetAmountToInvest(amount float64) {
    ic.AmountToInvest = amount
}

func (ic *InvestmentCalculator) CalculateAllocation() []structs.Category {
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
    sum := 0.0
    for i, category := range ic.Categories {
        if category.Locked {
            continue
        }
        adjustedTarget := category.Target / adjustedTotalTargetPercentage
        ic.Categories[i].Investment = roundFloat(adjustedRemainingInvestment * adjustedTarget - category.Current, 2)
        sum += ic.Categories[i].Investment
    }

    // Third pass: Adjust the last category to make the sum of investments equal to the total amount to invest
    ic.Categories[len(ic.Categories)-1].Investment += ic.AmountToInvest - sum

    // set achieved allocation
    for i, category := range ic.Categories {
        ic.Categories[i].AchievedAllocation = ((category.Current + category.Investment) / totalFutureInvestment) * 100
    }

    return ic.Categories
}

func roundFloat(val float64, precision uint) float64 {
    ratio := math.Pow(10, float64(precision))
    return math.Round(val*ratio) / ratio
}
