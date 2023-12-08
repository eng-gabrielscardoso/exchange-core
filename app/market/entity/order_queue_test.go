package entity

import (
	"container/heap"
	"testing"

	"github.com/google/uuid"
)

func OrderFactory() []*Order {
	orders := []*Order{}

	for range []int{1, 2, 3, 4, 5} {
		investor := NewInvestor(uuid.New().String())
		asset := NewAsset(uuid.New().String(), "Asset", 100)
		order := NewOrder(uuid.New().String(), investor, asset, 1000, 10.50, BuyOrder)
		orders = append(orders, order)
	}

	return orders
}

func TestNewOrderQueue(t *testing.T) {
	orderQueue := NewOrderQueue()

	if orderQueue == nil {
		t.Errorf("Expected an order queue to be defined.")
	}
}

func TestLess(t *testing.T) {
	orders := OrderFactory()

	orderQueue := NewOrderQueue()

	heap.Init(orderQueue)

	for _, order := range orders {
		orderQueue.Push(order)
	}

	if orderQueue.orders == nil {
		t.Errorf("Expected pushed orders to queue but nothing was sent.")
	}

	if orderQueue.Len() != 5 {
		t.Errorf("Expected order queue with %d length but got %d", 5, orderQueue.Len())
	}
}

func TestSwap(t *testing.T) {
	orders := OrderFactory()

	orderQueue := NewOrderQueue()

	heap.Init(orderQueue)

	for _, order := range orders {
		orderQueue.Push(order)
	}

	orderQueue.Swap(0, 4)

	if orderQueue.orders[0].price != orders[4].price || orderQueue.orders[4].price != orders[0].price {
		t.Error("Expected orders to be swapped, but they were not")
	}
}

func TestPop(t *testing.T) {
	orders := OrderFactory()

	orderQueue := NewOrderQueue()

	heap.Init(orderQueue)

	for _, order := range orders {
		orderQueue.Push(order)
	}

	poppedOrder := orderQueue.Pop().(*Order)

	lastOrder := orders[len(orders)-1]

	if poppedOrder.price != lastOrder.price {
		t.Error("Expected popped order to be the last order, but it was not")
	}

	if orderQueue.Len() != len(orders)-1 {
		t.Error("Expected order queue length to be reduced after Pop, but it was not")
	}
}
