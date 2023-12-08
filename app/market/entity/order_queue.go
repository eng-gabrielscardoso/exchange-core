package entity

// STRUCT DECLARATIONS

type OrderQueue struct {
	orders []*Order
}

// CONSTRUCTOR

func NewOrderQueue() *OrderQueue {
	return &OrderQueue{}
}

// METHODS

func (orderQueue *OrderQueue) Less(i, j int) bool {
	return orderQueue.orders[i].price < orderQueue.orders[j].price
}

func (orderQueue *OrderQueue) Swap(i, j int) {
	orderQueue.orders[i], orderQueue.orders[j] = orderQueue.orders[j], orderQueue.orders[i]
}

func (orderQueue *OrderQueue) Len() int {
	return len(orderQueue.orders)
}

func (orderQueue *OrderQueue) Push(x any) {
	orderQueue.orders = append(orderQueue.orders, x.(*Order))
}

func (orderQueue *OrderQueue) Pop() any {
	oldOrders := orderQueue.orders
	n := len(oldOrders)
	item := oldOrders[n-1]
	orderQueue.orders = oldOrders[0 : n-1]
	return item
}
