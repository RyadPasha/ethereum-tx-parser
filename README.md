
# Ethereum Transaction Parser

A Go-based Ethereum blockchain transaction parser that tracks and queries transactions for subscribed addresses. The parser interacts with the Ethereum blockchain using the Ethereum JSON-RPC interface and supports real-time transaction alerts and queries.

## Features

- **Ethereum Blockchain Integration**: Uses the Ethereum JSON-RPC interface to retrieve transaction data.
- **Address Subscription**: Subscribe to Ethereum addresses to track inbound and outbound transactions.
- **Concurrent Access**: Implements thread-safe access to transaction and subscription data.
- **Modular Architecture**: Easily extendable to support future `storage` solutions or `notification` integrations.
- **Storage Integration**: Storage functionality is created and ready to be extended in the future (currently uses in-memory storage).
- **Notification System**: Notification functionality  created and is ready to be extended in the future. Currently, it logs the notifications but does not send any actual alerts. It is designed for future integration with notification services to send real-time alerts about transactions.

## Architecture

![Ethereum Transaction Parser Architecture](https://i.ibb.co/7kTwHJH/Ethereum-Tx-Parser.png)

## Configuration

The configuration for the Ethereum Transaction Parser is managed through constants defined in the `config` package. Below are the key configuration settings used in the application:

- **DefaultAppPort**: The default port for the HTTP server (`8080`).
- **EthereumRPCURL**: The Ethereum JSON-RPC endpoint used to interact with the blockchain. It is set to `https://ethereum-rpc.publicnode.com` by default.
- **JSONRPCVersion**: The version of the JSON-RPC protocol (`2.0`).
- **BlockMethod**: The JSON-RPC method used to retrieve the current Ethereum block number (`eth_blockNumber`).
- **BlockByNumberMethod**: The JSON-RPC method used to retrieve a block by its number (`eth_getBlockByNumber`).

## Prerequisites

- **Go** 1.18 or later
- **Ethereum JSON-RPC Endpoint** (Public node or custom)

## Installation

### 1. Clone the Repository

```bash
git clone https://github.com/ryadpasha/ethereum-tx-parser.git
cd ethereum-tx-parser
```

### 2. Install Dependencies

This project does not require any external libraries, so Go's standard library is sufficient.

### 3. Build the Project

To build the Go application:

```bash
go build -o ethereum-tx-parser .
```

### 4. Run the Application

To start the Ethereum Transaction Parser:

```bash
./ethereum-tx-parser
```

By default, the server will start on port `8080`. You can customize the port by modifying the `main.go` file or passing an environment variable `PORT`.

### 5. Running the Server with Custom Port

To run the server on a custom port:

```bash
PORT=8080 ./ethereum-tx-parser
```

You should see the message:

```
Server is started on port 8080
```

### 6. Running with `make`

If you have `make` installed, you can use the provided `Makefile` to build and run the application more easily.

- **Build the Project**:

  ```bash
  make build
  ```

- **Run the Application**:

  To build and run the application:

  ```bash
  make run
  ```

- **Run the Application with Custom Port**:

  To run the application with a custom port (e.g., port 8081):

  ```bash
  make run-custom PORT=8081
  ```

### 7. Accessing the API

Once the server is running, the following endpoints are available:

#### 1. `GET /block`

Get the current Ethereum block being parsed.

Example request:

```bash
curl http://localhost:8080/block
```

#### 2. `POST /subscribe/?address={address}`

Subscribe an Ethereum address for transaction tracking.

Example request:

```bash
curl "http://localhost:8080/subscribe?address=0xYourEthereumAddress"
```

#### 3. `GET /transactions/?address={address}`

Retrieve the transactions for a subscribed Ethereum address.

Example request:

```bash
curl "http://localhost:8080/transactions?address=0xYourEthereumAddress"
```

## Testing

You can run the tests using Go's testing framework.

```bash
go test -v ./server
```

or

```bash
make test
```

## Contact

Created by [Mohamed Riyad](mailto:mohamed@ryad.dev).

## License

MIT License. See the [LICENSE](LICENSE) file for more information.
