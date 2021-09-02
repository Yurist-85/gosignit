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
          "hash": "fff7f7881a8099afa6940d42d1e7f6362bec38171ea3edf433541db4e4ad969f",
          "sequence": 4294967295
        }
      ],
      "outputs": [
        {
          "value": 1020304,
          "script": "76a914769bdff96a02f9135a1d19b749db6a78fe07dc9088ac"
        },
        {
          "value": 623979549,
          "script": "76a9149e089b6889e032d46e3b915a3392edfd616fb1c488ac"
        }
      ]
    },
    "encoded": "01000000000101fff7f7881a8099afa6940d42d1e7f6362bec38171ea3edf433541db4e4ad969f0000000000ffffffff0290910f00000000001976a914769bdff96a02f9135a1d19b749db6a78fe07dc9088ac1d2c3125000000001976a9149e089b6889e032d46e3b915a3392edfd616fb1c488ac02483045022100aa5a036b44389d78616caf62053f755388ea4779e571d4ded896f1929aa6ec410220351588f57b7354076ec794b05227a243e9aacba1e2ebbd479adc86995df04a74012103df8635c33eae028e20d912120ee1a7304c3bb1e454e69682b16fcf1e3a3128fd00000000",
    "transaction_id": "aabc41f7c15071c1cd95e5fb10280fad5f52bff49353c932149b93f78e898e6f"
  }
}
```
## NOTE: Bytes values (hashes, scripts, keys...) are HEX-encoded to strings.