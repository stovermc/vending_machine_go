package model

import "fmt"

type VendingMachine interface {
	Vend(id string, payment int) (string, int, error)
	Inventory() Inventory
}

type vendingMachine struct {
  inventory Inventory
	buttonMap map[string]string
}

func NewVendingMachine(i Inventory, bm map[string]string) VendingMachine{
  return &vendingMachine{
		inventory: i,
		buttonMap: bm,
	}
}

func (v *vendingMachine) Inventory() Inventory {
	return v.inventory
}

func (v *vendingMachine) Vend(id string, payment int) (string, int, error) {
	name := v.buttonMap[id];

  item := v.Inventory().Item(name)

	if item.Quantity() == 0 {
		return "", 0, fmt.Errorf("Out of Stock")
	}

	if payment < item.Price(){
		return "", 0, fmt.Errorf("Not enough payment")
	}

	change := payment - item.Price()

	item.Subtract()

	return item.Name(), change, nil
}