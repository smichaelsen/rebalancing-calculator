<?php

declare(strict_types=1);

namespace Smic\Rebalancing;

class Category
{
    private readonly string $name;
    private readonly float $currentInvestment;
    private readonly float $targetAllocation;

    private float $amountAdded = .0;
    private bool $locked = false;

    public function __construct(string $name, float $currentInvestment, float $targetAllocation)
    {
        $this->name = $name;
        $this->currentInvestment = $currentInvestment;
        $this->targetAllocation = $targetAllocation;
    }

    public function getName(): string
    {
        return $this->name;
    }

    public function getCurrentInvestment(): float
    {
        return $this->currentInvestment;
    }

    public function getTargetAllocation(): float
    {
        return $this->targetAllocation;
    }

    public function getAmountAdded(): float
    {
        return $this->amountAdded;
    }

    public function setAmountAdded(float $amountAdded): void
    {
        $this->amountAdded = $amountAdded;
    }

    public function getAmountAfter(): float
    {
        return $this->currentInvestment + $this->amountAdded;
    }

    public function lock(): void
    {
        $this->locked = true;
    }

    public function isLocked(): bool
    {
        return $this->locked;
    }
}