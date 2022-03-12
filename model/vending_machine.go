package model

import "fmt"

const (
	NICKEL = iota
	DIME
	QUARTER
	DOLLAR
)

type Coin struct {
	key   int
	value int
	label string
}

type VendingMachine struct {
	inventory     *Inventory
	buttonMap     map[string]string
	coins         map[int]*Coin
	coinsInserted []*Coin
	bank          map[*Coin]int
}

func NewVendingMachine(i *Inventory, bm map[string]string) *VendingMachine {
	coins := make(map[int]*Coin)
	coins[NICKEL] = &Coin{NICKEL, 5, "N"}
	coins[DIME] = &Coin{DIME, 10, "D"}
	coins[QUARTER] = &Coin{QUARTER, 25, "Q"}
	coins[DOLLAR] = &Coin{DOLLAR, 100, "DD"}

	bank := make(map[*Coin]int)
	bank[coins[NICKEL]] = 0
	bank[coins[DIME]] = 0
	bank[coins[QUARTER]] = 0
	bank[coins[DOLLAR]] = 0

	return &VendingMachine{
		inventory:     i,
		buttonMap:     bm,
		coins:         coins,
		coinsInserted: []*Coin{},
		bank:          bank,
	}
}

func (v *VendingMachine) Service() {
	v.bank[v.coins[NICKEL]] = 50
	v.bank[v.coins[DIME]] = 50
	v.bank[v.coins[QUARTER]] = 50
	v.bank[v.coins[DOLLAR]] = 50
}

func (v *VendingMachine) Vend(id string) (string, error) {
	name := v.buttonMap[id]

	item := v.Inventory().Item(name)

	if item.Quantity() == 0 {
		return "", fmt.Errorf("Out of Stock")
	}

	if v.AmountInserted() < item.Price() {
		return "", fmt.Errorf("Not enough payment")
	}

	change, err := v.makeChange(v.AmountInserted(), item.Price())
	if err != nil {
		return "", fmt.Errorf("Unable to make change. Please pay with exact change.")
	}

	item.Subtract()

	return item.Name() + change, nil
}

func (v *VendingMachine) Inventory() *Inventory {
	return v.inventory
}

func (v *VendingMachine) Insert(coin int) {
	v.coinsInserted = append(v.coinsInserted, v.coins[coin])
}

func (v *VendingMachine) AmountInserted() int {
	amountInserted := 0
	for _, coin := range v.coinsInserted {
		amountInserted += coin.value
	}

	return amountInserted
}

func (v *VendingMachine) CointReturn() []string {
	coinsReturned := make([]string, len(v.coinsInserted))

	for i, coin := range v.coinsInserted {
		coinsReturned[i] = coin.label
	}

	return coinsReturned
}

func (v *VendingMachine) makeChange(paid, price int) (string, error) {
	coinReturn := ""
	amountDue := paid - price
fmt.Printf("amountDue: %v\n", amountDue)

CoinLoop:
	for i := QUARTER; i >= NICKEL; i-- {
		coins := amountDue / v.coins[i].value

		if v.bank[v.coins[i]]-coins < 0 {
			continue CoinLoop
		}

		v.bank[v.coins[i]] -= coins

		for j := 1; j <= coins; j++ {
			coinReturn += ", " + v.coins[i].label
		}
		amountDue -= v.coins[i].value * coins
 
		if amountDue == 0 {
			break CoinLoop
		}
	}

	if amountDue != 0 {
		return "", fmt.Errorf("Unable to make change")
	}

	return coinReturn, nil
}
