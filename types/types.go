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
	// TODO: FILLME
}
