package entity

import (
	"container/heap"
	"sync"
)

// STRUCT DECLARATIONS

type Book struct {
	Orders       []*Order
	Transactions []*Transaction
	OrdersInput  chan *Order
	OrdersOutput chan *Order
	WaitGroup    *sync.WaitGroup
}

// CONSTRUCTOR

func NewBook(inputChannel chan *Order, outputChannel chan *Order, waitGroup *sync.WaitGroup) *Book {
	return &Book{
		Orders:       []*Order{},
		Transactions: []*Transaction{},
		OrdersInput:  inputChannel,
		OrdersOutput: outputChannel,
		WaitGroup:    waitGroup,
	}
}

// GETTERS AND SETTERS

func (book *Book) GetOrders() []*Order {
	return book.Orders
}

func (book *Book) GetTransactions() []*Transaction {
	return book.Transactions
}

func (book *Book) GetOrdersInput() chan *Order {
	return book.OrdersInput
}

func (book *Book) GetOrdersOutput() chan *Order {
	return book.OrdersOutput
}

func (book *Book) GetWaitGroup() *sync.WaitGroup {
	return book.WaitGroup
}

// METHODS

func (book *Book) AddTransaction(transaction *Transaction, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()

	sellingShares := transaction.SellingOrder.PendingShares
	buyingShares := transaction.BuyingOrder.PendingShares

	minShares := sellingShares

	if buyingShares < minShares {
		minShares = buyingShares
	}

	transaction.SellingOrder.Investor.UpdateAssetPosition(transaction.SellingOrder.Asset.ID, -minShares)
	transaction.SellingOrder.PendingShares -= minShares
	transaction.BuyingOrder.Investor.UpdateAssetPosition(transaction.BuyingOrder.Asset.ID, minShares)
	transaction.BuyingOrder.PendingShares -= minShares

	transaction.CalculateTotalAmount(transaction.Shares, transaction.BuyingOrder.Price)

	if transaction.BuyingOrder.PendingShares == 0 {
		transaction.BuyingOrder.Status = StatusFinished
	} else {
		transaction.BuyingOrder.Status = StatusPartiallyFilled
	}

	if transaction.SellingOrder.PendingShares == 0 {
		transaction.SellingOrder.Status = StatusFinished
	} else {
		transaction.SellingOrder.Status = StatusPartiallyFilled
	}

	book.Transactions = append(book.Transactions, transaction)
}

func (book *Book) Trade() {
	buyOrders := make(map[string]*OrderQueue)
	sellOrders := make(map[string]*OrderQueue)

	for order := range book.OrdersInput {
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
					transaction := NewTransaction(sellOrder, order, order.Shares, order.Price)

					book.AddTransaction(transaction, book.WaitGroup)

					sellOrder.Transactions = append(sellOrder.Transactions, transaction)
					order.Transactions = append(order.Transactions, transaction)

					book.OrdersOutput <- sellOrder
					book.OrdersOutput <- order

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

					book.AddTransaction(transaction, book.WaitGroup)

					buyOrder.Transactions = append(buyOrder.Transactions, transaction)
					order.Transactions = append(order.Transactions, transaction)

					book.OrdersOutput <- buyOrder
					book.OrdersOutput <- order

					if buyOrder.PendingShares > 0 {
						sellOrders[asset].Push(buyOrder)
					}
				}
			}
		}
	}
}
