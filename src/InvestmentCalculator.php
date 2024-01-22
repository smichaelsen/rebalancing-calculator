<?php
declare(strict_types=1);

namespace Smic\Rebalancing;

class InvestmentCalculator
{
    /** @var Category[] */
    private array $categories = [];
    private float $amountToInvest = 0;

    public function addCategory(Category $category): void
    {
        $this->categories[] = $category;
    }

    public function setAmountToInvest(float $amount): void
    {
        $this->amountToInvest = $amount;
    }

    public function calculateDistribution(): array
    {
        $totalCurrentInvestment = array_reduce($this->categories, function (float $carry, Category $category) {
            return $carry + $category->getCurrentInvestment();
        }, .0);
        $adjustedTotalTargetPercentage = 0;
        $lockedInvestment = 0;
        $totalFutureInvestment = $totalCurrentInvestment + $this->amountToInvest;

        foreach ($this->categories as $category) {
            $targetInvestment = $totalFutureInvestment * ($category->getTargetAllocation() / 100);
            if ($category->getCurrentInvestment() >= $targetInvestment) {
                $category->lock();
                $lockedInvestment += $category->getCurrentInvestment();
            } else {
                $adjustedTotalTargetPercentage += $category->getTargetAllocation();
            }
        }

        $adjustedRemainingInvestment = $totalFutureInvestment - $lockedInvestment;

        foreach ($this->categories as $category) {
            if ($category->isLocked()) {
                continue;
            }

            $adjustedTarget = $category->getTargetAllocation() / $adjustedTotalTargetPercentage;
            $category->setAmountAdded($adjustedRemainingInvestment * $adjustedTarget - $category->getCurrentInvestment());
        }

        return $this->formatResults();
    }

    private function formatResults(): array
    {
        $results = [];
        $totalCurrentInvestment = array_reduce($this->categories, function (float $carry, Category $category) {
            return $carry + $category->getCurrentInvestment();
        }, .0);
        $totalFutureInvestment = $totalCurrentInvestment + $this->amountToInvest;

        foreach ($this->categories as $category) {
            $achievedAllocation = ($category->getAmountAfter() / $totalFutureInvestment) * 100;
            $percentOfInvestment = ($this->amountToInvest > 0) ? ($category->getAmountAdded() / $this->amountToInvest) * 100 : 0;

            $results[] = [
                'Category' => $category->getName(),
                'Target Allocation (%)' => number_format($category->getTargetAllocation(), 2, ','),
                'Amount Before' => number_format($category->getCurrentInvestment(), 2, ',', '.'),
                'Amount Added' => number_format($category->getAmountAdded(), 2, ',', '.'),
                'Total Investment After' => number_format($category->getAmountAfter(), 2, ',', '.'),
                'Achieved Allocation (%)' => number_format($achievedAllocation, 2, ','),
                'Investment Percentage (%)' => number_format((float)$percentOfInvestment, 2, ','),
                '_totalRaw' => $category->getAmountAfter(),
            ];
        }

        return $results;
    }
}
