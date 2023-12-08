package entity

// STRUCT DECLARATIONS

type Asset struct {
	ID           string
	Name         string
	MarketVolume int
}

// CONSTRUCTORS

func NewAsset(ID string, Name string, volume int) *Asset {
	return &Asset{
		ID:           ID,
		Name:         Name,
		MarketVolume: volume,
	}
}

// GETTERS AND SETTERS

func (asset *Asset) GetName() string {
	return asset.Name
}

func (asset *Asset) GetMarketVolume() int {
	return asset.MarketVolume
}
