package types

type Market struct {
	// Symbol of the market.
	Market string `json:"market"`
	// Status of the market. Can be one of ONLINE, OFFLINE, POST_ONLY or CANCEL_ONLY.
	Status string `json:"status"`
	// Symbol of the base asset. e.g. "BTC".
	BaseAsset string `json:"baseAsset"`
	// Symbol of the quote asset. e.g. "BTC".
	QuoteAsset string `json:"quoteAsset"`
	// The minimum step size (in base currency) of trade sizes for the market.
	StepSize string `json:"stepSize"`
	// The Tick size of the market.
	TickSize string `json:"tickSize"`
	// The current index price of the market.
	IndexPrice string `json:"indexPrice"`
	// The current oracle price of the market.
	OraclePrice string `json:"oraclePrice"`
	// The absolute price change of the index price over the past 24 hours.
	PriceChange string `json:"priceChange24H"`
	// The predicted next funding rate (as a 1-hour rate). Can be up to 5 seconds delayed.
	NextFundingRate string `json:"nextFundingRate"`
	// The timestamp of the next funding update.
	NextFundingAt string `json:"nextFundingAt"`
	// Minimum order size for the market.
	MinOrderSize string `json:"minOrderSize"`
	// Type of the market. This will always be PERPETUAL for now.
	Type string `json:"type"`
	// The margin fraction needed to open a position.
	InitialMarginFraction string `json:"initialMarginFraction"`
	// The margin fraction required to prevent liquidation.
	MaintenanceMarginFraction string `json:"maintenanceMarginFraction"`
	// The max position size (in base token) before increasing the initial-margin-fraction.
	BaselinePositionSize string `json:"baselinePositionSize"`
	// The step size (in base token) for increasing the initialMarginFraction by (incrementalInitialMarginFraction per step).
	IncrementalPositionSize string `json:"incrementalPositionSize"`
	// The increase of initialMarginFraction for each incrementalPositionSize above the baselinePositionSize the position is.
	IncrementalInitialMarginFraction string `json:"incrementalInitialMarginFraction"`
	// The max position size for this market in base token.
	MaxPositionSize string `json:"maxPositionSize"`
	// The USD volume of the market in the previous 24 hours.
	Volume string `json:"volume24H"`
	// The number of trades in the market in the previous 24 hours.
	Trades string `json:"trades24H"`
	// The open interest in base token.
	OpenInterest string `json:"openInterest"`
	// The asset resolution is the number of quantums (Starkware units) that fit within one "human-readable" unit of the asset.
	AssetResolution string `json:"assetResolution"`
}

type Orderbook struct {
	// Sorted by price in descending order.
	Bids []OrderbookOrder `json:"bids"`
	// Sorted by price in ascending order.
	Asks []OrderbookOrder `json:"asks"`
}

type OrderbookOrder struct {
	// The price of the order (in quote / base currency).
	Price string `json:"price"`
	// The size of the order (in base currency).
	Size string `json:"size"`
}

type MarketStats struct {
	// The symbol of the market, e.g. ETH-USD.
	Market string `json:"market"`
	// The open price of the market.
	Open string `json:"open"`
	// The high price of the market.
	High string `json:"high"`
	// The low price of the market.
	Low string `json:"low"`
	// The close price of the market.
	Close string `json:"close"`
	// The total amount of base asset traded.
	BaseVolume string `json:"baseVolume"`
	// The total amount of quote asset traded.
	QuoteVolume string `json:"quoteVolume"`
	// Type of the market. This will always be PERPETUAL for now.
	Type string `json:"type"`
}

type Trade struct {
	// Either BUY or SELL.
	Side string `json:"side"`
	// The size of the trade.
	Size string `json:"size"`
	// The price of the trade.
	Price string `json:"price"`
	// The time of the trade.
	CreatedAt string `json:"createdAt"`
}

type HistoricalFunding struct {
	// Market for which to query historical funding.
	Market string `json:"market"`
	// The funding rate (as a 1-hour rate).
	Rate string `json:"rate"`
	// Oracle price used to calculate the funding rate.
	Price string `json:"price"`
	// Time at which funding payments were exchanged at this rate.
	EffectiveAt string `json:"effectiveAt"`
}

type LiquidityProvider struct {
	// The funds available for the LP.
	AvailableFunds string `json:"availableFunds"`
	// The public stark key for the LP.
	StarkKey string `json:"starkKey"`
	// The Liquidity Provider Quote given the user's request.
	// Null if no request from the user or the request is unfillable by this LP.
	Quote *LiquidityProviderQuote `json:"quote"`
}

