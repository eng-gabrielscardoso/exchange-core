package entity

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuyAsset(t *testing.T) {
	asset1 := NewAsset("asset1", "Asset 1", 100)

	investor := NewInvestor("1")
	investor2 := NewInvestor("2")

	investorAssetPosition := NewInvestorAssetPosition("asset1", 10)
	investor.AddAssetPosition(investorAssetPosition)

	waitGroup := sync.WaitGroup{}
	inputChannel := make(chan *Order)
	outputChannel := make(chan *Order)

	book := NewBook(inputChannel, outputChannel, &waitGroup)

	go book.Trade()

	waitGroup.Add(1)
	order := NewOrder("1", investor, asset1, 5, 5, SellOrder)
	inputChannel <- order

	order2 := NewOrder("2", investor2, asset1, 5, 5, BuyOrder)
	inputChannel <- order2
	waitGroup.Wait()

	assert := assert.New(t)
	assert.Equal(StatusFinished, order.Status, "Order 1 should be finished")
	assert.Equal(0, order.PendingShares, "Order 1 should have 0 PendingShares")
	assert.Equal(StatusFinished, order2.Status, "Order 2 should be finished")
	assert.Equal(0, order2.PendingShares, "Order 2 should have 0 PendingShares")

	assert.Equal(5, investorAssetPosition.Shares, "Investor 1 should have 5 shares of asset 1")
	assert.Equal(5, investor2.GetAssetPosition("asset1").Shares, "Investor 2 should have 5 shares of asset 1")
}

func TestBuyAssetWithDifferentAssents(t *testing.T) {
	asset1 := NewAsset("asset1", "Asset 1", 100)
	asset2 := NewAsset("asset2", "Asset 2", 100)

	investor := NewInvestor("1")
	investor2 := NewInvestor("2")

	investorAssetPosition := NewInvestorAssetPosition("asset1", 10)
	investor.AddAssetPosition(investorAssetPosition)

	investorAssetPosition2 := NewInvestorAssetPosition("asset2", 10)
	investor2.AddAssetPosition(investorAssetPosition2)

	waitGroup := sync.WaitGroup{}
	inputChannel := make(chan *Order)
	outputChannel := make(chan *Order)

	book := NewBook(inputChannel, outputChannel, &waitGroup)

	go book.Trade()

	order := NewOrder("1", investor, asset1, 5, 5, SellOrder)
	inputChannel <- order

	order2 := NewOrder("2", investor2, asset2, 5, 5, BuyOrder)
	inputChannel <- order2

	assert := assert.New(t)
	assert.Equal(StatusOpen, order.Status, "Order 1 should be finished")
	assert.Equal(5, order.PendingShares, "Order 1 should have 5 PendingShares")
	assert.Equal(StatusOpen, order2.Status, "Order 2 should be finished")
	assert.Equal(5, order2.PendingShares, "Order 2 should have 5 PendingShares")
}

func TestBuyPartialAsset(t *testing.T) {
	asset1 := NewAsset("asset1", "Asset 1", 100)

	investor := NewInvestor("1")
	investor2 := NewInvestor("2")
	investor3 := NewInvestor("3")

	investorAssetPosition := NewInvestorAssetPosition("asset1", 3)
	investor.AddAssetPosition(investorAssetPosition)

	investorAssetPosition2 := NewInvestorAssetPosition("asset1", 5)
	investor3.AddAssetPosition(investorAssetPosition2)

	waitGroup := sync.WaitGroup{}
	inputChannel := make(chan *Order)
	outputChannel := make(chan *Order)

	book := NewBook(inputChannel, outputChannel, &waitGroup)

	go book.Trade()

	waitGroup.Add(1)
	order2 := NewOrder("1", investor2, asset1, 5, 5.0, BuyOrder)
	inputChannel <- order2

	order := NewOrder("2", investor, asset1, 3, 5.0, SellOrder)
	inputChannel <- order

	assert := assert.New(t)

	go func() {
		for range outputChannel {
		}
	}()

	waitGroup.Wait()

	assert.Equal(StatusFinished, order.Status, "Order 1 should be finished")
	assert.Equal(0, order.PendingShares, "Order 1 should have 0 PendingShares")

	assert.Equal(StatusOpen, order2.Status, "Order 2 should be open")
	assert.Equal(2, order2.PendingShares, "Order 2 should have 2 PendingShares")

	assert.Equal(0, investorAssetPosition.Shares, "Investor 1 should have 0 shares of asset 1")
	assert.Equal(3, investor2.GetAssetPosition("asset1").Shares, "Investor 2 should have 3 shares of asset 1")

	waitGroup.Add(1)
	order3 := NewOrder("3", investor3, asset1, 2, 5.0, SellOrder)
	inputChannel <- order3
	waitGroup.Wait()

	assert.Equal(StatusFinished, order3.Status, "Order 3 should be finished")
	assert.Equal(0, order3.PendingShares, "Order 3 should have 0 PendingShares")

	assert.Equal(StatusFinished, order2.Status, "Order 2 should be finished")
	assert.Equal(0, order2.PendingShares, "Order 2 should have 0 PendingShares")

	assert.Equal(2, len(book.Transactions), "Should have 2 transactions")
	assert.Equal(15.0, float64(book.Transactions[0].TotalAmount), "Transaction should have price 15")
	assert.Equal(10.0, float64(book.Transactions[1].TotalAmount), "Transaction should have price 10")
}

