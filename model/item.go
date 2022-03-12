package model

type Item struct {
	name string
	price int
	quantity int
}

func NewItem(n string, p, q int) *Item {
  return &Item{
		name: n,
		price: p,
		quantity: q,
	}
}

func (i *Item) Name() string {
	return i.name
}

func (i *Item) Price() int {
	return i.price
}

func (i *Item) Quantity() int {
	return i.quantity
}

func (i *Item) Subtract() {
	i.quantity = i.quantity - 1
	return
}