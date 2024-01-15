package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close()

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			return
		}

		message := string(p)
		fmt.Printf("Recebido: %s\n", message)

		err = conn.WriteMessage(messageType, p)
		if err != nil {
			return
		}
	}

}

func main() {
	http.HandleFunc("/chat", handleConnection)

	fmt.Println("Chat initialize in http://localhost:8080/chat")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
