# Rebalancing calculator

Where to put your investments to reach your desired asset allocation?

## Intended use case

This calculator is intended to help you allocate investments according to your desired asset allocation. It assumes you do the rebalancing from time to time, e.g. once per year.
After a year of investment the allocation in your portfolio will differ from your desired allocation. That is because the assets will have performed differently. Also you might have changed your opinion about the desired allocation.

This calculator will never advise to **remove** money from a category - because selling assets often results in transaction costs and taxes. Instead it will advise you how to **add** money in the future to achieve or approach your desired allocation.

## Usage

### Interactive

```bash
rebalance
```

The program will ask for your inputs interactively.

### With command line arguments for categories

```bash
rebalance --category "Developed Marktes;80" --category "Europe;10" --category "Emerging Markets;10"
```

The program will only ask for your holdings and the amount to invest.

## Remarks

* For the "amount to invest" enter the amount of money you will presumably invest until the next rebalancing (for example for one year).
* You can also use the calculator before _every_ investment to determine how to allocate the money - that's how I do it.
* You can enter all amounts with , or . as decimal separator. Don't input any thousands separators.
* You have to make sure your target allocation adds up to 100% - the calculator will not check this.
* If you get an unexpected opportunity to invest significantly more money (lucky you!) just do another rebalancing to decide how to allocate the additional money.
* It should be obvious, but given that any software can and probably will contain bugs: Use this calculator at your own risk. I am not responsible for any losses you might incur by using this calculator.

![Example](./example.png)

If you find this helpful, consider allocating a few bucks to [Unicef](https://www.unicef.org/).
