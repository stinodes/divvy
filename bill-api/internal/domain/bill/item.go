package bill

import (
	"github.com/google/uuid"
	internalerrors "github.com/stinodes/bill-api/internal/errors"
)

type Item struct {
	ID       string
	Name     string
	Price    float64
	Quantity int
}

func NewItem(name string, price float64, quantity int) (*Item, error) {
	if name == "" {
		return nil, internalerrors.ErrBadInput
	}
	if price <= 0 {
		return nil, internalerrors.ErrBadInput
	}
	return &Item{
		ID:       uuid.NewString(),
		Name:     name,
		Price:    price,
		Quantity: quantity,
	}, nil
}

func ItemFromDB(id string, name string, price float64, quantity int) *Item {
	return &Item{
		ID:       id,
		Name:     name,
		Price:    price,
		Quantity: quantity,
	}
}

func (i *Item) Increment() error {
	i.Quantity++
	return nil
}
func (i *Item) Decrement() error {
	i.Quantity = max(i.Quantity-1, 0)
	return nil
}
