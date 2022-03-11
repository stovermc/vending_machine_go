package model

type Inventory interface{
	Item(name string) Item
	RemoveItem(name string)
}

type inventory struct {
	items map[string]Item	
}

func NewInventory(i map[string]Item) Inventory {
	return &inventory{
		items: i,
	}
}

func (i *inventory) Item(name string) Item {
	return i.items[name]
}

func (i *inventory) RemoveItem(name string)  {
	item := i.items[name]

	item.Subtract()

	return
}