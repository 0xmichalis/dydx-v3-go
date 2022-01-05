package starkex

import (
	"fmt"
	"math/big"

	"github.com/tselementes/dydx-v3-go/constants"
)

// More info can be found here:
// https://docs.starkware.co/starkex-v3/starkex-deep-dive/starkex-specific-concepts

func ToQuantumExact(amount *big.Float, asset string) (int64, error) {
	resolution := big.NewFloat(float64(constants.ASSET_RESOLUTION[asset]))
	a := new(big.Float).Mul(amount, resolution)

	if !a.IsInt() {
		// TODO: Fix printing
		return 0, fmt.Errorf("amount %s is not a multiple of the quantum size %.0e", amount.Text('g', 15), float64(constants.ASSET_RESOLUTION[asset]))
	}
	got, _ := a.Int64()

	return got, nil
}

func ToQuantumRoundUp(amount *big.Float, asset string) int64 {
	resolution := big.NewFloat(float64(constants.ASSET_RESOLUTION[asset]))
	a := new(big.Float).Mul(amount, resolution)

	got, _ := a.Int64()
	if !a.IsInt() {
		// Round up if needed
		got++
	}

	return got
}

func ToQuantumRoundDown(amount *big.Float, asset string) int64 {
	resolution := big.NewFloat(float64(constants.ASSET_RESOLUTION[asset]))
	amount.Mul(amount, resolution)

	// Rounding down happens automatically by Int64 if needed
	got, _ := amount.Int64()

	return got
}

func NonceFromClientID(clientID string) string {
	// TODO: FILLME
	return ""
}
