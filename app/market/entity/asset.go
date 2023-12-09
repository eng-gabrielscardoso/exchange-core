package entity

// TYPE DECLARATIONS

type Asset struct {
	ID           string
	Name         string
	MarketVolume int
}

// CONSTRUCTORS

func NewAsset(id string, name string, marketVolume int) *Asset {
	return &Asset{
		ID:           id,
		Name:         name,
		MarketVolume: marketVolume,
	}
}
