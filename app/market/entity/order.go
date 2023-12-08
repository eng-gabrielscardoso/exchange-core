package entity

// STRUCT DECLARATIONS

type OrderType string

const (
	BuyOrder  OrderType = "buy"
	SellOrder OrderType = "sell"
)

type OrderStatus string

const (
	StatusOpen            OrderStatus = "open"
	StatusFilled          OrderStatus = "filled"
	StatusPartiallyFilled OrderStatus = "partially_filled"
	StatusCancelled       OrderStatus = "cancelled"
	StatusLiquidated      OrderStatus = "liquidated"
	StatusPending         OrderStatus = "pending"
	StatusRejected        OrderStatus = "rejected"
	StatusFailed          OrderStatus = "failed"
	StatusExpired         OrderStatus = "expired"
	StatusUnderReview     OrderStatus = "under_review"
	StatusFinished        OrderStatus = "finished"
)

type Order struct {
	id             string
	investor       *Investor
	asset          *Asset
	shares         int
	pending_shares int
	price          float64
	order_type     OrderType
	status         OrderStatus
	transactions   []*Transaction
}

// CONSTRUCTORS

func NewOrder(orderId string, investor *Investor, asset *Asset, shares int, price float64, orderType OrderType) *Order {
	return &Order{
		id:           orderId,
		investor:     investor,
		asset:        asset,
		shares:       shares,
		price:        price,
		order_type:   orderType,
		status:       StatusOpen,
		transactions: []*Transaction{},
	}
}

// GETTERS AND SETTERS

func (order *Order) GetInvestor() *Investor {
	return order.investor
}

func (order *Order) GetAsset() *Asset {
	return order.asset
}

func (order *Order) GetShares() int {
	return order.shares
}

func (order *Order) GetPendingShares() int {
	return order.pending_shares
}

func (order *Order) GetPrice() float64 {
	return order.price
}

func (order *Order) GetOrderType() OrderType {
	return order.order_type
}

func (order *Order) GetStatus() OrderStatus {
	return order.status
}

func (order *Order) GetTransactions() []*Transaction {
	return order.transactions
}
