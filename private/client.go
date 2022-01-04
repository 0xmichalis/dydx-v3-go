package private

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/tselementes/dydx-v3-go/types"
)

const (
	// DYDX API credentials keys
	Key        = "key"
	Passphrase = "passphrase"
	Secret     = "secret"
)

type Client struct {
	host              string
	client            *http.Client
	networkId         int
	starkPrivateKey   string
	defaultAddress    common.Address
	apiKeyCredentials map[string]string
}

func New(
	host string,
	timeout time.Duration,
	networkId int,
	starkPrivateKey string,
	defaultAddress common.Address,
	apiKeyCredentials map[string]string,
) (*Client, error) {
	_, err := url.Parse(host)
	if err != nil {
		return nil, err
	}
	// TODO: Validate that apiKeyCredentials has a
	// key, passphrase, and secret
	return &Client{
		host: host,
		client: &http.Client{
			Timeout: timeout,
		},
		networkId:         networkId,
		starkPrivateKey:   starkPrivateKey,
		defaultAddress:    defaultAddress,
		apiKeyCredentials: apiKeyCredentials,
	}, nil
}

func (c Client) doRequest(method, path string, urlParams map[string]string, data []byte) (*http.Response, error) {
	host, err := url.Parse(c.host)
	if err != nil {
		return nil, fmt.Errorf("failed to parse host (%s): %w", c.host, err)
	}
	host.Path = "/v3/" + path

	if len(urlParams) > 0 {
		q := host.Query()
		for k, v := range urlParams {
			q.Set(k, v)
		}
		host.RawQuery = q.Encode()
	}

	req, err := http.NewRequest(method, host.String(), bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("failed to build new request: %w", err)
	}

	now := time.Now().UTC().Format(time.RFC3339)
	signature, err := c.sign(method, host.Path, now, data)
	if err != nil {
		return nil, fmt.Errorf("failed to sign request: %w", err)
	}

	req.Header.Set("DYDX-SIGNATURE", signature)
	req.Header.Set("DYDX-API-KEY", c.apiKeyCredentials[Key])
	req.Header.Set("DYDX-TIMESTAMP", now)
	req.Header.Set("DYDX-PASSPHRASE", c.apiKeyCredentials[Passphrase])

	// execute the request
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to %s %s: %w", method, host.Path, err)
	}
	return resp, nil
}

func (c Client) sign(method, path, timestamp string, data []byte) (string, error) {
	var dataJSON string
	var err error

	if len(data) > 0 {
		dataJSON, err = jsonStringifyWithoutNils(data)
		if err != nil {
			return "", fmt.Errorf("cannot stringify JSON: %w", err)
		}
	}

	message := timestamp +
		method +
		path +
		dataJSON

	s, err := base64.URLEncoding.DecodeString(c.apiKeyCredentials[Secret])
	if err != nil {
		return "", fmt.Errorf("cannot decode from base64: %w", err)
	}
	h := hmac.New(sha256.New, s)
	if _, err := h.Write([]byte(message)); err != nil {
		return "", fmt.Errorf("cannot hash to buffer: %w", err)
	}
	digest := h.Sum(nil)
	return base64.URLEncoding.EncodeToString(digest), nil
}

func jsonStringifyWithoutNils(data []byte) (string, error) {
	deserialized := map[string]interface{}{}
	if err := json.Unmarshal(data, &deserialized); err != nil {
		return "", fmt.Errorf("cannot unmarshal to drop nils: %w", err)
	}
	return jsonStringify(removeNils(deserialized))
}

func jsonStringify(data map[string]interface{}) (string, error) {
	out, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("cannot marshal to JSON: %w", err)
	}
	return string(out), nil
}

func removeNils(initialMap map[string]interface{}) map[string]interface{} {
	withoutNils := map[string]interface{}{}
	for key, value := range initialMap {
		_, ok := value.(map[string]interface{})
		if ok {
			value = removeNils(value.(map[string]interface{}))
			withoutNils[key] = value
			continue
		}
		if value != nil {
			withoutNils[key] = value
		}
	}
	return withoutNils
}

// Does not handle HTTP errors.
func (c Client) get(path string, urlParams map[string]string) ([]byte, error) {
	resp, err := c.doRequest(http.MethodGet, path, urlParams, nil)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected response status: %v", resp.Status)
	}
	return ioutil.ReadAll(resp.Body)
}

// Does not handle HTTP errors.
func (c Client) post(path string, data []byte) (*http.Response, error) {
	return c.doRequest(http.MethodPost, path, nil, data)
}

// Does not handle HTTP errors.
func (c Client) put(path string, data []byte) (*http.Response, error) {
	return c.doRequest(http.MethodPut, path, nil, data)
}

// Does not handle HTTP errors.
func (c Client) delete(path string, urlParams map[string]string) (*http.Response, error) {
	return c.doRequest(http.MethodDelete, path, urlParams, nil)
}

// GetApiKeys fetches all api keys associated with an Ethereum address.
func (c Client) GetApiKeys() ([]types.ApiKey, error) {
	data, err := c.get("api-keys", nil)
	if err != nil {
		return nil, err
	}
	resp := &types.GetApiKeysResponse{}
	if err := json.Unmarshal(data, resp); err != nil {
		return nil, err
	}
	return resp.ApiKeys, nil
}

