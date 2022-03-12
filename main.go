package main

import (
	"fmt"
	"vending_machine_go/model"
)

func buttonMap() map[string]string {
	return map[string]string {
		"a1": "Pepsi",
		"b1": "Pop Tarts",
		"c1": "Lays",
	}
}

func main() {
	pepsi := model.NewItem("Pepsi", 50, 10)
	popTart := model.NewItem("Pop Tart", 45, 10)
	lays := model.NewItem("Lays", 35, 10)

	items := map[string]*model.Item{
		"Pepsi": pepsi,
		"Pop Tart": popTart,
		"Lays": lays,
	}

	inventory := model.NewInventory(items)

	machine := model.NewVendingMachine(inventory, buttonMap())

	dispensedItem, err := machine.Vend("a1")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(dispensedItem)
}