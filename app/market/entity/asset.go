package entity

// STRUCT DECLARATIONS

type Asset struct {
	id            string
	name          string
	market_volume int
}

// CONSTRUCTORS

func NewAsset(id string, name string, volume int) *Asset {
	return &Asset{
		id:            id,
		name:          name,
		market_volume: volume,
	}
}

// GETTERS AND SETTERS

func (asset *Asset) GetName() string {
	return asset.name
}

func (asset *Asset) GetMarketVolume() int {
	return asset.market_volume
}
