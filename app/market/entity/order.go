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
	ID            string
	Investor      *Investor
	Asset         *Asset
	Shares        int
	PendingShares int
	Price         float64
	OrderType     OrderType
	Status        OrderStatus
	Transactions  []*Transaction
}

// CONSTRUCTORS

func NewOrder(orderId string, investor *Investor, asset *Asset, shares int, price float64, orderType OrderType) *Order {
	return &Order{
		ID:           orderId,
		Investor:     investor,
		Asset:        asset,
		Shares:       shares,
		Price:        price,
		OrderType:    orderType,
		Status:       StatusOpen,
		Transactions: []*Transaction{},
	}
}

// GETTERS AND SETTERS

func (order *Order) GetInvestor() *Investor {
	return order.Investor
}

func (order *Order) GetAsset() *Asset {
	return order.Asset
}

func (order *Order) GetShares() int {
	return order.Shares
}

func (order *Order) GetPendingShares() int {
	return order.PendingShares
}

func (order *Order) GetPrice() float64 {
	return order.Price
}

func (order *Order) GetOrderType() OrderType {
	return order.OrderType
}

func (order *Order) GetStatus() OrderStatus {
	return order.Status
}

func (order *Order) GetTransactions() []*Transaction {
	return order.Transactions
}
