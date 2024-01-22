<?php

declare(strict_types=1);

use PHPUnit\Framework\TestCase;
use Smic\Rebalancing\Category;
use Smic\Rebalancing\InvestmentCalculator;

final class InvestmentCalculatorTest extends TestCase
{
    private InvestmentCalculator $subject;

    protected function setUp(): void
    {
        $this->subject = new InvestmentCalculator();
    }

    /**
     * @test
     */
    public function testCategoryNamesAreReturnedAsProvided(): void
    {
        $this->subject->addCategory(new Category('risk_free', 10000, 45));
        $this->subject->addCategory(new Category('stable_return', 5000, 50));
        $this->subject->addCategory(new Category('high_risk', 0, 5));

        $results = $this->subject->calculateDistribution();

        $this->assertEquals('risk_free', $results[0]['Category']);
        $this->assertEquals('stable_return', $results[1]['Category']);
        $this->assertEquals('high_risk', $results[2]['Category']);
    }

    /**
     * @test
     */
    public function testTargetAllocationsAreReturnedAsProvided(): void
    {
        $this->subject->addCategory(new Category('risk_free', 10000, 45));
        $this->subject->addCategory(new Category('stable_return', 5000, 50));
        $this->subject->addCategory(new Category('high_risk', 0, 5));

        $results = $this->subject->calculateDistribution();

        $this->assertEquals('45,00', $results[0]['Target Allocation (%)']);
        $this->assertEquals('50,00', $results[1]['Target Allocation (%)']);
        $this->assertEquals('5,00', $results[2]['Target Allocation (%)']);
    }

    /**
     * @test
     */
    public function testAmountsBeforeAreReturnedAsProvided(): void
    {
        $this->subject->addCategory(new Category('risk_free', 10000, 45));
        $this->subject->addCategory(new Category('stable_return', 5000, 50));
        $this->subject->addCategory(new Category('high_risk', 0, 5));

        $results = $this->subject->calculateDistribution();

        $this->assertEquals('10.000,00', $results[0]['Amount Before']);
        $this->assertEquals('5.000,00', $results[1]['Amount Before']);
        $this->assertEquals('0,00', $results[2]['Amount Before']);
    }

    /**
     * @test
     */
    public function testAmountAdded(): void
    {
        $this->subject->addCategory(new Category('risk_free', 10000, 45));
        $this->subject->addCategory(new Category('stable_return', 5000, 50));
        $this->subject->addCategory(new Category('high_risk', 0, 5));
        $this->subject->setAmountToInvest(5000);

        $results = $this->subject->calculateDistribution();

        $this->assertEquals('0,00', $results[0]['Amount Added']);
        $this->assertEquals('4.090,91', $results[1]['Amount Added']);
        $this->assertEquals('909,09', $results[2]['Amount Added']);
    }
}
