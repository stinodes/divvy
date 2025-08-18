package bill

import (
	"math"
	"slices"

	"github.com/google/uuid"
	internalerrors "github.com/stinodes/bill-api/internal/errors"
)

type DivisionType string

const (
	DivisionTypeFlat  DivisionType = "flat"
	DivisionTypeShare DivisionType = "share"
)

type Bill struct {
	ID           string
	Name         string
	DivisionType DivisionType

	Items        []*Item
	Participants []*Participant
}

func NewBill(name string, divisionType *DivisionType) (*Bill, error) {
	if name == "" {
		return nil, internalerrors.ErrBadInput
	}
	divType := DivisionTypeShare
	if divisionType != nil {
		divType = *divisionType
	}

	return &Bill{
		ID:           uuid.NewString(),
		Name:         name,
		DivisionType: divType,
		Items:        []*Item{},
		Participants: []*Participant{},
	}, nil
}

func BillFromDB(id string, name string, divisionType DivisionType, items []*Item, participants []*Participant) *Bill {
	return &Bill{
		ID:           id,
		Name:         name,
		DivisionType: divisionType,
		Items:        items,
		Participants: participants,
	}
}

func (b *Bill) GetTotal() float64 {
	total := 0.0
	for _, item := range b.Items {
		total += float64(item.Quantity) * item.Price
	}
	return total
}

func (b *Bill) GetTotalShares() int {
	if b.DivisionType != DivisionTypeShare {
		return 0
	}
	total := 0
	for _, participant := range b.Participants {
		total += int(math.Floor(participant.Amount))
	}
	return total
}

func (b *Bill) GetMoneyForShare(id string) float64 {
	totalMoney := b.GetTotal()
	shares := b.GetTotalShares()
	return totalMoney / float64(shares)
}

func (b *Bill) GetMoneyForParticipant(id string) float64 {
	participant := b.GetParticipant(id)
	if b.DivisionType == DivisionTypeFlat {
		return participant.Amount
	}
	mPerShare := b.GetMoneyForShare(id)
	return participant.Amount * mPerShare
}

func (b *Bill) GetItem(id string) *Item {
	index := slices.IndexFunc(b.Items, func(i *Item) bool {
		return i.ID == id
	})
	if index == -1 {
		return nil
	}
	return b.Items[index]
}

func (b *Bill) GetParticipant(id string) *Participant {
	index := slices.IndexFunc(b.Participants, func(i *Participant) bool {
		return i.ID == id
	})
	if index == -1 {
		return nil
	}
	return b.Participants[index]
}

func (b *Bill) SetParticipantAmount(id string, amount float64) error {

}

func (b *Bill) AddParticipant(participant *Participant) error {
	b.Participants = append(b.Participants, participant)
	if b.DivisionType == DivisionTypeShare {
		participant.Amount = 1
	}
	return nil
}

func (b *Bill) RemoveParticipant(id string) error {
	for i, p := range b.Participants {
		if p.ID == id {
			b.Participants = slices.Delete(b.Participants, i, i+1)
			return nil
		}
	}
	return internalerrors.ErrNotFound
}

func (b *Bill) AddItem(item *Item) error {
	b.Items = append(b.Items, item)
	return nil
}
func (b *Bill) RemoveItem(id string) error {
	for i, p := range b.Items {
		if p.ID == id {
			b.Items = slices.Delete(b.Items, i, i+1)
			return nil
		}
	}
	return internalerrors.ErrNotFound
}

func (b *Bill) IncrementItem(id string) error {
	item := b.GetItem(id)
	if item == nil {
		return internalerrors.ErrNotFound
	}
	err := item.Increment()
	if err != nil {
		return err
	}
	return nil
}
func (b *Bill) DecrementItem(id string) error {
	item := b.GetItem(id)
	if item == nil {
		return internalerrors.ErrNotFound
	}
	err := item.Decrement()
	if err != nil {
		return err
	}
	if item.Quantity == 0 {
		err := b.RemoveItem(id)
		if err != nil {
			return err
		}
	}
	return nil
}
