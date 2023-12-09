package entity

// STRUCT DECLARATIONS

type InvestorAssetPosition struct {
	ID     string
	Shares int
}

type Investor struct {
	ID            string
	Name          string
	AssetPosition []*InvestorAssetPosition
}

// CONSTRUCTORS

func NewInvestorAssetPosition(assetId string, shares int) *InvestorAssetPosition {
	return &InvestorAssetPosition{
		ID:     assetId,
		Shares: shares,
	}
}

func NewInvestor(id string) *Investor {
	return &Investor{
		ID:            id,
		AssetPosition: []*InvestorAssetPosition{},
	}
}

// GETTERS AND SETTERS

func (investor *Investor) GetInvestorName() string {
	return investor.Name
}

func (investor *Investor) GetAssetPosition(assetId string) *InvestorAssetPosition {
	for _, assetPosition := range investor.AssetPosition {
		if assetPosition.ID == assetId {
			return assetPosition
		}
	}

	return nil
}

// METHODS

func (investor *Investor) AddAssetPosition(assetPosition *InvestorAssetPosition) {
	investor.AssetPosition = append(investor.AssetPosition, assetPosition)
}

func (investor *Investor) UpdateAssetPosition(assetId string, shares int) {
	assetPosition := investor.GetAssetPosition(assetId)

	if assetPosition == nil {
		investor.AssetPosition = append(investor.AssetPosition, NewInvestorAssetPosition(assetId, shares))
	} else {
		assetPosition.Shares += shares
	}
}
