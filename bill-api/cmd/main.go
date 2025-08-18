package main

import "github.com/stinodes/bill-api/internal/domain/bill"

func main() {
	// TODO: Implement
	bill, err := bill.NewBill("Test Bill", nil)
	if err != nil {
		panic(err)
	}
	println(bill.Name)
}
