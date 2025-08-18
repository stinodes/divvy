package bill

import (
	"errors"
	"testing"

	internalerrors "github.com/stinodes/bill-api/internal/errors"
)

func TestCreateBill(t *testing.T) {
	bill, err := NewBill("Test Bill", nil)
	if err != nil {
		t.Errorf("Error creating bill: %v", err)
	}
	if bill.DivisionType != DivisionTypeShare {
		t.Errorf("Expected default division type to be %s, got %s", DivisionTypeShare, bill.DivisionType)
	}
	if bill.Name != "Test Bill" {
		t.Errorf("Bill name does not match argument', got %s", bill.Name)
	}
}

func TestItems(t *testing.T) {
	bill, err := NewBill("Test Bill", nil)

	carolus, err := NewItem("Carolus Classic", 4.00, 2)
	if err != nil {
		t.Errorf("Error creating carolus: %v", err)
	}

	bill.AddItem(carolus)

	_, err = NewItem("", 4.00, 2)
	if err == nil {
		t.Errorf("Item without name should return an error: %v", err)
	}

	_, err = NewItem("Carolus Classic", -1, 2)
	if err == nil {
		t.Errorf("Item with negative price should return an error: %v", err)
	}

	bill.IncrementItem(carolus.ID)
	if carolus.Quantity != 3 {
		t.Errorf("Item quantity should be 3, got %d", carolus.Quantity)
	}

	err = bill.DecrementItem(carolus.ID)
	if err != nil {
		t.Errorf("Error decrementing item: %v", err)
	}
	if carolus.Quantity != 2 {
		t.Errorf("Item quantity should be 2, got %d", carolus.Quantity)
	}

	err = bill.DecrementItem(carolus.ID)
	if err != nil {
		t.Errorf("Error decrementing item: %v", err)
	}

	err = bill.DecrementItem(carolus.ID)
	if err != nil {
		t.Errorf("Error decrementing item: %v", err)
	}
	if carolus.Quantity != 0 {
		t.Errorf("Item quantity should be 0, got %d", carolus.Quantity)
	}

	err = bill.RemoveItem("fake-id")
	if !errors.Is(err, internalerrors.ErrNotFound) {
		t.Errorf("Error deleting non-existent item should return ErrNotFound, got %v", err)
	}
}

func TestParticipants(t *testing.T) {
	bill, err := NewBill("Test Bill", nil)

	anduin, err := NewParticipant("Anduin")
	if err != nil {
		t.Errorf("Error creating anduin: %v", err)
	}
	err = bill.AddParticipant(anduin)
	if err != nil {
		t.Errorf("Error adding anduin to bill: %v", err)
	}
	if anduin.Amount != 1 {
		t.Errorf("Anduin amount should be set to 1 by default, got %f", anduin.Amount)
	}
	if anduin.DivisionType != nil {
		t.Errorf("Anduin type should be nil by default, got %s", *anduin.DivisionType)
	}

	_, err = NewParticipant("")
	if err == nil {
		t.Errorf("Participant without name should return an error: %v", err)
	}

	err = bill.RemoveParticipant(anduin.ID)
	if err != nil {
		t.Errorf("Error removing anduin from bill: %v", err)
	}

	err = bill.RemoveParticipant("fake-id")
	if !errors.Is(err, internalerrors.ErrNotFound) {
		t.Errorf("Error removing non-existent participant should return ErrNotFound, got %v", err)
	}
}

func TestBillCalculations(t *testing.T) {
	bill, err := NewBill("Test Bill", nil)

	carolus, err := NewItem("Carolus Classic", 4.00, 2)
	chouffe, err := NewItem("Chouffe", 3.60, 3)
	nachos, err := NewItem("Nachos", 4.50, 1)

	bill.AddItem(carolus)
	bill.AddItem(chouffe)
	bill.AddItem(nachos)

	anduin, err := NewParticipant("Anduin")
	wrynn, err := NewParticipant("Wrynn")
	thrall, err := NewParticipant("Thrall")
	lothar, err := NewParticipant("Lothar")

	bill.AddParticipant(anduin)
	bill.AddParticipant(wrynn)
	bill.AddParticipant(thrall)
	bill.AddParticipant(lothar)

}
