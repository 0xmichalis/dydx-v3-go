package starkex_test

import (
	"math/big"
	"testing"

	"github.com/tselementes/dydx-v3-go/constants"
	"github.com/tselementes/dydx-v3-go/starkex"
)

func TestToQuantumExact(t *testing.T) {
	humanAmount := big.NewFloat(145.000600001)
	asset := constants.SYNTHETIC_ASSET_MAP[constants.MARKET_ETH_USD]

	expected := int64(145000600001)
	got, err := starkex.ToQuantumExact(humanAmount, asset)
	if err != nil {
		t.Fatalf("got unexpected error: %v", err)
	}
	if got != expected {
		t.Fatalf("expected %d, got %d", expected, got)
	}
}

func TestToQuantumRoundUp(t *testing.T) {
	humanAmount := big.NewFloat(145.0006000011)
	asset := constants.SYNTHETIC_ASSET_MAP[constants.MARKET_ETH_USD]

	expected := int64(145000600002)
	got := starkex.ToQuantumRoundUp(humanAmount, asset)
	if got != expected {
		t.Fatalf("expected %d, got %d", expected, got)
	}
}

func TestToQuantumRoundDown(t *testing.T) {
	humanAmount := big.NewFloat(145.0006000001)
	asset := constants.SYNTHETIC_ASSET_MAP[constants.MARKET_ETH_USD]

	expected := int64(145000600000)
	got := starkex.ToQuantumRoundDown(humanAmount, asset)
	if got != expected {
		t.Fatalf("expected %d, got %d", expected, got)
	}
}
