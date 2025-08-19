package bill

import (
	"github.com/google/uuid"
	internalerrors "github.com/stinodes/bill-api/internal/errors"
)

type Participant struct {
	ID     string
	Name   string
	Amount float64
	Shares int
}

func NewParticipant(name string) (*Participant, error) {
	if name == "" {
		return nil, internalerrors.ErrBadInput
	}

	return &Participant{
		ID:     uuid.NewString(),
		Name:   name,
		Amount: 0,
		Shares: 0,
	}, nil
}

func ParticipantFromDB(id string, name string, amount float64, shares int) *Participant {
	return &Participant{
		ID:     id,
		Name:   name,
		Amount: amount,
		Shares: shares,
	}
}

func (p *Participant) Edit(amount float64, shares int) error {
	if amount < 0 || shares < 0 {
		return internalerrors.ErrBadInput
	}
	p.Amount = amount
	p.Shares = shares
	return nil
}
