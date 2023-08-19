package app

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// This is just for development. Don't check any origin, just return true.
		return true
	},
	// there are additional parameters which can be used for production grade code.
	// As of now, I will be commenting out this code for dev purpose.
	/*
		HandshakeTimeout:  0,
		WriteBufferPool:   nil,
		Subprotocols:      nil,
		Error:             nil,
		EnableCompression: false,
	*/
}

func setupRoutes() {
	http.HandleFunc("/", home)
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/websocket", websocketEndpoint)
}

func hello(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hello world")
	if err != nil {
		log.Println("Error: ", err)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "It is just a home page")
	if err != nil {
		log.Println("Error: ", err)
	}
}

// this method will be handling websocket connection.
func websocketEndpoint(w http.ResponseWriter, r *http.Request) {
	// create a websocket connection.
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	log.Println("Websocket end point connected successfully....")
	readAndWrite(conn)
}

// This method will actually read message from client and write it back to client.
func readAndWrite(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		log.Println("Message read: ", string(p))
		if err = conn.WriteMessage(messageType, []byte("From server: You request received")); err != nil {
			log.Println(err)
			return
		}
	}
}

func StartServer() {
	fmt.Println("Starting server....")
	setupRoutes()
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatalln(err)
	}
}
