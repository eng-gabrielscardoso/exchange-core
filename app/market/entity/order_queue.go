package entity

// STRUCT DECLARATIONS

type OrderQueue struct {
	Orders []*Order
}

// CONSTRUCTOR

func NewOrderQueue() *OrderQueue {
	return &OrderQueue{}
}

// METHODS

func (orderQueue *OrderQueue) Less(i, j int) bool {
	return orderQueue.Orders[i].Price < orderQueue.Orders[j].Price
}

func (orderQueue *OrderQueue) Swap(i, j int) {
	orderQueue.Orders[i], orderQueue.Orders[j] = orderQueue.Orders[j], orderQueue.Orders[i]
}

func (orderQueue *OrderQueue) Len() int {
	return len(orderQueue.Orders)
}

func (orderQueue *OrderQueue) Push(x any) {
	orderQueue.Orders = append(orderQueue.Orders, x.(*Order))
}

func (orderQueue *OrderQueue) Pop() any {
	oldOrders := orderQueue.Orders
	n := len(oldOrders)
	item := oldOrders[n-1]
	orderQueue.Orders = oldOrders[0 : n-1]
	return item
}
