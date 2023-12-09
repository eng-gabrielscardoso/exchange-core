package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewTransaction(t *testing.T) {
	sellingOrder := &Order{
		ID:       "sell-order-1",
		Shares:   5,
		Price:    10.0,
		Status:   StatusOpen,
		Asset:    &Asset{"asset-1", "test-name", 10.0},
		Investor: &Investor{ID: "investor-1"},
	}

	buyingOrder := &Order{
		ID:       "buy-order-1",
		Shares:   5,
		Price:    10.0,
		Status:   StatusOpen,
		Asset:    &Asset{"asset-1", "A1", 10.0},
		Investor: &Investor{ID: "investor-2"},
	}

	transaction := NewTransaction(sellingOrder, buyingOrder, 5, 10.0)

	assert.NotNil(t, transaction, "Transaction should not be nil")
	assert.NotEmpty(t, transaction.ID, "Transaction ID should not be empty")
	assert.Equal(t, sellingOrder, transaction.SellingOrder, "Selling order should match")
	assert.Equal(t, buyingOrder, transaction.BuyingOrder, "Buying order should match")
	assert.Equal(t, 5, transaction.Shares, "Shares should match")
	assert.Equal(t, 10.0, transaction.Price, "Price should match")
	assert.Equal(t, 50.0, transaction.TotalAmount, "Total amount should match")
	assert.WithinDuration(t, time.Now(), transaction.Datetime, time.Second, "Datetime should be close to current time")
}

func TestCalculateTotalAmount(t *testing.T) {
	transaction := &Transaction{
		Shares: 5,
		Price:  10.0,
	}

	transaction.CalculateTotalAmount(3, 15.0)

	assert.Equal(t, 45.0, transaction.TotalAmount, "Total amount should match")
}

func TestFinishBuyOrder(t *testing.T) {
	buyingOrder := &Order{
		PendingShares: 0,
		Status:        StatusOpen,
	}

	transaction := &Transaction{
		BuyingOrder: buyingOrder,
	}

	transaction.FinishBuyOrder()

	assert.Equal(t, StatusFinished, buyingOrder.Status, "Buy order status should be finished")
}

func TestFinishSellOrder(t *testing.T) {
	sellingOrder := &Order{
		PendingShares: 0,
		Status:        StatusOpen,
	}

	transaction := &Transaction{
		SellingOrder: sellingOrder,
	}

	transaction.FinishSellOrder()

	assert.Equal(t, StatusFinished, sellingOrder.Status, "Sell order status should be finished")
}

func TestAddBuyOrderPendingShares(t *testing.T) {
	buyingOrder := &Order{
		PendingShares: 5,
	}

	transaction := &Transaction{
		BuyingOrder: buyingOrder,
	}

	transaction.AddBuyOrderPendingShares(3)

	assert.Equal(t, 8, buyingOrder.PendingShares, "Pending shares should be updated")
}

func TestAddSellOrderPendingShares(t *testing.T) {
	sellingOrder := &Order{
		PendingShares: 3,
	}

	transaction := &Transaction{
		SellingOrder: sellingOrder,
	}

	transaction.AddSellOrderPendingShares(2)

	assert.Equal(t, 5, sellingOrder.PendingShares, "Pending shares should be updated")
}
