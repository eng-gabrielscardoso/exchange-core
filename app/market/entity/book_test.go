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
	assert.Equal(StatusFinished, order.status, "Order 1 should be closed")
	assert.Equal(0, order.pending_shares, "Order 1 should have 0 pending_shares")
	assert.Equal(StatusFinished, order2.status, "Order 2 should be closed")
	assert.Equal(0, order2.pending_shares, "Order 2 should have 0 pending_shares")

	assert.Equal(5, investorAssetPosition.shares, "Investor 1 should have 5 shares of asset 1")
	assert.Equal(5, investor2.GetAssetPosition("asset1").shares, "Investor 2 should have 5 shares of asset 1")
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
	assert.Equal(StatusOpen, order.status, "Order 1 should be closed")
	assert.Equal(5, order.pending_shares, "Order 1 should have 5 pending_shares")
	assert.Equal(StatusOpen, order2.status, "Order 2 should be closed")
	assert.Equal(5, order2.pending_shares, "Order 2 should have 5 pending_shares")
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

	assert.Equal(StatusFinished, order.status, "Order 1 should be closed")
	assert.Equal(0, order.pending_shares, "Order 1 should have 0 pending_shares")

	assert.Equal(StatusOpen, order2.status, "Order 2 should be OPEN")
	assert.Equal(2, order2.pending_shares, "Order 2 should have 2 pending_shares")

	assert.Equal(0, investorAssetPosition.shares, "Investor 1 should have 0 shares of asset 1")
	assert.Equal(3, investor2.GetAssetPosition("asset1").shares, "Investor 2 should have 3 shares of asset 1")

	waitGroup.Add(1)
	order3 := NewOrder("3", investor3, asset1, 2, 5.0, SellOrder)
	inputChannel <- order3
	waitGroup.Wait()

	assert.Equal(StatusFinished, order3.status, "Order 3 should be closed")
	assert.Equal(0, order3.pending_shares, "Order 3 should have 0 pending_shares")

	assert.Equal(StatusFinished, order2.status, "Order 2 should be CLOSED")
	assert.Equal(0, order2.pending_shares, "Order 2 should have 0 pending_shares")

	assert.Equal(2, len(book.transactions), "Should have 2 transactions")
	assert.Equal(15.0, float64(book.transactions[0].total_amount), "Transaction should have price 15")
	assert.Equal(10.0, float64(book.transactions[1].total_amount), "Transaction should have price 10")
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
	assert.Equal(StatusFinished, order.status, "Order 1 should be closed")
	assert.Equal(0, order.pending_shares, "Order 1 should have 0 pending_shares")

	assert.Equal(StatusOpen, order2.status, "Order 2 should be OPEN")
	assert.Equal(2, order2.pending_shares, "Order 2 should have 2 pending_shares")

	assert.Equal(0, investorAssetPosition.shares, "Investor 1 should have 0 shares of asset 1")
	assert.Equal(3, investor2.GetAssetPosition("asset1").shares, "Investor 2 should have 3 shares of asset 1")

	waitGroup.Add(1)
	order3 := NewOrder("3", investor3, asset1, 3, 4.5, SellOrder)
	inputChannel <- order3

	waitGroup.Wait()

	assert.Equal(StatusOpen, order3.status, "Order 3 should be open")
	assert.Equal(1, order3.pending_shares, "Order 3 should have 1 pending_shares")

	assert.Equal(StatusFinished, order2.status, "Order 2 should be CLOSED")
	assert.Equal(0, order2.pending_shares, "Order 2 should have 0 pending_shares")
}

func TestNoMatch(t *testing.T) {
	asset1 := NewAsset("asset1", "Asset 1", 100)

	investor := NewInvestor("1")
	investor2 := NewInvestor("2")

	investorAssetPosition := NewInvestorAssetPosition(asset1.id, 3)
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
	assert.Equal(StatusOpen, order.status, "Order 1 should be closed")
	assert.Equal(StatusOpen, order2.status, "Order 2 should be OPEN")
	assert.Equal(3, order.pending_shares, "Order 1 should have 3 pending_shares")
	assert.Equal(5, order2.pending_shares, "Order 2 should have 5 pending_shares")
}
