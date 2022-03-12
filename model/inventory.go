package model

type Inventory struct {
	items map[string]*Item
}

func NewInventory(i map[string]*Item) *Inventory {
	return &Inventory{
		items: i,
	}
}

func (i *Inventory) Items() []*Item {
	items := make([]*Item, len(i.items))
	idx := 0
	for _, item := range i.items {
		items[idx] = item

		idx++
	}

	return items
}

func (i *Inventory) Item(name string) *Item {
	return i.items[name]
}

func (i *Inventory) RemoveItem(name string) {
	item := i.items[name]

	item.Subtract()

	return
}
