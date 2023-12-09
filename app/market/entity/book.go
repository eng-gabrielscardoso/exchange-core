package entity

import (
	"container/heap"
	"sync"
)

// TYPE DECLARATIONS

type Book struct {
	Order         []*Order
	Transactions  []*Transaction
	InputChannel  chan *Order // input
	OutputChannel chan *Order
	WaitGroup     *sync.WaitGroup
}

// CONSTRUCTORS

func NewBook(inputChannel chan *Order, outputChannel chan *Order, wg *sync.WaitGroup) *Book {
	return &Book{
		Order:         []*Order{},
		Transactions:  []*Transaction{},
		InputChannel:  inputChannel,
		OutputChannel: outputChannel,
		WaitGroup:     wg,
	}
}

// METHODS

func (b *Book) Trade() {
	buyOrders := make(map[string]*OrderQueue)
	sellOrders := make(map[string]*OrderQueue)

	for order := range b.InputChannel {
		asset := order.Asset.ID

		if buyOrders[asset] == nil {
			buyOrders[asset] = NewOrderQueue()
			heap.Init(buyOrders[asset])
		}

		if sellOrders[asset] == nil {
			sellOrders[asset] = NewOrderQueue()
			heap.Init(sellOrders[asset])
		}

		if order.OrderType == BuyOrder {
			buyOrders[asset].Push(order)

			if sellOrders[asset].Len() > 0 && sellOrders[asset].Orders[0].Price <= order.Price {
				sellOrder := sellOrders[asset].Pop().(*Order)

				if sellOrder.PendingShares > 0 {
					transaction := NewTransaction(sellOrder, order, order.Shares, sellOrder.Price)

					b.AddTransaction(transaction, b.WaitGroup)

					sellOrder.Transactions = append(sellOrder.Transactions, transaction)
					order.Transactions = append(order.Transactions, transaction)

					b.OutputChannel <- sellOrder
					b.OutputChannel <- order

					if sellOrder.PendingShares > 0 {
						sellOrders[asset].Push(sellOrder)
					}
				}
			}
		} else if order.OrderType == SellOrder {
			sellOrders[asset].Push(order)

			if buyOrders[asset].Len() > 0 && buyOrders[asset].Orders[0].Price >= order.Price {
				buyOrder := buyOrders[asset].Pop().(*Order)

				if buyOrder.PendingShares > 0 {
					transaction := NewTransaction(order, buyOrder, order.Shares, buyOrder.Price)

					b.AddTransaction(transaction, b.WaitGroup)

					buyOrder.Transactions = append(buyOrder.Transactions, transaction)
					order.Transactions = append(order.Transactions, transaction)

					b.OutputChannel <- buyOrder
					b.OutputChannel <- order

					if buyOrder.PendingShares > 0 {
						buyOrders[asset].Push(buyOrder)
					}
				}
			}
		}
	}
}

func (b *Book) AddTransaction(transaction *Transaction, wg *sync.WaitGroup) {
	defer wg.Done()

	sellingShares := transaction.SellingOrder.PendingShares
	buyingShares := transaction.BuyingOrder.PendingShares

	minShares := sellingShares
	if buyingShares < minShares {
		minShares = buyingShares
	}

	transaction.SellingOrder.Investor.UpdateAssetPosition(transaction.SellingOrder.Asset.ID, -minShares)
	transaction.AddSellOrderPendingShares(-minShares)

	transaction.BuyingOrder.Investor.UpdateAssetPosition(transaction.BuyingOrder.Asset.ID, minShares)
	transaction.AddBuyOrderPendingShares(-minShares)

	transaction.CalculateTotal(transaction.Shares, transaction.BuyingOrder.Price)
	transaction.CloseBuyOrder()
	transaction.CloseSellOrder()
	b.Transactions = append(b.Transactions, transaction)
}
