# Go Network Programming Exercises

A collection of low-level network programming examples in Go, demonstrating how protocols work under the hood without relying on high-level web frameworks.

## 📦 Implementations
- **`tcpConnection/`**: Concurrent TCP server handling multiple clients using Goroutines.
- **`udpConnection/`**: Connectionless UDP server/client for low-latency messaging.
- **`http/`**: Building a minimal HTTP server from raw TCP sockets to understand HTTP/1.1 framing.

## 🧠 Key Concepts Covered
- Socket creation and binding (`net.Listen`).
- Concurrent connection handling (`go handleConnection(conn)`).
- Handling timeouts and graceful shutdowns.
- Manual data serialization over the wire.

## 🚀 How to Run
Navigate into any subdirectory and run the server:
```bash
cd tcpConnection
go run server.go

In a separate terminal, run corresponding client.
