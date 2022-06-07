package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type Message struct {
	Greeting string `json:"greeting"`
}

var(
	wsUpgrader = websocket.Upgrader{
		ReadBufferSize:1024,
		WriteBufferSize:1024,
	}
	wsConn *websocket.Conn
)

func WsEndpoint(w http.ResponseWriter, r *http.Request){
	wsUpgrader.CheckOrigin=func(r *http.Request) bool{
		//check http.Request
		return true

	}
	wsConn, err:= wsUpgrader.Upgrade(w,r,nil)
	if err != nil {
		fmt.Printf("Error upgrading websocket: %s\n", err.Error())
	}

	defer wsConn.Close()
	for{
		var msg Message 
		err:= wsConn.ReadJSON(&msg)
		if err != nil{
			fmt.Printf("Error reading json: %s\n", err.Error())
		}
		fmt.Printf("Message Recieved: %s\n", msg.Greeting)
	}
} 

func main(){
	router := mux.NewRouter()
	log.Fatal(http.ListenAndServe(":9100",router))
}