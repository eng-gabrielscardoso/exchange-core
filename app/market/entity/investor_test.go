package entity

import (
	"testing"

	"github.com/google/uuid"
)

func TestNewInvestor(t *testing.T) {
	id := uuid.New().String()
	investor := NewInvestor(id)

	if investor == nil {
		t.Error("Expected a valid Investor object, but got nil")
	}

	if investor.id != id {
		t.Errorf("Expected id to be %s, but got %s", id, investor.id)
	}

	if len(investor.asset_positions) != 0 {
		t.Error("Expected no asset positions initially, but got some")
	}
}

func TestInvestorMethods(t *testing.T) {
	id := uuid.New().String()
	investor := NewInvestor(id)
	assetId := uuid.New().String()
	shares := 100

	assetPosition := NewInvestorAssetPosition(assetId, shares)
	investor.AddAssetPosition(assetPosition)

	if len(investor.asset_positions) != 1 {
		t.Error("Expected one asset position after adding, but got a different count")
	}

	newShares := 150
	investor.UpdateAssetPosition(assetId, newShares)

	if len(investor.asset_positions) != 1 {
		t.Error("Expected one asset position after updating, but got a different count")
	}

	updatedAssetPosition := investor.GetAssetPosition(assetId)
	if updatedAssetPosition == nil {
		t.Error("Expected to find the updated asset position, but got nil")
	}

	updatedShares := shares + newShares
	if updatedAssetPosition.shares != updatedShares {
		t.Errorf("Expected shares to be %d, but got %d", updatedShares, updatedAssetPosition.shares)
	}

	newAssetId := "XYZ"
	investor.UpdateAssetPosition(newAssetId, 200)

	if len(investor.asset_positions) != 2 {
		t.Error("Expected two asset positions after updating with a new asset, but got a different count")
	}

	newAssetPosition := investor.GetAssetPosition(newAssetId)
	if newAssetPosition == nil {
		t.Error("Expected to find the new asset position, but got nil")
	}
}
