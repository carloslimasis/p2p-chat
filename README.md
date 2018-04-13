# Simple P2P chat

This is a simple chat to switch messages.

## How It Works

First, starts the server

```sh
$ go run server/server.go
```

Then, starts the client with your username

```sh
$ go run client/client.go Carlos
```

```sh
$ go run client/client.go Maria
```

Then, Carlos send a message to Maria
```sh
$ go run client/client.go Carlos
$ Trying to connect with 127.0.0.1:8888...
$ Connected with 127.0.0.1:8888!
$ Hello Maria!
```

```sh
$ go run client/client.go Maria
$ Trying to connect with 127.0.0.1:8888...
$ Connected with 127.0.0.1:8888!
$ Carlos -> Hello Maria!
```