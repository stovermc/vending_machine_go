package model_test

import (
	"fmt"
	"testing"
	"vending_machine_go/model"
)

func TestVendItem(t *testing.T) {
	want := "Pepsi"

	items := inventoryItemsInStock()
	inventory := model.NewInventory(items)
	machine := model.NewVendingMachine(inventory, buttonMap())

	got, err := machine.Vend("a1")

	if err != nil {
		t.Fatalf("expected err to be nil. err: %s", err)
	}

	if got != want {
		t.Errorf("expected: %s\n got: %s", want, got)
	}
}

func TestInventoryIsDecrementedAfterVendingAnItem(t *testing.T) {
	want := 4
	items := inventoryItemsInStock()
	inventory := model.NewInventory(items)

	machine := model.NewVendingMachine(inventory, buttonMap())
	_, err := machine.Vend("a1")

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

	items := inventoryItemsOutOfStock()
	inventory := model.NewInventory(items)
	machine := model.NewVendingMachine(inventory, buttonMap())

	_, got := machine.Vend("a1")

	if got.Error() != want.Error() {
		t.Errorf("expected: %s\n got: %s", want, got)
	}
}

func TestListItems(t *testing.T) {
	pepsi := model.NewItem("Pepsi", 50, 5)

	items := inventoryItemsInStock()
	inventory := model.NewInventory(items)

	machine := model.NewVendingMachine(inventory, buttonMap())

	list := machine.Inventory().Items()

	for _, item := range list {
		if item.Name() != pepsi.Name() {
			t.Errorf("expected: %s\n got: %s", pepsi.Name(), item.Name())
		}
		if item.Price() != pepsi.Price() {
			t.Errorf("expected: %s\n got: %s", pepsi.Name(), item.Name())
		}
		if item.Quantity() != pepsi.Quantity() {
			t.Errorf("expected: %s\n got: %s", pepsi.Name(), item.Name())
		}
	}
}

func TestInsertingMoney(t *testing.T) {
	items := inventoryItemsInStock()
	inventory := model.NewInventory(items)

	machine := model.NewVendingMachine(inventory, buttonMap())
	machine.Service()

	machine.Insert(model.DOLLAR)
	got := machine.AmountInserted()

	want := 100
	if got != want {
		t.Errorf("expected: %d\n got: %d", want, got)
	}
}

func TestCoinReturn(t *testing.T) {
	items := inventoryItemsInStock()
	inventory := model.NewInventory(items)

	machine := model.NewVendingMachine(inventory, buttonMap())
	machine.Service()

	machine.Insert(model.NICKEL)
	machine.Insert(model.QUARTER)
	machine.Insert(model.DIME)

	got := machine.CointReturn()
	labels := []string{"N", "Q", "D"}

	for i, want := range labels {
		if got[i] != want {
			t.Errorf("expected: %s\n got: %s", want, got)
		}
	}

}

func TestWhenPaymentIsntEnough(t *testing.T) {
	want := fmt.Errorf("Not enough payment")

	items := inventoryItemsInStock()
	inventory := model.NewInventory(items)

	machine := model.NewVendingMachine(inventory, buttonMap())

	_, got := machine.Vend("a1")

	if got.Error() != want.Error() {
		t.Fatalf("expected: %s\n got: %s", want, got)
	}

}

func TestCorrectChangeIsReturned(t *testing.T) {
	want := "Pepsi, Q, Q"

	items := inventoryItemsInStock()
	inventory := model.NewInventory(items)
	machine := model.NewVendingMachine(inventory, buttonMap())

	machine.Service()
	machine.Insert(model.DOLLAR)
	got, err := machine.Vend("a1")

	if err != nil {
		t.Fatalf("expected err to be nil. err: %s", err)
	}

	if got != want {
		t.Errorf("expected: %s\n got: %s", want, got)
	}
}

func inventoryItemsInStock() map[string]*model.Item {
	item := model.NewItem("Pepsi", 50, 5)
	return map[string]*model.Item{
		"Pepsi": item,
	}
}

func inventoryItemsOutOfStock() map[string]*model.Item {
	item := model.NewItem("Pepsi", 50, 0)
	return map[string]*model.Item{
		"Pepsi": item,
	}
}

func buttonMap() map[string]string {
	return map[string]string{
		"a1": "Pepsi",
		"a2": "Lays",
	}
}
