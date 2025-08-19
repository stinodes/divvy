package bill

import (
	"errors"
	"testing"

	"github.com/charmbracelet/log"
	internalerrors "github.com/stinodes/bill-api/internal/errors"
)

func TestCreateBill(t *testing.T) {
	bill, err := NewBill("Test Bill")
	if err != nil {
		t.Errorf("Error creating bill: %v", err)
	}
	if bill.Name != "Test Bill" {
		t.Errorf("Bill name does not match argument', got %s", bill.Name)
	}
}

func TestItems(t *testing.T) {
	bill, err := NewBill("Test Bill")

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
	bill, err := NewBill("Test Bill")

	anduin, err := NewParticipant("Anduin")
	if err != nil {
		t.Errorf("Error creating anduin: %v", err)
	}
	err = bill.AddParticipant(anduin)
	if err != nil {
		t.Errorf("Error adding anduin to bill: %v", err)
	}
	if anduin.Amount != 0 {
		t.Errorf("Anduin amount should be set to 1 by default, got %f", anduin.Amount)
	}
	if anduin.Shares != 0 {
		t.Errorf("Anduin shares should be set to 1 by default, got %f", anduin.Amount)
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
	bill, err := NewBill("Test Bill")
	if err != nil {
		t.Errorf("Error creating bill: %v", err)
	}

	billTotal, err := NewItem("Bill Total", 48.40, 1)
	bill.AddItem(billTotal)

	if bill.GetTotal() != 48.40 {
		t.Errorf("Bill total should be same as item, got %f", bill.GetTotal())
	}

	anduin, err := NewParticipant("Anduin")
	wrynn, err := NewParticipant("Wrynn")
	thrall, err := NewParticipant("Thrall")
	lothar, err := NewParticipant("Lothar")

	bill.AddParticipant(anduin)
	bill.AddParticipant(wrynn)
	bill.AddParticipant(thrall)
	bill.AddParticipant(lothar)

	bill.EditParticipant(anduin.ID, 2.80, 0)
	bill.EditParticipant(wrynn.ID, 0, 3)
	bill.EditParticipant(thrall.ID, 0, 2)
	bill.EditParticipant(lothar.ID, 0, 4)

	if bill.GetMoneyForParticipant(anduin.ID) != 2.80 {
		t.Errorf("Anduin money should be 2.80, got %f", bill.GetMoneyForParticipant(anduin.ID))
	}

	if bill.GetMoneyForParticipant(wrynn.ID) != (bill.GetTotal()-2.80)/9*3 {
		t.Errorf("Wrynn money not correct, got %f", bill.GetMoneyForParticipant(wrynn.ID))
	}
	if bill.GetMoneyForParticipant(thrall.ID) != (bill.GetTotal()-2.80)/9*2 {
		t.Errorf("Thrall money not correct, got %f", bill.GetMoneyForParticipant(thrall.ID))
	}
	if bill.GetMoneyForParticipant(lothar.ID) != (bill.GetTotal()-2.80)/9*4 {
		t.Errorf("Lothar money not correct, got %f", bill.GetMoneyForParticipant(lothar.ID))
	}
}
