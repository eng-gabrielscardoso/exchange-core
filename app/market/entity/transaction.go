package entity

import (
	"time"

	"github.com/google/uuid"
)

// STRUCT DECLARATIONS

type Transaction struct {
	id            string
	buying_order  *Order
	selling_order *Order
	shares        int
	price         float64
	total_amount  float64
	datetime      time.Time
}

// CONSTRUCTOR

func NewTransaction(buyingOrder *Order, sellingOrder *Order, shares int, price float64) *Transaction {
	total := float64(shares) * price

	return &Transaction{
		id:            uuid.New().String(),
		buying_order:  buyingOrder,
		selling_order: sellingOrder,
		shares:        shares,
		price:         price,
		total_amount:  total,
		datetime:      time.Now(),
	}
}

// GETTERS AND SETTERS

func (transaction *Transaction) GetBuyingOrder() *Order {
	return transaction.buying_order
}

func (transaction *Transaction) GetSellingOrder() *Order {
	return transaction.selling_order
}

func (transaction *Transaction) GetShares() int {
	return transaction.shares
}

func (transaction *Transaction) GetPrice() float64 {
	return transaction.price
}

func (transaction *Transaction) GetTotalAmount() float64 {
	return transaction.total_amount
}

func (transaction *Transaction) GetDatetime() time.Time {
	return transaction.datetime
}

// METHODS

func (transaction *Transaction) CalculateTotalAmount(shares int, price float64) {
	transaction.total_amount = float64(shares) * price
}
