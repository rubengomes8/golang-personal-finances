# golang-personal-finances

[![Actions Status](https://github.com/rubengomes8/golang-personal-finances/workflows/build/badge.svg)](https://github.com/rubengomes8/golang-personal-finances/actions)
[![codecov](https://codecov.io/gh/rubengomes8/golang-personal-finances/branch/main/graph/badge.svg)](https://codecov.io/gh/rubengomes8/golang-personal-finances)

## Objective

This repo builds a service whose goal is to manage personal finances. It creates:
- a database
- a HTTP Server to manage expenses/incomes and its metadata (categories, cards)
- a gRPC Server to manage expenses/incomes and its metadata (categories, cards)

Here is a simple diagram that describes the project:
![Project Diagram](resources/images/golang-personal-finances-backend-diagram.png)

## How to run locally

To run locally you need to:
1. set up the docker environment with a postgres image database
2. build the gRPC and HTTP Server binaries, and run one/both of them
3. you can test the HTTP Server using postman (**TODO**: add postman collection here)
4. you can test the gRPC Server using this client https://github.com/rubengomes8/golang-personal-finances-client -(**TODO**: need to update it)

### 1. Set up docker environment
Open the Docker Desktop and run on the project root folder:

```make up```

### 2. Build server binaries
On the project root folder run:
```make build```

It will create the binaries on the `bin/` folder. You can execute them by running:

```
./bin/grpc_server
./bin/http_server
``` 

### 3. Test the HTTP Server
You can test the HTTP Server using this postman collection (**TODO**)

### 4. Test the gRPC Server
You can test the gRPC Server using this gRPC client (**TODO**) or creating your own.





