package client

import (
	"context"
	"crypto/ecdsa"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/tselementes/dydx-v3-go/private"
	"github.com/tselementes/dydx-v3-go/public"
)

type Client struct {
	host    string
	chainId int

	ethClient  *ethclient.Client
	pubClient  *public.Client
	privClient *private.Client
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
	apiKeyCredentials map[string]string,
) (*Client, error) {
	// TODO: Create separate constructor to accept context
	ethClient, err := ethclient.DialContext(context.Background(), providerURL)
	if err != nil {
		return nil, err
	}

	pubClient, err := public.New(host, timeout)
	if err != nil {
		return nil, err
	}

	privClient, err := private.New(
		host,
		timeout,
		chainId,
		starkPrivateKey,
		defaultEthereumAddress,
		apiKeyCredentials,
	)
	if err != nil {
		return nil, err
	}

	return &Client{
		host:    host,
		chainId: chainId,

		ethClient:  ethClient,
		pubClient:  pubClient,
		privClient: privClient,
	}, nil
}
