package entity

import "testing"

func TestNewAsset(t *testing.T) {
	id := "BTC"
	name := "Bitcoin"
	volume := 1000000

	asset := NewAsset(id, name, volume)

	if asset.ID != id {
		t.Errorf("Expected: %s, Received: %s", id, asset.ID)
	}

	if asset.Name != name {
		t.Errorf("Expected: %s, Received: %s", name, asset.Name)
	}

	if asset.MarketVolume != volume {
		t.Errorf("Expected: %d, Received: %d", volume, asset.MarketVolume)
	}
}

func TestAssetGetters(t *testing.T) {
	id := "ETH"
	name := "Ethereum"
	volume := 500000

	asset := NewAsset(id, name, volume)

	if asset.GetName() != name {
		t.Errorf("Expected: %s, Received: %s", name, asset.GetName())
	}

	if asset.GetMarketVolume() != volume {
		t.Errorf("Expected: %d, Received: %d", volume, asset.GetMarketVolume())
	}
}
