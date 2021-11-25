package public

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/tselementes/dydx-v3-go/types"
)

type Client struct {
	host   string
	client *http.Client
}

func New(host string, timeout time.Duration) (*Client, error) {
	_, err := url.Parse(host)
	if err != nil {
		return nil, err
	}
	return &Client{
		host: host,
		client: &http.Client{
			Timeout: timeout,
		},
	}, nil
}

func (c Client) get(reqPath string, urlParams map[string]string) (*http.Response, error) {
	// build the request
	host, err := url.Parse(c.host)
	if err != nil {
		return nil, err
	}
	host.Path = reqPath

	if len(urlParams) > 0 {
		q := host.Query()
		for k, v := range urlParams {
			q.Set(k, v)
		}
		host.RawQuery = q.Encode()
	}

	req, err := http.NewRequest(http.MethodGet, host.String(), nil)
	if err != nil {
		return nil, err
	}

	// execute the request
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Note that put will not handle HTTP errors.
func (c Client) put(reqPath string, data io.Reader) (*http.Response, error) {
	// build the request
	host, err := url.Parse(c.host)
	if err != nil {
		return nil, err
	}
	host.Path = reqPath

	req, err := http.NewRequest(http.MethodPut, host.String(), data)
	if err != nil {
		return nil, err
	}

	// execute the request
	return c.client.Do(req)
}

// UserExists checks whether the provided Ethereum address
// has been onboarded as a user.
func (c Client) UserExists(ethereumAddress string) (bool, error) {
	path := "/v3/users/exists"
	params := map[string]string{
		"ethereumAddress": ethereumAddress,
	}

	resp, err := c.get(path, params)
	if err != nil {
		return false, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	u := struct {
		Exists bool `json:"exists"`
	}{}
	if err := json.Unmarshal(body, &u); err != nil {
		return false, err
	}

	return u.Exists, nil
}

// UsernameExists checks whether the provided username exists
func (c Client) UsernameExists(username string) (bool, error) {
	path := "/v3/usernames"
	params := map[string]string{
		"username": username,
	}

	resp, err := c.get(path, params)
	if err != nil {
		return false, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	u := struct {
		Exists bool `json:"exists"`
	}{}
	if err := json.Unmarshal(body, &u); err != nil {
		return false, err
	}

	return u.Exists, nil
}

// GetMarkets fetches information about all available markets if an empty string
// is provided or information about the specific market is one is specified.
func (c Client) GetMarkets(market *string) (map[string]types.Market, error) {
	path := "/v3/markets"
	var params map[string]string
	if market != nil && *market != "" {
		params = map[string]string{
			"market": *market,
		}
	}

	resp, err := c.get(path, params)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	m := struct {
		Markets map[string]types.Market `json:"markets"`
	}{}
	if err := json.Unmarshal(body, &m); err != nil {
		return nil, err
	}

	return m.Markets, nil
}

// GetOrderbook fetches the orderbook for a market
func (c Client) GetOrderbook(market string) (*types.Orderbook, error) {
	path := "/v3/orderbook/" + market

	resp, err := c.get(path, nil)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	orderbook := &types.Orderbook{}
	if err := json.Unmarshal(body, orderbook); err != nil {
		return nil, err
	}

	return orderbook, nil
}

// GetStats fetches one or more day statistics for a market.
// days is an optional day range for the statistics to have been
// compiled over. Can be one of 1, 7, 30. Defaults to 1.
func (c Client) GetStats(market *string, days *int32) (*types.MarketStats, error) {
	path := "/v3/stats"
	if market != nil && *market != "" {
		path += "/" + *market
	}
	var params map[string]string
	if days != nil {
		params = map[string]string{
			"days": fmt.Sprint(*days),
		}
	}

	resp, err := c.get(path, params)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	stats := &types.MarketStats{}
	if err := json.Unmarshal(body, stats); err != nil {
		return nil, err
	}

	return stats, nil
}

// GetTrades fetches the trades for a market
// startingBeforeOrAt is optional and should be of 2021-09-05T17:33:43.163Z format.
// Trades will include information for all users and as such
// includes less information on individual transactions than the fills endpoint.
// TODO: Use limit - (Optional): The number of candles to fetch (Max 100).
func (c Client) GetTrades(market, startingBeforeOrAt, limit string) ([]types.Trade, error) {
	path := "/v3/trades/" + market
	var params map[string]string
	if startingBeforeOrAt != "" {
		params = map[string]string{
			"startingBeforeOrAt": startingBeforeOrAt,
		}
	}

	resp, err := c.get(path, params)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	t := struct {
		Trades []types.Trade `json:"trades"`
	}{}
	if err := json.Unmarshal(body, &t); err != nil {
		return nil, err
	}

	return t.Trades, nil
}

// GetHistoricalFunding fetches the historical funding for a market
func (c Client) GetHistoricalFunding(market string, effectiveBeforeOrAt *string) ([]types.HistoricalFunding, error) {
	path := "/v3/historical-funding/" + market
	var params map[string]string
	if effectiveBeforeOrAt != nil {
		params = map[string]string{
			"effectiveBeforeOrAt": *effectiveBeforeOrAt,
		}
	}

	resp, err := c.get(path, params)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	hf := struct {
		HistoricalFunding []types.HistoricalFunding `json:"historicalFunding"`
	}{}
	if err := json.Unmarshal(body, &hf); err != nil {
		return nil, err
	}

	return hf.HistoricalFunding, nil
}

// GetFastWithdrawal fetches all fast withdrawal account information.
// Returns a map of all LP provider accounts that have available funds
// for fast withdrawals. Given a debitAmount and asset the user wants
// sent to L1, this endpoint also returns the predicted amount of the
// desired asset the user will be credited on L1. Given a creditAmount
// and asset the user wants sent to L1, this endpoint also returns the
// predicted amount the user will be debited on L2.
// TODO: Use amounts if provided
func (c Client) GetFastWithdrawal(creditAsset, creditAmount, debitAmount *string) (map[string]types.LiquidityProvider, error) {
	path := "/v3/fast-withdrawals"

	resp, err := c.get(path, nil)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	lp := struct {
		LiquidityProviders map[string]types.LiquidityProvider `json:"liquidityProviders"`
	}{}
	if err := json.Unmarshal(body, &lp); err != nil {
		return nil, err
	}

	return lp.LiquidityProviders, nil
}

func (c Client) GetCandles(market string, resolution, fromISO, toISO, limit *string) ([]types.Candle, error) {
	path := "/v3/candles/" + market
	params := make(map[string]string)
	if resolution != nil {
		params["resolution"] = *resolution
	}
	if fromISO != nil {
		params["fromISO"] = *fromISO
	}
	if toISO != nil {
		params["toISO"] = *toISO
	}
	if limit != nil {
		params["limit"] = *limit
	}

	resp, err := c.get(path, params)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	cs := struct {
		Candles []types.Candle `json:"candles"`
	}{}
	if err := json.Unmarshal(body, &cs); err != nil {
		return nil, err
	}

	return cs.Candles, nil
}

func (c Client) GetTime() (*types.Time, error) {
	path := "/v3/time"

	resp, err := c.get(path, nil)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	t := &types.Time{}
	if err := json.Unmarshal(body, t); err != nil {
		return nil, err
	}

	return t, nil
}

// VerifyEmail verifies an email address by providing the verification
// token sent to the email address.
func (c Client) VerifyEmail(token string) error {
	path := "/v3/emails/verify-email"

	t := struct {
		Token string `json:"token"`
	}{
		Token: token,
	}

	body, err := json.Marshal(t)
	if err != nil {
		return err
	}

	resp, err := c.put(path, bytes.NewReader(body))
	if err != nil {
		return err
	}
	if resp.StatusCode != 204 {
		return fmt.Errorf("invalid response status for email verification: %s (%d)", resp.Status, resp.StatusCode)
	}

	return nil
}

// GetPublicRetroactiveMiningRewards gets the retroactive mining rewards for
// an ethereum address.
func (c Client) GetPublicRetroactiveMiningRewards(ethereumAddress string) (*types.PublicRetroactiveMiningReward, error) {
	path := "/v3/rewards/public-retroactive-mining"
	params := map[string]string{
		"ethereumAddress": ethereumAddress,
	}

	resp, err := c.get(path, params)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	r := &types.PublicRetroactiveMiningReward{}
	if err := json.Unmarshal(body, r); err != nil {
		return nil, err
	}

	return r, nil
}

// Get global config variables for the exchange as a whole.
// This includes (but is not limited to) details on the exchange,
// including addresses, fees, transfers, and rate limits.
func (c Client) GetConfig() (*types.Config, error) {
	path := "/v3/config"

	resp, err := c.get(path, nil)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	config := &types.Config{}
	if err := json.Unmarshal(body, config); err != nil {
		return nil, err
	}

	return config, nil
}
