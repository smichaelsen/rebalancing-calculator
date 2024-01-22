#!/usr/bin/env php
<?php
require __DIR__.'/vendor/autoload.php';

use Smic\Rebalancing\Category;
use Smic\Rebalancing\InvestmentCalculator;
use Symfony\Component\Console\Helper\QuestionHelper;
use Symfony\Component\Console\Input\InputInterface;
use Symfony\Component\Console\Output\OutputInterface;
use Symfony\Component\Console\Command\Command;
use Symfony\Component\Console\Question\Question;
use Symfony\Component\Console\Helper\Table;

class InvestmentCalculatorCommand extends Command
{
    protected static $defaultName = 'app:rebalance';

    protected function execute(InputInterface $input, OutputInterface $output): int
    {
        /** @var QuestionHelper $helper */
        $helper = $this->getHelper('question');
        $calculator = new InvestmentCalculator();

        while (true) {
            $categoryQuestion = new Question('Enter category name (leave empty to finish): ');
            $categoryName = $helper->ask($input, $output, $categoryQuestion);
            if (empty($categoryName)) {
                break;
            }

            $currentInvestmentQuestion = new Question('Enter amount already invested in ' . $categoryName . ': ');
            $currentInvestment = str_replace(',', '.', $helper->ask($input, $output, $currentInvestmentQuestion));

            $targetAllocationQuestion = new Question('Enter target allocation in % for ' . $categoryName . ': ');
            $targetAllocation = str_replace(',', '.', $helper->ask($input, $output, $targetAllocationQuestion));

            $category = new Category($categoryName, (float)$currentInvestment, (float)$targetAllocation);
            $calculator->addCategory($category);
        }

        $amountToInvestQuestion = new Question('Enter amount to invest: ');
        $amountToInvest = str_replace(',', '.', $helper->ask($input, $output, $amountToInvestQuestion));
        $calculator->setAmountToInvest((float)$amountToInvest);

        $results = $calculator->calculateDistribution();

        // remove all keys starting with _
        $formattedResults = array_map(function ($result) {
            return array_filter($result, function ($key) {
                return !str_starts_with($key, '_');
            }, ARRAY_FILTER_USE_KEY);
        }, $results);

        $table = new Table($output);
        $table
            ->setHeaders(array_keys($formattedResults[0]))
            ->setRows($formattedResults)
            ->setFooterTitle('Total invested: ' . number_format((float)array_sum(array_column($results, '_totalRaw')), 2, ',', '.'));

        $table->render();

        return Command::SUCCESS;
    }
}

$app = new Symfony\Component\Console\Application();
$app->add(new InvestmentCalculatorCommand());
$app->run();