type LiquidityProviderQuote struct {
	// The asset that would be sent to the user on L1.
	CreditAsset string `json:"creditAsset"`
	// The amount of creditAsset that would be sent to the user (human readable).
	CreditAmount string `json:"creditAmount"`
	// The amount of USD that would be deducted from the users L2 account (human readable).
	DebitAmount string `json:"debitAmount"`
}

type Candle struct {
	// When the candle started, time of first trade in candle.
	StartedAt string `json:"startedAt"`
	// When the candle was last updated
	UpdatedAt string `json:"updatedAt"`
	// Market the candle is for.
	Market string `json:"market"`
	// Time-period of candle (currently 1HOUR or 1DAY).
	Resolution string `json:"resolution"`
	// The open price of the candle.
	Open string `json:"open"`
	// The high price of the candle.
	High string `json:"high"`
	// Low trade price of the candle.
	Low string `json:"low"`
	// The close price of the candle.
	Close string `json:"close"`
	// Volume of trade in baseToken currency for the candle.
	BaseTokenVolume string `json:"baseTokenVolume"`
	// Count of trades during the candle.
	Trades string `json:"trades"`
	// Volume of trade in USD for the candle.
	USDVolume string `json:"usdVolume"`
	// The open interest in baseToken at the start of the candle.
	StartingOpenInterest string `json:"startingOpenInterest"`
}

type Time struct {
	// ISO time of the server in UTC.
	ISO string `json:"iso"`
	// Epoch time in seconds with milliseconds.
	Epoch string `json:"epoch"`
}

type PublicRetroactiveMiningReward struct {
	// The number of allocated dYdX tokens for the address.
	Allocation string `json:"allocation"`
	// The addresses' required trade volume (in USD) to be able to claim the allocation.
	TargetVolume string `json:"targetVolume"`
}

type Config struct {
	CollateralAssetId             string                  `json:"collateralAssetId"`
	CollateralTokenAddress        string                  `json:"collateralTokenAddress"`
	DefaultMakerFee               string                  `json:"defaultMakerFee"`
	DefaultTakerFee               string                  `json:"defaultTakerFee"`
	ExchangeAddress               string                  `json:"exchangeAddress"`
	MaxExpectedBatchLengthMinutes string                  `json:"maxExpectedBatchLengthMinutes"`
	MaxFastWithdrawalAmount       string                  `json:"maxFastWithdrawalAmount"`
	CancelOrderRateLimiting       CancelOrderRateLimiting `json:"cancelOrderRateLimiting"`
	PlaceOrderRateLimiting        PlaceOrderRateLimiting  `json:"placeOrderRateLimiting"`
}

type CancelOrderRateLimiting struct {
	MaxPointsMulti  int32 `json:"maxPointsMulti"`
	MaxPointsSingle int32 `json:"maxPointsSingle"`
	WindowSecMulti  int32 `json:"windowSecMulti"`
	WindowSecSingle int32 `json:"windowSecSingle"`
}

type PlaceOrderRateLimiting struct {
	MaxPoints                 int32 `json:"maxPoints"`
	WindowSec                 int32 `json:"windowSec"`
	TargetNotional            int32 `json:"targetNotional"`
	MinLimitConsumption       int32 `json:"minLimitConsumption"`
	MinMarketConsumption      int32 `json:"minMarketConsumption"`
	MinTriggerableConsumption int32 `json:"minTriggerableConsumption"`
	MaxOrderConsumption       int32 `json:"maxOrderConsumption"`
}

type ApiKey string

type GetApiKeysResponse struct {
	ApiKeys []ApiKey `json:"apiKeys"`
}

type Registration struct {
	// Ethereum signature authorizing the user's Ethereum address to register
	// for the corresponding position id.
	Signature string `json:"signature"`
}

type User struct {
	// The 20-byte Ethereum address.
	EthereumAddress string `json:"ethereumAddress"`
	// True if the user is registered on the starkware smart contract. This is false otherwise.
	IsRegistered bool `json:"isRegistered"`
	// Email address.
	Email string `json:"email"`
	// User defined username.
	Username string `json:"username"`
	// The affiliate link that referred this user, or null if the user was not referred.
	ReferredByAffiliateLink *string `json:"referredByAffiliateLink,omitempty"`
	// The fee rate the user would be willing to take as the maker. Note, 1% would be represented as 0.01.
	MakerFeeRate string `json:"makerFeeRate"`
	// The fee rate the user would be willing to take as the taker. Note, 1% would be represented as 0.01.
	TakerFeeRate string `json:"takerFeeRate"`
	// The user's thirty day maker volume. Note, this is in USD (eg $12.34 -> 12.34).
	MakerVolume string `json:"makerVolume30D"`
	// The user's thirty day maker volume. Note, this is in USD (eg $12.34 -> 12.34).
	TakerVolume string `json:"takerVolume30D"`
	// The user's thirty day fees. Note, this is in USD (eg $12.34 -> 12.34).
	Fees string `json:"fees30D"`
	// The user's unstructured user data.
	UserData UnstructuredData `json:"userData"`
	// The user's DYDX token holdings.
	DydxTokenBalance string `json:"dydxTokenBalance"`
	// The user's staked DYDX token holdings
	StakedDydxTokenBalance string `json:"stakedDydxTokenBalance"`
	// If the user's email address is verified to receive emails from dYdX.
	IsEmailVerified bool `json:"isEmailVerified"`
}

