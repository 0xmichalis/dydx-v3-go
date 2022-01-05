package starkex

import (
	"fmt"
	"math"
	"math/big"
	"strconv"

	"github.com/tselementes/dydx-v3-go/constants"
)

type OrderSignatureRequest struct {
	NetworkID              int
	Market                 string
	Side                   string
	PositionID             string
	Size                   string
	Price                  string
	LimitFee               string
	ClientID               string
	ExpirationEpochSeconds int
}

type OrderSignature struct {
	OrderType                string `json:"order_type"`
	AssetIDSynthetic         string `json:"asset_id_synthetic"`
	AssetIDCollateral        string `json:"asset_id_collateral"`
	AssetIDFee               string `json:"asset_id_fee"`
	QuantumsAmountSynthetic  int64  `json:"quantums_amount_synthetic"`
	QuantumsAmountCollateral int64  `json:"quantums_amount_collateral"`
	QuantumsAmountFee        int64  `json:"quantums_amount_fee"`
	IsBuyingSynthetic        string `json:"is_buying_synthetic"`
	PositionID               string `json:"position_id"`
	Nonce                    string `json:"nonce"`
	ExpirationEpochHours     int    `json:"expiration_epoch_hours"`
}

func NewOrderSignature(req *OrderSignatureRequest) (*OrderSignature, error) {
	syntheticAsset := constants.SYNTHETIC_ASSET_MAP[req.Market]
	syntheticAssetID := constants.SYNTHETIC_ASSET_ID_MAP[syntheticAsset]
	collateralAssetID := constants.COLLATERAL_ASSET_ID_BY_NETWORK_ID[req.NetworkID]

	price, _, err := new(big.Float).Parse(req.Price, 10)
	if err != nil {
		return nil, fmt.Errorf("cannot parse req.Price: %w", err)
	}
	size, _, err := new(big.Float).Parse(req.Size, 10)
	if err != nil {
		return nil, fmt.Errorf("cannot parse req.Size: %w", err)
	}

	quantumAmountSynthetic, err := ToQuantumExact(price, syntheticAsset)
	if err != nil {
		return nil, err
	}

	// TODO: Check this works for all cases
	price = price.Mul(price, size)

	var quantumAmountCollateral int64
	isBuyingSynthetic := req.Side == constants.ORDER_SIDE_BUY
	if isBuyingSynthetic {
		quantumAmountCollateral = ToQuantumRoundUp(price, constants.COLLATERAL_ASSET)
	} else {
		quantumAmountCollateral = ToQuantumRoundDown(price, constants.COLLATERAL_ASSET)
	}

	// The limitFee is a fraction, e.g. 0.01 is a 1 % fee.
	// It is always paid in the collateral asset.
	// Constrain the limit fee to six decimals of precision.
	// The final fee amount must be rounded up.
	limitFee, _, err := new(big.Float).Parse(req.LimitFee, 10)
	if err != nil {
		return nil, fmt.Errorf("cannot parse req.LimitFee: %w", err)
	}
	// TODO: Check that this works as expected
	limitFee = limitFee.SetPrec(6)
	qaf := new(big.Float).Mul(new(big.Float).SetInt64(quantumAmountCollateral), limitFee)
	quantumAmountFee, _ := qaf.Int64()
	if !qaf.IsInt() {
		// Round up if needed
		quantumAmountFee++
	}

	// Orders may have a short time-to-live on the orderbook, but we need
	// to ensure their signatures are valid by the time they reach the
	// blockchain. Therefore, we enforce that the signed expiration includes
	// a buffer relative to the expiration timestamp sent to the dYdX API.
	expirationEpochHours := int(math.Ceil(
		float64(req.ExpirationEpochSeconds)/float64(ONE_HOUR_IN_SECONDS),
	)) + ORDER_SIGNATURE_EXPIRATION_BUFFER_HOURS

	return &OrderSignature{
		OrderType:         "LIMIT_ORDER_WITH_FEES",
		AssetIDSynthetic:  syntheticAssetID,
		AssetIDCollateral: collateralAssetID,
		// TODO: Copied from the Python SDK.
		// Looks like a limitation in the dYdX server?
		AssetIDFee:               collateralAssetID,
		QuantumsAmountSynthetic:  quantumAmountSynthetic,
		QuantumsAmountCollateral: quantumAmountCollateral,
		QuantumsAmountFee:        quantumAmountFee,
		IsBuyingSynthetic:        strconv.FormatBool(isBuyingSynthetic),
		PositionID:               req.PositionID,
		Nonce:                    NonceFromClientID(req.ClientID),
		ExpirationEpochHours:     expirationEpochHours,
	}, nil
}
