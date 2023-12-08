package entity

import (
	"container/heap"
	"sync"
)

// STRUCT DECLARATIONS

type Book struct {
	orders        []*Order
	transactions  []*Transaction
	orders_input  chan *Order
	orders_output chan *Order
	wait_group    *sync.WaitGroup
}

// CONSTRUCTOR

func NewBook(inputChannel chan *Order, outputChannel chan *Order, waitGroup *sync.WaitGroup) *Book {
	return &Book{
		orders:        []*Order{},
		transactions:  []*Transaction{},
		orders_input:  inputChannel,
		orders_output: outputChannel,
		wait_group:    waitGroup,
	}
}

// GETTERS AND SETTERS

func (book *Book) GetOrders() []*Order {
	return book.orders
}

func (book *Book) GetTransactions() []*Transaction {
	return book.transactions
}

func (book *Book) GetOrdersInput() chan *Order {
	return book.orders_input
}

func (book *Book) GetOrdersOutput() chan *Order {
	return book.orders_output
}

func (book *Book) GetWaitGroup() *sync.WaitGroup {
	return book.wait_group
}

// METHODS

func (book *Book) AddTransaction(transaction *Transaction, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()

	buyingShares := transaction.buying_order.pending_shares
	sellingShares := transaction.selling_order.pending_shares

	minShares := sellingShares

	if buyingShares < minShares {
		minShares = buyingShares
	}

	transaction.buying_order.investor.UpdateAssetPosition(transaction.buying_order.asset.id, minShares)
	transaction.buying_order.pending_shares -= minShares
	transaction.selling_order.investor.UpdateAssetPosition(transaction.selling_order.asset.id, -minShares)
	transaction.selling_order.pending_shares -= minShares

	transaction.CalculateTotalAmount(transaction.shares, transaction.buying_order.price)

	if transaction.buying_order.pending_shares == 0 {
		transaction.buying_order.status = StatusFinished
	} else {
		transaction.buying_order.status = StatusPartiallyFilled
	}

	if transaction.selling_order.pending_shares == 0 {
		transaction.selling_order.status = StatusFinished
	} else {
		transaction.selling_order.status = StatusPartiallyFilled
	}

	book.transactions = append(book.transactions, transaction)
}

func (book *Book) Trade() {
	buyOrders := NewOrderQueue()
	sellOrders := NewOrderQueue()

	heap.Init(buyOrders)
	heap.Init(sellOrders)

	for order := range book.orders_input {
		if order.order_type == BuyOrder {
			buyOrders.Push(order)

			if sellOrders.Len() > 0 && sellOrders.orders[0].price <= order.price {
				sellOrder := sellOrders.Pop().(*Order)

				if sellOrder.pending_shares > 0 {
					transaction := NewTransaction(sellOrder, order, order.shares, order.price)

					book.AddTransaction(transaction, book.wait_group)

					sellOrder.transactions = append(sellOrder.transactions, transaction)
					order.transactions = append(order.transactions, transaction)

					book.orders_output <- order
					book.orders_output <- sellOrder

					if sellOrder.pending_shares > 0 {
						sellOrders.Push(sellOrder)
					}
				}
			}
		}

		if order.order_type == SellOrder {
			sellOrders.Push(order)

			if buyOrders.Len() > 0 && buyOrders.orders[0].price >= order.price {
				buyOrder := buyOrders.Pop().(*Order)

				if buyOrder.pending_shares > 0 {
					transaction := NewTransaction(order, buyOrder, order.shares, buyOrder.price)

					book.AddTransaction(transaction, book.wait_group)

					buyOrder.transactions = append(buyOrder.transactions, transaction)
					order.transactions = append(order.transactions, transaction)

					book.orders_output <- order
					book.orders_output <- buyOrder

					if buyOrder.pending_shares > 0 {
						sellOrders.Push(buyOrder)
					}
				}
			}
		}
	}
}
