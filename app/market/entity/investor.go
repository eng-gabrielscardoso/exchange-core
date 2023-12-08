package entity

// STRUCT DECLARATIONS

type InvestorAssetPosition struct {
	ID     string
	Shares int
}

type Investor struct {
	ID             string
	Name           string
	AssetPositions []*InvestorAssetPosition
}

// CONSTRUCTORS

func NewInvestor(id string) *Investor {
	return &Investor{
		ID:             id,
		AssetPositions: []*InvestorAssetPosition{},
	}
}

func NewInvestorAssetPosition(assetId string, shares int) *InvestorAssetPosition {
	return &InvestorAssetPosition{
		ID:     assetId,
		Shares: shares,
	}
}

// GETTERS AND SETTERS

func (investor *Investor) GetInvestorName() string {
	return investor.Name
}

func (investor *Investor) GetAssetPosition(assetId string) *InvestorAssetPosition {
	for _, assetPosition := range investor.AssetPositions {
		if assetPosition.ID == assetId {
			return assetPosition
		}
	}

	return nil
}

// METHODS

func (investor *Investor) AddAssetPosition(assetPosition *InvestorAssetPosition) {
	investor.AssetPositions = append(investor.AssetPositions, assetPosition)
}

func (investor *Investor) UpdateAssetPosition(assetId string, shares int) {
	assetPosition := investor.GetAssetPosition(assetId)

	if assetPosition == nil {
		investor.AssetPositions = append(investor.AssetPositions, NewInvestorAssetPosition(assetId, shares))
	} else {
		assetPosition.Shares += shares
	}
}
