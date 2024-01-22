package calculator

import (
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

// AddCategory adds a new category to the calculator
func (ic *InvestmentCalculator) AddCategory(category structs.Category) {
	ic.Categories = append(ic.Categories, category)
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