// GetRegistration fetches the dYdX provided Ethereum signature required to
// send a registration transaction to the Starkware smart contract.
func (c Client) GetRegistration() (*types.Registration, error) {
	data, err := c.get("registration", nil)
	if err != nil {
		return nil, err
	}
	registration := types.Registration{}
	if err := json.Unmarshal(data, &registration); err != nil {
		return nil, err
	}
	return &registration, nil
}

// GetUser fetches user information.
func (c Client) GetUser() (*types.User, error) {
	data, err := c.get("users", nil)
	if err != nil {
		return nil, err
	}
	resp := &types.UserResponse{}
	if err := json.Unmarshal(data, resp); err != nil {
		return nil, err
	}
	return resp.User, nil
}

// UpdateUser updates user information and return the updated user.
func (c Client) UpdateUser(req *types.UpdateUserRequest) (*types.User, error) {
	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	resp, err := c.put("users", data)
	if err != nil {
		return nil, err
	}
	// TODO: Check response status code
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	uResp := &types.UserResponse{}
	if err := json.Unmarshal(body, uResp); err != nil {
		return nil, err
	}
	return uResp.User, nil
}

// GetAccount fetches ethereumAddress or if ethereumAddress is nil, it will
// default to defaultAddress which is the default address the Client was
// initialized with.
func (c Client) GetAccount(ethereumAddress *common.Address) (*types.Account, error) {
	address := c.defaultAddress
	if ethereumAddress != nil {
		address = *ethereumAddress
	}
	// TODO: Need to check address format
	data, err := c.get(fmt.Sprintf("accounts/%x", address), nil)
	if err != nil {
		return nil, err
	}
	resp := &types.GetAccountResponse{}
	if err := json.Unmarshal(data, resp); err != nil {
		return nil, err
	}
	return resp.Account, nil
}

// GetAccounts fetches all accounts for a user.
func (c Client) GetAccounts() ([]*types.Account, error) {
	data, err := c.get("accounts", nil)
	if err != nil {
		return nil, err
	}
	resp := &types.GetAccountsResponse{}
	if err := json.Unmarshal(data, resp); err != nil {
		return nil, err
	}
	return resp.Accounts, nil
}

// GetPositions fetches all user positions. Filters can be provided via
// GetPositionsFilter or pass nil to fetch all positions.
func (c Client) GetPositions(filters *types.GetPositionsFilter) ([]*types.Position, error) {
	params := make(map[string]string)
	if filters != nil {
		if filters.Market != nil {
			params["market"] = *filters.Market
		}
		if filters.Limit != nil {
			params["limit"] = *filters.Limit
		}
		if filters.Status != nil {
			params["status"] = *filters.Status
		}
		if filters.CreatedBeforeOrAt != nil {
			params["createdBeforeOrAt"] = *filters.CreatedBeforeOrAt
		}
	}
	data, err := c.get("positions", params)
	if err != nil {
		return nil, err
	}
	resp := &types.GetPositionsResponse{}
	if err := json.Unmarshal(data, resp); err != nil {
		return nil, err
	}
	return resp.Positions, nil
}

// GetOrders fetches active (not filled or canceled) orders for a user by specified parameters.
func (c Client) GetOrders(filters *types.GetOrdersFilter) ([]*types.Order, error) {
	params := make(map[string]string)
	if filters != nil {
		if filters.Market != nil {
			params["market"] = *filters.Market
		}
		if filters.Limit != nil {
			params["limit"] = *filters.Limit
		}
		if filters.Status != nil {
			params["status"] = *filters.Status
		}
		if filters.Side != nil {
			params["side"] = *filters.Side
		}
		if filters.Type != nil {
			params["type"] = *filters.Type
		}
		if filters.CreatedBeforeOrAt != nil {
			params["createdBeforeOrAt"] = *filters.CreatedBeforeOrAt
		}
		if filters.ReturnLatestOrders != nil {
			params["returnLatestOrders"] = strconv.FormatBool(*filters.ReturnLatestOrders)
		}
	}
	data, err := c.get("orders", params)
	if err != nil {
		return nil, err
	}
	resp := &types.GetOrdersResponse{}
	if err := json.Unmarshal(data, resp); err != nil {
		return nil, err
	}
	return resp.Orders, nil
}

// GetOrderByID fetches an order by its id
func (c Client) GetOrderByID(id string) (*types.Order, error) {
	data, err := c.get(fmt.Sprintf("orders/%s", id), nil)
	if err != nil {
		return nil, err
	}
	resp := &types.GetOrderByIdResponse{}
	if err := json.Unmarshal(data, resp); err != nil {
		return nil, err
	}
	return resp.Order, nil
}

// GetOrderByClientID fetches an order by its client id
func (c Client) GetOrderByClientID(id string) (*types.Order, error) {
	data, err := c.get(fmt.Sprintf("orders/client/%s", id), nil)
	if err != nil {
		return nil, err
	}
	resp := &types.GetOrderByIdResponse{}
	if err := json.Unmarshal(data, resp); err != nil {
		return nil, err
	}
	return resp.Order, nil
}

func (c Client) CreateOrder(req *types.OrderRequest) (*types.Order, error) {
	if req.Signature == "" && c.starkPrivateKey == "" {
		return nil, errors.New("No signature provided and client was not initialized with stark private key")
	}

	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	resp, err := c.post("orders", data)
	if err != nil {
		return nil, err
	}

	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	o := &types.GetOrderByIdResponse{}
	if err := json.Unmarshal(respData, o); err != nil {
		return nil, err
	}

	return o.Order, nil
}
