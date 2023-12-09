package dto

import e "github.com/eng-gabrielscardoso/exchange-core/app/market/entity"

// TYPE DECLARATIONS

type TradeInput struct {
	OrderID       string      `json:"order_id"`
	InvestorID    string      `json:"investor_id"`
	AssetID       string      `json:"asset_id"`
	CurrentShares int         `json:"current_shares"`
	Shares        int         `json:"shares"`
	Price         float64     `json:"price"`
	OrderType     e.OrderType `json:"order_type"`
}

type TradeOutput struct {
	OrderID           string               `json:"order_id"`
	InvestorID        string               `json:"investor_id"`
	AssetID           string               `json:"asset_id"`
	OrderType         e.OrderType          `json:"order_type"`
	Status            e.OrderStatus        `json:"status"`
	Partial           int                  `json:"partial"`
	Shares            int                  `json:"shares"`
	TransactionOutput []*TransactionOutput `json:"transactions"`
}

type TransactionOutput struct {
	TransactionId string  `json:"transaction_id"`
	BuyerID       string  `json:"buyer_id"`
	SellerID      string  `json:"seller_id"`
	AssetID       string  `json:"asset_id"`
	Price         float64 `json:"price"`
	Shares        int     `json:"shares"`
}
