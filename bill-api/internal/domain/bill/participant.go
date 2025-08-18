package bill

import (
	"github.com/google/uuid"
	internalerrors "github.com/stinodes/bill-api/internal/errors"
)

type Participant struct {
	ID           string
	Name         string
	DivisionType *DivisionType
	Amount       float64
}

func NewParticipant(name string) (*Participant, error) {
	if name == "" {
		return nil, internalerrors.ErrBadInput
	}

	return &Participant{
		ID:           uuid.NewString(),
		Name:         name,
		DivisionType: nil,
		Amount:       0,
	}, nil
}

func ParticipantFromDB(id string, name string, divisionType *DivisionType, amount float64) *Participant {
	return &Participant{
		ID:           id,
		Name:         name,
		DivisionType: divisionType,
		Amount:       amount,
	}
}

func (p *Participant) SetType(divisionType *DivisionType) error {
	p.DivisionType = divisionType
	return nil
}

func (p *Participant) SetAmount(amount float64) error {
	p.Amount = amount
	return nil
}