func TestBuyWithDifferentPrice(t *testing.T) {
	asset1 := NewAsset("asset1", "Asset 1", 100)

	investor := NewInvestor("1")
	investor2 := NewInvestor("2")
	investor3 := NewInvestor("3")

	investorAssetPosition := NewInvestorAssetPosition("asset1", 3)
	investor.AddAssetPosition(investorAssetPosition)

	investorAssetPosition2 := NewInvestorAssetPosition("asset1", 5)
	investor3.AddAssetPosition(investorAssetPosition2)

	waitGroup := sync.WaitGroup{}
	inputChannel := make(chan *Order)

	outputChannel := make(chan *Order)

	book := NewBook(inputChannel, outputChannel, &waitGroup)

	go book.Trade()

	waitGroup.Add(1)
	order2 := NewOrder("2", investor2, asset1, 5, 5.0, BuyOrder)
	inputChannel <- order2

	order := NewOrder("1", investor, asset1, 3, 4.0, SellOrder)
	inputChannel <- order

	go func() {
		for range outputChannel {
		}
	}()

	waitGroup.Wait()

	assert := assert.New(t)
	assert.Equal(StatusFinished, order.Status, "Order 1 should be finished")
	assert.Equal(0, order.PendingShares, "Order 1 should have 0 PendingShares")

	assert.Equal(StatusOpen, order2.Status, "Order 2 should be open")
	assert.Equal(2, order2.PendingShares, "Order 2 should have 2 PendingShares")

	assert.Equal(0, investorAssetPosition.Shares, "Investor 1 should have 0 shares of asset 1")
	assert.Equal(3, investor2.GetAssetPosition("asset1").Shares, "Investor 2 should have 3 shares of asset 1")

	waitGroup.Add(1)
	order3 := NewOrder("3", investor3, asset1, 3, 4.5, SellOrder)
	inputChannel <- order3

	waitGroup.Wait()

	assert.Equal(StatusOpen, order3.Status, "Order 3 should be open")
	assert.Equal(1, order3.PendingShares, "Order 3 should have 1 PendingShares")

	assert.Equal(StatusFinished, order2.Status, "Order 2 should be finished")
	assert.Equal(0, order2.PendingShares, "Order 2 should have 0 PendingShares")
}

func TestNoMatch(t *testing.T) {
	asset1 := NewAsset("asset1", "Asset 1", 100)

	investor := NewInvestor("1")
	investor2 := NewInvestor("2")

	investorAssetPosition := NewInvestorAssetPosition("asset1", 3)
	investor.AddAssetPosition(investorAssetPosition)

	waitGroup := sync.WaitGroup{}

	inputChannel := make(chan *Order)

	outputChannel := make(chan *Order)

	book := NewBook(inputChannel, outputChannel, &waitGroup)

	go book.Trade()

	waitGroup.Add(0)
	order := NewOrder("1", investor, asset1, 3, 6.0, SellOrder)
	inputChannel <- order

	order2 := NewOrder("2", investor2, asset1, 5, 5.0, BuyOrder)
	inputChannel <- order2

	go func() {
		for range outputChannel {
		}
	}()

	waitGroup.Wait()

	assert := assert.New(t)
	assert.Equal(StatusOpen, order.Status, "Order 1 should be finished")
	assert.Equal(StatusOpen, order2.Status, "Order 2 should be open")
	assert.Equal(3, order.PendingShares, "Order 1 should have 3 PendingShares")
	assert.Equal(5, order2.PendingShares, "Order 2 should have 5 PendingShares")
}