type UnstructuredData map[string]interface{}

type UserResponse struct {
	User *User `json:"user"`
}

type UpdateUserRequest struct {
	// User metadata in a map. This is serialized
	// into a JSON blob.
	UserData UnstructuredData `json:"userData,omitempty"`
	// Email to be used with the user.
	Email string `json:"email,omitempty"`
	// Username to be used for the user.
	Username string `json:"username,omitempty"`
	// Share username publically on leaderboard rankings.
	IsSharingUsername string `json:"isSharingUsername,omitempty"`
	// Share ETH address publically on leaderboard rankings.
	IsSharingAddress string `json:"isSharingAddress,omitempty"`
	// Country of the user's residence. Must be ISO 3166-1 Alpha-2 compliant.
	Country string `json:"country,omitempty"`
}

type GetAccountResponse struct {
	Account *Account `json:"account"`
}

type Account struct {
	// Public StarkKey associated with an account.
	StarkKey string `json:"starkKey"`
	// Starkware-specific positionId.
	PositionId string `json:"positionId"`
	// The amount of equity (value) in the account. Uses balances and oracle-prices to calculate.
	Equity string `json:"equity"`
	// The amount of collateral that is withdrawable from the account.
	FreeCollateral string `json:"freeCollateral"`
	// Human readable quote token balance. Can be negative.
	QuoteBalance string `json:"quoteBalance"`
	// The sum amount of all pending deposits.
	PendingDeposits string `json:"pendingDeposits"`
	// The sum amount of all pending withdrawal requests.
	PendingWithdrawals string `json:"pendingWithdrawals"`
	// When the account was first created in UTC.
	CreatedAt string `json:"createdAt"`
	// Markets where the user has no position are not returned in the map.
	OpenPositions Positions `json:"openPositions"`
	// Unique accountNumber for the account.
	AccountNumber string `json:"accountNumber"`
	// Unique id of the account hashed from the userId and the accountNumber.
	ID string `json:"id"`
}

type Positions map[string]Position

type Position struct {
	// The market of the position.
	Market string `json:"market"`
	// The status of the position.
	Status string `json:"status"`
	// The side of the position. LONG or SHORT.
	Side string `json:"side"`
	// The current size of the position. Positive if long, negative if short, 0 if closed.
	Size string `json:"size"`
	// The maximum (absolute value) size of the position. Positive if long, negative if short.
	MaxSize string `json:"maxSize"`
	// Average price paid to enter the position.
	EntryPrice string `json:"entryPrice"`
	// Average price paid to exit the position.
	ExitPrice *string `json:"exitPrice,omitempty"`
	// The unrealized pnl of the position in quote currency using the market's index-price
	// (https://docs.dydx.exchange/#index-prices) for the position to calculate.
	UnrealizedPNL string `json:"unrealizedPnl"`
	// The realized pnl of the position in quote currency.
	RealizedPNL string `json:"realizedPnl"`
	// Timestamp of when the position was opened.
	CreatedAt string `json:"createdAt"`
	// Timestamp of when the position was closed.
	ClosedAt *string `json:"closedAt,omitempty"`
	// Sum of all funding payments for this position.
	NetFunding string `json:"netFunding"`
	// Sum of all trades sizes that increased the size of this position.
	SumOpen string `json:"sumOpen"`
	// Sum of all trades sizes that decreased the size of this position.
	SumClose string `json:"sumClose"`
}

type GetAccountsResponse struct {
	Accounts []*Account `json:"accounts"`
}

type GetPositionsFilter struct {
	// Market of the position.
	Market *string `json:"market,omitempty"`
	// Status of the position. Can be OPEN, CLOSED or LIQUIDATED.
	Status *string `json:"status,omitempty"`
	// The maximum number of positions that can be fetched via this request.
	// Note, this cannot be greater than 100.
	Limit *string `json:"limit,omitempty"`
	// Set a date by which the positions had to be created.
	CreatedBeforeOrAt *string `json:"createdBeforeOrAt,omitempty"`
}

type GetPositionsResponse struct {
	Positions []*Position `json:"positions"`
}
