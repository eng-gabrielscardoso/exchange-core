package entity

import (
	"time"

	"github.com/google/uuid"
)

// STRUCT DECLARATIONS

type Transaction struct {
	ID           string
	BuyingOrder  *Order
	SellingOrder *Order
	Shares       int
	Price        float64
	TotalAmount  float64
	Datetime     time.Time
}

// CONSTRUCTOR

func NewTransaction(buyingOrder *Order, sellingOrder *Order, Shares int, price float64) *Transaction {
	total := float64(Shares) * price

	return &Transaction{
		ID:           uuid.New().String(),
		BuyingOrder:  buyingOrder,
		SellingOrder: sellingOrder,
		Shares:       Shares,
		Price:        price,
		TotalAmount:  total,
		Datetime:     time.Now(),
	}
}

// GETTERS AND SETTERS

func (transaction *Transaction) GetBuyingOrder() *Order {
	return transaction.BuyingOrder
}

func (transaction *Transaction) GetSellingOrder() *Order {
	return transaction.SellingOrder
}

func (transaction *Transaction) GetShares() int {
	return transaction.Shares
}

func (transaction *Transaction) GetPrice() float64 {
	return transaction.Price
}

func (transaction *Transaction) GetTotalAmount() float64 {
	return transaction.TotalAmount
}

func (transaction *Transaction) GetDatetime() time.Time {
	return transaction.Datetime
}

// METHODS

func (transaction *Transaction) CalculateTotalAmount(Shares int, price float64) {
	transaction.TotalAmount = float64(Shares) * price
}
