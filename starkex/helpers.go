package starkex

import (
	"fmt"
	"math/big"

	"github.com/tselementes/dydx-v3-go/constants"
)

// More info can be found here:
// https://docs.starkware.co/starkex-v3/starkex-deep-dive/starkex-specific-concepts

func ToQuantumExact(amount float64, asset string) (int64, error) {
	resolution := big.NewFloat(float64(constants.ASSET_RESOLUTION[asset]))
	fAmount := big.NewFloat(amount)
	fAmount.Mul(fAmount, resolution)

	if !fAmount.IsInt() {
		// TODO: Fix float printing here
		return 0, fmt.Errorf("amount %s is not a multiple of the quantum size %.0e", big.NewFloat(amount).String(), float64(constants.ASSET_RESOLUTION[asset]))
	}
	got, _ := fAmount.Int64()
	return got, nil
}

func ToQuantumRoundUp(amount float64, asset string) int64 {
	resolution := big.NewFloat(float64(constants.ASSET_RESOLUTION[asset]))
	fAmount := big.NewFloat(amount)
	fAmount.Mul(fAmount, resolution)

	got, _ := fAmount.Int64()
	if !fAmount.IsInt() {
		got++
	}

	return got
}

func ToQuantumRoundDown(amount float64, asset string) int64 {
	resolution := big.NewFloat(float64(constants.ASSET_RESOLUTION[asset]))
	fAmount := big.NewFloat(amount)
	fAmount.Mul(fAmount, resolution)

	got, _ := fAmount.Int64()
	return got
}
