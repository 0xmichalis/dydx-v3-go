package client

import (
	"context"
	"crypto/ecdsa"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/tselementes/dydx-v3-go/public"
)

type Client struct {
	host    string
	chainId int

	ethClient *ethclient.Client
	pubClient *public.Client
}

func New(
	host string,
	timeout time.Duration,
	defaultEthereumAddress common.Address,
	ethPrivateKey *ecdsa.PrivateKey,
	chainId int,
	starkPublicKey string,
	starkPrivateKey string,
	starkPrivateKeyYCoordinate string,
	providerURL string,
	apiKeyCredentials string,
) (*Client, error) {
	// TODO: Create separate constructor to accept context
	c, err := ethclient.DialContext(context.Background(), providerURL)
	if err != nil {
		return nil, err
	}

	return &Client{
		host:    host,
		chainId: chainId,

		ethClient: c,
		pubClient: public.New(host, timeout),
	}, nil
}
