package entity

import (
	"time"

	"github.com/google/uuid"
)

// STRUCT DECLARATIONS

type Transaction struct {
	ID           string
	SellingOrder *Order
	BuyingOrder  *Order
	Shares       int
	Price        float64
	TotalAmount  float64
	Datetime     time.Time
}

// CONSTRUCTOR

func NewTransaction(sellingOrder *Order, buyingOrder *Order, Shares int, price float64) *Transaction {
	total := float64(Shares) * price

	return &Transaction{
		ID:           uuid.New().String(),
		SellingOrder: sellingOrder,
		BuyingOrder:  buyingOrder,
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

func (transaction *Transaction) CalculateTotalAmount(shares int, price float64) {
	transaction.TotalAmount = float64(shares) * price
}

func (transaction *Transaction) FinishBuyOrder() {
	if transaction.BuyingOrder.PendingShares == 0 {
		transaction.BuyingOrder.Status = StatusFinished
	}
}

func (transaction *Transaction) FinishSellOrder() {
	if transaction.SellingOrder.PendingShares == 0 {
		transaction.SellingOrder.Status = StatusFinished
	}
}

func (transaction *Transaction) AddBuyOrderPendingShares(shares int) {
	transaction.BuyingOrder.PendingShares += shares
}

func (transaction *Transaction) AddSellOrderPendingShares(shares int) {
	transaction.SellingOrder.PendingShares += shares
}
