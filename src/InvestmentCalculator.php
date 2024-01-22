<?php
declare(strict_types=1);

namespace Smic\Rebalancing;

class InvestmentCalculator
{
    private array $categories;
    private int $amountToInvest;

    public function __construct()
    {
        $this->categories = [];
        $this->amountToInvest = 0;
    }

    public function addCategory($name, $currentInvestment, $targetAllocation): void
    {
        $this->categories[$name] = [
            "current" => $currentInvestment,
            "target" => $targetAllocation,
            "locked" => false,
            "investment" => 0
        ];
    }

    public function setAmountToInvest(int $amount): void
    {
        $this->amountToInvest = $amount;
    }

    public function calculateDistribution(): array
    {
        $totalCurrentInvestment = array_sum(array_column($this->categories, 'current'));
        $adjustedTotalTargetPercentage = 0;
        $lockedInvestment = 0;
        $totalFutureInvestment = $totalCurrentInvestment + $this->amountToInvest;

        foreach ($this->categories as $name => $category) {
            $targetInvestment = $totalFutureInvestment * ($category['target'] / 100);
            if ($category['current'] >= $targetInvestment) {
                $this->categories[$name]['locked'] = true;
                $lockedInvestment += $category['current'];
            } else {
                $adjustedTotalTargetPercentage += $category['target'];
                $this->categories[$name]['locked'] = false;
            }
        }

        $adjustedRemainingInvestment = $totalFutureInvestment - $lockedInvestment;

        foreach ($this->categories as $name => $category) {
            if (!$category['locked']) {
                $adjustedTarget = $category['target'] / $adjustedTotalTargetPercentage;
                $this->categories[$name]['investment'] = $adjustedRemainingInvestment * $adjustedTarget - $category['current'];
            } else {
                $this->categories[$name]['investment'] = 0;
            }
        }

        return $this->formatResults();
    }

    private function formatResults(): array
    {
        $results = [];
        $totalFutureInvestment = array_sum(array_column($this->categories, 'current')) + $this->amountToInvest;

        foreach ($this->categories as $name => $category) {
            $totalInvestment = $category['current'] + $category['investment'];
            $achievedAllocation = ($totalInvestment / $totalFutureInvestment) * 100;
            $percentOfInvestment = ($this->amountToInvest > 0) ? ($category['investment'] / $this->amountToInvest) * 100 : 0;

            $results[] = [
                'Category' => $name,
                'Target Allocation (%)' => number_format((float)$category['target'], 2, ','),
                'Amount Before' => number_format((float)$category['current'], 2, ',', '.'),
                'Amount Added' => number_format((float)$category['investment'], 2, ',', '.'),
                'Total Investment After' => number_format((float)$totalInvestment, 2, ',', '.'),
                'Achieved Allocation (%)' => number_format((float)$achievedAllocation, 2, ','),
                'Investment Percentage (%)' => number_format((float)$percentOfInvestment, 2, ','),
                '_totalRaw' => $totalInvestment,
            ];
        }

        return $results;
    }
}
