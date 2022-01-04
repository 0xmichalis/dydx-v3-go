package starkex

import (
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

type OrderSignature struct{}

func New(req *OrderSignatureRequest) (*OrderSignature, error) {
	syntheticAsset := constants.SYNTHETIC_ASSET_MAP[req.Market]
	syntheticAssetID := constants.SYNTHETIC_ASSET_ID_MAP[syntheticAsset]
	_ = syntheticAssetID

	collateralAssetID := constants.COLLATERAL_ASSET_ID_BY_NETWORK_ID[req.NetworkID]
	_ = collateralAssetID

	isBuyingSynthetic := req.Side == constants.ORDER_SIDE_BUY
	_ = isBuyingSynthetic

	return nil, nil
}
