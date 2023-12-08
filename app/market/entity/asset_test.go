package entity

import "testing"

func TestNewAsset(t *testing.T) {
	id := "BTC"
	name := "Bitcoin"
	volume := 1000000

	asset := NewAsset(id, name, volume)

	if asset.id != id {
		t.Errorf("Expected: %s, Received: %s", id, asset.id)
	}

	if asset.name != name {
		t.Errorf("Expected: %s, Received: %s", name, asset.name)
	}

	if asset.market_volume != volume {
		t.Errorf("Expected: %d, Received: %d", volume, asset.market_volume)
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
