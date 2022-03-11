package model_test

import (
	"fmt"
	"testing"
	"vending_machine_go/model"
)

func TestVendItem(t *testing.T) {
	want := "Pepsi"

	item := model.NewItem("Pepsi", 50, 5)
	items := map[string]model.Item{
		"Pepsi": item,
	}
	inventory := model.NewInventory(items)
	machine := model.NewVendingMachine(inventory, buttonMap())

	got, _, err := machine.Vend("a1", 50)

	if err != nil {
		t.Fatalf("expected err to be nil. err: %s", err)
	}

	if got != want {
		t.Errorf("expected: %s\n got: %s", want, got)
	}
}

func TestInventoryIsDecrementedAfterVendingAnItem(t *testing.T) {
	want := 4
	item := model.NewItem("Pepsi", 50, 5)

	items := map[string]model.Item{
		"Pepsi": item,
	}
	inventory := model.NewInventory(items)

	machine := model.NewVendingMachine(inventory, buttonMap())
	_, _, err := machine.Vend("a1", 50)

	if err != nil {
		t.Fatalf("expected err to be nil. err: %s", err)
	}

	got := machine.Inventory().Item("Pepsi").Quantity()

	if got != want {
		t.Errorf("expected: %d\n got: %d", want, got)
	}
}

func TestWhenItemIsOutOfStock(t *testing.T) {
	want := fmt.Errorf("Out of Stock")

	item := model.NewItem("Pepsi", 50, 0)

	items := map[string]model.Item{
		"Pepsi": item,
	}
	inventory := model.NewInventory(items)
	machine := model.NewVendingMachine(inventory, buttonMap())

	_, _, got := machine.Vend("a1", 50)

	if got.Error() != want.Error() {
		t.Errorf("expected: %s\n got: %s", want, got)
	}
}

func TestWhenPaymentIsntEnough(t *testing.T) {
	want := fmt.Errorf("Not enough payment")

	item := model.NewItem("Pepsi", 50, 5)

	items := map[string]model.Item{
		"Pepsi": item,
	}
	inventory := model.NewInventory(items)

	machine := model.NewVendingMachine(inventory, buttonMap())

	_, _, got := machine.Vend("a1", 0)

	if got.Error() != want.Error() {
		t.Fatalf("expected: %s\n got: %s", want, got)
	}

}

func TestCorrectChangeIsReturned(t *testing.T) {
	want := 50

	item := model.NewItem("Pepsi", 50, 5)
	items := map[string]model.Item{
		"Pepsi": item,
	}
	inventory := model.NewInventory(items)
	machine := model.NewVendingMachine(inventory, buttonMap())

	_, got, err := machine.Vend("a1", 100)

	if err != nil {
		t.Fatalf("expected err to be nil. err: %s", err)
	}

	if got != want {
		t.Errorf("expected: %d\n got: %d", want, got)
	}
}


func buttonMap() map[string]string {
	return map[string]string{
		"a1": "Pepsi",
		"a2": "Lays",
	}
}
