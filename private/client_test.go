package private_test

// import (
// 	"testing"
// 	"time"

// 	"github.com/ethereum/go-ethereum/common"

// 	client "github.com/tselementes/dydx-v3-go"
// 	"github.com/tselementes/dydx-v3-go/private"
// )

// func TestClient(t *testing.T) {
// 	c, err := private.New(
// 		client.API_HOST_ROPSTEN,
// 		10*time.Second,
// 		client.NETWORK_ID_ROPSTEN,
// 		"00035a53c9762e16e13afdd0da9b2605621a5ff1c5bf1d1b3f1f1482f1cf92fa",
// 		common.HexToAddress("0xBdC85027BCDBe20B3430523a773bf3008888FA9d"),
// 		map[string]string{
// 			private.Key:        "d94fdca8-d2d4-bba3-8e4b-f26368c13fa8",
// 			private.Passphrase: "_vtaQ2udhTD7rhHj-VAn",
// 			private.Secret:     "AbxObiZJ5TOaMEmQcto4544U_WXIsacGD90TVBv9",
// 		},
// 	)
// 	if err != nil {
// 		t.Fatalf("failed to initialize client: %v", err)
// 	}

// 	user, err := c.GetUser()
// 	if err != nil {
// 		t.Fatalf("cannot get user: %v", err)
// 	}
// 	t.Logf("got user %#v", *user)
// 	t.Fatal("on purpose")
// }
