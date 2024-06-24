# GOEDA

This is a simple POC to test [nats.io](https://nats.io/) using [watermill lib](https://watermill.io/)
It have two services:
- foo: publishing messages
- bar: consuming messages

## Running
It need to have a nats server instance with [Jetstream](https://docs.nats.io/nats-concepts/jetstream) enabled running locally.
Example:
```
nats-server -js
```
Run the main.go to start both services:
```
go run main.go
```
