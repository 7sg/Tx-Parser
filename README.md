# Tx Parser

## Run the service
`go run cmd/main.go`

## To Test the service
- Get current block number
```
curl --location --request GET 'http://localhost:8080/current_block' 
```

- Subscribe a new address

```
curl --location 'http://localhost:8080/subscribe' \
--header 'Content-Type: application/json' \
--data '{
    "address": "0xbfe663dd7e363f04d3d088bc0d6f5ed087a80f6a"
}'
```

- List inbound or outbound transactions for an address
Sample Request:
```
curl --location 'http://localhost:8080/transactions?address=0x00000000a991c429ee2ec6df19d40fe0c80088b8'
```