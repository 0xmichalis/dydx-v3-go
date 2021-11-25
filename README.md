# dydx-v3-go
DYDX v3 Golang client

## Get tokens in Ropsten

1. Onboard as a user in https://trade.stage.dydx.exchange
2. Once onboarded, open the browser console, navigate to `Storage > Local Storage`, find the `API_KEY_PAIRS` value and copy `key`, `secret`, and `passphrase`.
3. [TODO: Run from our own SDK] Install the Python SDK and in the following code stanza replace the credentials retrieved from step. 2
```python
from dydx3 import Client
from dydx3 import private_key_to_public_key_pair_hex
from dydx3.constants import API_HOST_ROPSTEN
from dydx3.constants import NETWORK_ID_ROPSTEN
from web3 import Web3

ETHEREUM_ADDRESS = '<ETHEREUM_ADDRESS>'

WEB_PROVIDER_URL = 'https://ropsten.infura.io/v3/<INFURA_API_KEY>'

client = Client(
    network_id=NETWORK_ID_ROPSTEN,
    host=API_HOST_ROPSTEN,
    default_ethereum_address=ETHEREUM_ADDRESS,
    web3=Web3(Web3.HTTPProvider(WEB_PROVIDER_URL)),
    api_key_credentials={
        'key': '<key>',
        'secret': '<secret>',
        'passphrase': '<passphrase>',
    },
)

accounts_response = client.private.request_testnet_tokens()
print(accounts_response)
```
4. After executing the script above, you should get USDC directly in the DYDX app.
