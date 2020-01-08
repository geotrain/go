package main

import (
	"fmt"
)

//Trade is a trade in the stocks
type Trades struct {
	Symbol string  // Stock symbol
	Volume int     // Number of shares
	Price  float64 // Trade price
	Buy    bool    // True if buy trade, false if sell trade
}

// Value returns the trade value
func (t *Trades) Value() float64 {
	value := float64(t.Volume) * t.Price
	if t.Buy {
		value = -value
	}
	return value
}

func main() {
	t := Trades{
		Symbol: "Microsoft",
		Volume: 10,
		Price:  99.98,
		Buy:    true,
	}
	fmt.Println(t.Value())
}
