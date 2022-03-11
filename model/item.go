package model

type Item interface{
	Name() string
	Price() int
	Quantity() int
	Subtract()
}

type item struct {
	name string
	price int
	quantity int
}

func NewItem(n string, p, q int) Item {
  return &item{
		name: n,
		price: p,
		quantity: q,
	}
}

func (i *item) Name() string {
	return i.name
}

func (i *item) Price() int {
	return i.price
}

func (i *item) Quantity() int {
	return i.quantity
}

func (i *item) Subtract() {
	i.quantity = i.quantity - 1
	return
}