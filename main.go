package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

var upgrader = websocket.Upgrader{}

func liwaSocket(w http.ResponseWriter, req *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	c, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatalln("Upgrade is failed ", err)
		return
	}
	defer c.Close()
	for {
		mt, message, error := c.ReadMessage()
		if error != nil {
			log.Fatalln("Error parsing the message", error)
		}
		log.Printf("recv: %s", message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("Write failed", error)
		}
	}

}
func main() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/socket", liwaSocket)
	http.ListenAndServe(*addr, nil)
}
