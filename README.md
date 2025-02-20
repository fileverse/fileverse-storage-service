# Fileverse Storage Service

An upload service for Fileverse that enables file uploads to IPFS using a pinning service (currently Pinata).

## Prerequisites
- Go (version 1.21.11 or later)
- Git
- Docker (optional, for containerized deployment)

## Installation

1. **Clone the repository:**
   ```sh
   git clone https://github.com/fileverse/fileverse-storage-service.git
   cd fileverse-storage-service
   ```

2. **Install dependencies:**
   ```sh
   go mod tidy
   go mod vendor
   ```

## Configuration

Create a `.env` file in the root directory:
```sh
MONGO_URI=
PINATA_JWT=
ETH_GOERLI_RPC_URL=
ETH_MAINNET_RPC_URL=
ETH_SEPOLIA_RPC_URL=
GNOSIS_MAINNET_RPC_URL=
GNOSIS_TESTNET_RPC_URL=
POLYGON_MAINNET_RPC_URL=
SERVICE_DID=
HOST=local
# Add other configuration variables as needed
```

## Running the Service

### Local Development
```sh
go run main.go
```

### Using Docker
```sh
docker compose up --build -d
```

## API Endpoints

base url: `<host>/new-storage-service/api/v1`

- **POST** `upload/identity` - Upload an identity file
  ```sh
  curl --location '<host>/new-storage-service/api/v1/upload/identity' \
   --header 'x-api-key: <api-key>' \
   --header 'Authorization: <ucan-token>' \
   --form 'file=@"<file-path>"'
   ```


## Project Structure
```
src
├── config
│   ├── config.go
│   ├── dev.yaml // dev config
│   ├── local.yaml // local config
│   └── prod.yaml // prod config
├── constants // constants for the project
│   └── db.go // db constants
├── pkg
│   ├── db
│   │   ├── client.go // db client connection
│   │   ├── query.go
│   │   └── utils.go
│   ├── goucan
│   │   ├── capabilities.go
│   │   ├── constants.go
│   │   ├── ucan.go // ucan verification module in golang
│   │   ├── ucan_test.go
│   │   └── utils.go
│   ├── logger
│   │   └── logger.go // service logger using slog with context and custom fields
│   └── reporter
│       ├── slack.go // slack reporter for error and warning messages, not used currently
│       └── slack_test.go
└── services
    ├── contracts
    │   ├── FileversePortal // FileversePortal contract
    │   ├── IdentityModule // IdentityModule contract
    │   └── utils.go 
    ├── dbservice
    │   ├── files.go 
    │   └── limit.go 
    ├── ipfs
    │   └── pinata // ipfs pinning service
    ├── server
    │   ├── handler // handler for the api
    │   ├── middleware // middleware for the api
    │   ├── server.go // server setup
    │   └── v1.go // v1 api routes
    └── upload
        ├── structs.go // structs for the upload api
        └── upload.go // upload api implementation
```
