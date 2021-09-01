# gosignit

## Prerequisites:
Install TrustWallet
```
cd lib
git clone https://github.com/trustwallet/wallet-core trustwallet
cd trustwallet
./bootstrap.sh
```

## API Reference
### Ping
#### Request:
```http request
GET /ping
```
#### Response:
```json
{
  "ping": "pong"
}
```
### Get supported coins
#### Request:
```http request
GET /api/v1/coins
```
#### Response:
```json
[
  {
    "id": "bitcoin",
    "name": "Bitcoin",
    "derivation_path": "m/84'/0'/0'/0/0",
    "symbol": "BTC",
    "decimals": 8,
    "explorer_transaction_url": "https://blockchair.com/bitcoin/transaction/",
    "explorer_account_url": "https://blockchair.com/bitcoin/address/"
  },
  {
    "id": "litecoin",
    "name": "Litecoin",
    "derivation_path": "m/84'/2'/0'/0/0",
    "symbol": "LTC",
    "decimals": 8,
    "explorer_transaction_url": "https://blockchair.com/litecoin/transaction/",
    "explorer_account_url": "https://blockchair.com/litecoin/address/"
  },
  {
    "id": "doge",
    "name": "Dogecoin",
    "derivation_path": "m/44'/3'/0'/0/0",
    "symbol": "DOGE",
    "decimals": 8,
    "explorer_transaction_url": "https://blockchair.com/dogecoin/transaction/",
    "explorer_account_url": "https://blockchair.com/dogecoin/address/"
  },
  <...>
]
```
### Sign Transactions
### Bitcoin blockchain
#### Request:
Use certain coin `id` from the list above as `coinId` path parameter.
```http request
GET /api/v1/{coinId}/sign
Content-Type: application/json

{
  "amount": 1020304,
  "byte_fee": 1,
  "to_address": "1Bp9U1ogV3A14FMvKbRJms7ctyso4Z4Tcx",
  "change_address": "1FQc5LdgGHMHEN9nwkjmz6tWkxhPpxBvBU",
  "use_max_amount": false,
  "utxo": [
    {
      "hash": "fff7f7881a8099afa6940d42d1e7f6362bec38171ea3edf433541db4e4ad969f",
      "index": 0,
      "sequence": 4294967295,
      "amount": 625000000
    }
  ]
}
```
#### Resopnse:
```json
{
  "gate": "bitcoin",
  "tx": {
    "amount": 1020304,
    "byte_fee": 1,
    "to_address": "1Bp9U1ogV3A14FMvKbRJms7ctyso4Z4Tcx",
    "utxo": [
      {
        "hash": "fff7f7881a8099afa6940d42d1e7f6362bec38171ea3edf433541db4e4ad969f",
        "sequence": 4294967295,
        "amount": 625000000
      }
    ],
    "change_address": "1FQc5LdgGHMHEN9nwkjmz6tWkxhPpxBvBU"
  },
  "signed": {
    "transaction": {
      "version": 1,
      "inputs": [
        {
          "previousOutput": {
            "hash": "//f3iBqAma+mlA1C0ef2NivsOBceo+30M1QdtOStlp8="
          },
          "sequence": 4294967295
        }
      ],
      "outputs": [
        {
          "value": 1020304,
          "script": "dqkUdpvf+WoC+RNaHRm3SdtqeP4H3JCIrA=="
        },
        {
          "value": 623979549,
          "script": "dqkUngibaIngMtRuO5FaM5Lt/WFvscSIrA=="
        }
      ]
    },
    "encoded": "AQAAAAABAf/394gagJmvppQNQtHn9jYr7DgXHqPt9DNUHbTkrZafAAAAAAD/////ApCRDwAAAAAAGXapFHab3/lqAvkTWh0Zt0nbanj+B9yQiKwdLDElAAAAABl2qRSeCJtoieAy1G47kVozku39YW+xxIisAkgwRQIhAKpaA2tEOJ14YWyvYgU/dVOI6kd55XHU3tiW8ZKapuxBAiA1FYj1e3NUB27HlLBSJ6JD6arLoeLrvUea3IaZXfBKdAEhA9+GNcM+rgKOINkSEg7hpzBMO7HkVOaWgrFvzx46MSj9AAAAAA==",
    "transaction_id": "aabc41f7c15071c1cd95e5fb10280fad5f52bff49353c932149b93f78e898e6f"
  }
}
```
## NOTE: Hash strings now base64-encoded. This behaviour is going to be changed asap.
