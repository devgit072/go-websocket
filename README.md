## Websocket protocol
Websocket is persistent bidirectional protocol used in chat applications. Unlike other protocol like http, 
you don't have to establish connection again and again. Its a persistent duplex communication. 
Client can send message to server and server can send the message to client in the same established TCP connection. 
The use for websocket chat application, trading application,  

## Depedencies required.
We can use gorilla library which implement websocket.
```bash
go get github.com/gorilla/websocket v1.5.0
```

## Demo
This code is implementation for websocket server. Once you start the server, you need websocket client. Either 
you can use command-line curl for websocket connection, or write some script in JS or Go which can open websocket connection with 
server. The best to use some tool. For Mac, we can use ```websocat``` tool.
```bash
brew install websocat
```
Next call the websocket end-point:
```bash
websocat ws://localhost:8081/wsendpoint
```
Output:
```bash
from client: My request
From server: You request received
```