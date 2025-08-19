package bill

import (
	"slices"

	"github.com/google/uuid"
	internalerrors "github.com/stinodes/bill-api/internal/errors"
)

type Bill struct {
	ID   string
	Name string

	Items        []*Item
	Participants []*Participant
}

func NewBill(name string) (*Bill, error) {
	if name == "" {
		return nil, internalerrors.ErrBadInput
	}

	return &Bill{
		ID:           uuid.NewString(),
		Name:         name,
		Items:        []*Item{},
		Participants: []*Participant{},
	}, nil
}

func BillFromDB(id string, name string, items []*Item, participants []*Participant) *Bill {
	return &Bill{
		ID:           id,
		Name:         name,
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

func (b *Bill) GetFlatTotal() float64 {
	total := 0.0
	for _, participant := range b.Participants {
		total += participant.Amount
	}
	return total
}

func (b *Bill) GetTotalShares() int {
	total := 0
	for _, participant := range b.Participants {
		total += participant.Shares
	}
	return total
}

func (b *Bill) GetMoneyPerShare() float64 {
	totalShareMoney := b.GetTotal() - b.GetFlatTotal()
	shares := b.GetTotalShares()
	return totalShareMoney / float64(shares)
}

func (b *Bill) GetMoneyForParticipant(id string) float64 {
	participant := b.GetParticipant(id)
	mPerShare := b.GetMoneyPerShare()
	shareAmount := float64(participant.Shares) * mPerShare
	return participant.Amount + shareAmount
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

func (b *Bill) AddParticipant(participant *Participant) error {
	if b.GetParticipant(participant.ID) != nil {
		return internalerrors.ErrDuplicate
	}
	b.Participants = append(b.Participants, participant)

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

func (b *Bill) EditParticipant(id string, amount float64, shares int) error {
	participant := b.GetParticipant(id)
	if participant == nil {
		return internalerrors.ErrNotFound
	}
	err := participant.Edit(amount, shares)
	return err

}

func (b *Bill) AddItem(item *Item) error {
	if b.GetItem(item.ID) != nil {
		return internalerrors.ErrDuplicate
	}
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
