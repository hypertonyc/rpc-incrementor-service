# RPC-INCREMENTOR-SERVICE

> Test task 

## Description

The task involved developing an RPC service in the Go programming language. This service needed to fulfill two main functions: returning a numerical value and incrementing this value through a designated method call. Dynamic parameters, including the increment step and an upper threshold, were configurable during the service's runtime. It was imperative that the service retained both the numeric value and the configuration settings within a repository, ensuring that these values could be seamlessly restored following any service restarts.

Available gRPC methods:

* GetNumber - returns current value
* IncrementNumber - increments current value
* SetSettings - set current upper limit and increment step

### Used packages

* google.golang.org/grpc - gRPC & Protobuf package
* github.com/lib/pq - PostgreSQL driver
* github.com/jessevdk/go-flags - Package to parse commandline args (used in client app)
* go.uber.org/zap - Logger package
* github.com/golang-migrate/migrate/v4 - Creation of DB scheme and addition of initial data 
* github.com/stretchr/testify - Unit testing utils

## Dependencies

* docker
* make

## Commands

- Generate go files based on proto files:

    ```bash
    make proto
    ```
- Run unit tests:

    ```bash
    make test
    ```

- Run server in debug mode (with in-memory storage):

    ```bash
    make debug
    ```

- Run End-to-End test:

    ```bash
    make e2e
    ```

- Run server (build docker images and run them in background):

    ```bash
    make up
    ```

- Stop server:

    ```bash
    make stop
    ```

## TODO

* create domain entities at startup only and don't execute db queries for every rpc request
* increase tests coverage
* add mocks for tests