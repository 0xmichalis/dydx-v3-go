package private

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

const (
	// DYDX API credentials keys
	key        = "key"
	passphrase = "passphrase"
	secret     = "secret"
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
		return nil, err
	}
	host.Path = path

	if len(urlParams) > 0 {
		q := host.Query()
		for k, v := range urlParams {
			q.Set(k, v)
		}
		host.RawQuery = q.Encode()
	}

	req, err := http.NewRequest(method, host.String(), bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	now := time.Now().Format(time.RFC3339)
	signature, err := c.sign(method, path, now, data)
	if err != nil {
		return nil, err
	}

	req.Header.Set("DYDX-SIGNATURE", signature)
	req.Header.Set("DYDX-API-KEY", c.apiKeyCredentials[key])
	req.Header.Set("DYDX-TIMESTAMP", now)
	req.Header.Set("DYDX-PASSPHRASE", c.apiKeyCredentials[passphrase])

	// execute the request
	return c.client.Do(req)
}

func (c Client) sign(method, path, timestamp string, data []byte) (string, error) {
	dataJSON, err := jsonStringify(data)
	if err != nil {
		return "", err
	}
	message := timestamp +
		strings.ToUpper(method) +
		path +
		dataJSON

	s, err := base64.StdEncoding.DecodeString(c.apiKeyCredentials[secret])
	if err != nil {
		return "", err
	}
	h := hmac.New(sha256.New, s)
	if _, err := h.Write([]byte(message)); err != nil {
		return "", err
	}
	digest := h.Sum(nil)
	return base64.StdEncoding.EncodeToString(digest), nil
}

func jsonStringify(data []byte) (string, error) {
	deserialized := map[string]interface{}{}
	if err := json.Unmarshal(data, &deserialized); err != nil {
		return "", err
	}
	out, err := json.Marshal(removeNils(deserialized))
	if err != nil {
		return "", err
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
func (c Client) get(path string, urlParams map[string]string) (*http.Response, error) {
	return c.doRequest(http.MethodGet, path, urlParams, nil)
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
