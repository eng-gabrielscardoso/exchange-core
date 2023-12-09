package entity

// TYPE DECLARATIONS

type OrderType string

const (
	BuyOrder  OrderType = "buy"
	SellOrder OrderType = "sell"
)

type OrderStatus string

const (
	StatusOpen   OrderStatus = "open"
	StatusClosed OrderStatus = "closed"
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

func NewOrder(orderID string, investor *Investor, asset *Asset, shares int, price float64, orderType OrderType) *Order {
	return &Order{
		ID:            orderID,
		Investor:      investor,
		Asset:         asset,
		Shares:        shares,
		PendingShares: shares,
		Price:         price,
		OrderType:     orderType,
		Status:        StatusOpen,
		Transactions:  []*Transaction{},
	}
}
