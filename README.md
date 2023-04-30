# go-grpc-sqlc-example

An example project that I created to show how to build a Go application that uses modern
tech that easily develops and maintains in the long term.

It is the most simple ToDo List application;
- sign-in to account
- Create, read, update, and delete ToDos.

This application assumes running on k8s and connects to RDB (MySQL).

Mainly, It focused on showing how to work on the following things.

- Protobuf / gRPC toolchains (gRPC Gateway, protoc-gen-validate, Buf)
- sqlc

The basic philosophy of this project is "Schema-driven" development.
Go is declarative means you should write codes explicitly.
You may have had an experience you realized a lot of duplicated codes. However, still, those are required.
"Code generation" is an effective way to build your application in Go.

This project has two sorts of schema. 

- Protobuf
    - Client request, server response, validations, and documents (openapiv2).
- sqlc
    - Database schema, queries, column validations.

## Contribution

Feel free to create any PR.
I always welcome your contribution.
