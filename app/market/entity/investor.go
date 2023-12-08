package entity

// STRUCT DECLARATIONS

type InvestorAssetPosition struct {
	id     string
	shares int
}

type Investor struct {
	id              string
	name            string
	asset_positions []*InvestorAssetPosition
}

// CONSTRUCTORS

func NewInvestor(id string) *Investor {
	return &Investor{
		id:              id,
		asset_positions: []*InvestorAssetPosition{},
	}
}

func NewInvestorAssetPosition(assetId string, shares int) *InvestorAssetPosition {
	return &InvestorAssetPosition{
		id:     assetId,
		shares: shares,
	}
}

// GETTERS AND SETTERS

func (investor *Investor) GetInvestorName() string {
	return investor.name
}

func (investor *Investor) GetAssetPosition(assetId string) *InvestorAssetPosition {
	for _, assetPosition := range investor.asset_positions {
		if assetPosition.id == assetId {
			return assetPosition
		}
	}

	return nil
}

// METHODS

func (investor *Investor) AddAssetPosition(assetPosition *InvestorAssetPosition) {
	investor.asset_positions = append(investor.asset_positions, assetPosition)
}

func (investor *Investor) UpdateAssetPosition(assetId string, shares int) {
	assetPosition := investor.GetAssetPosition(assetId)

	if assetPosition == nil {
		investor.asset_positions = append(investor.asset_positions, NewInvestorAssetPosition(assetId, shares))
	} else {
		assetPosition.shares += shares
	}
}
