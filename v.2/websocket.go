package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strings"
	"time"
)

var clients = make(map[*websocket.Conn]string)
var broadcast = make(chan wsMessage)

var wsUp = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	log.Println("handleConnections")
	// Get token
	token := strings.TrimPrefix(r.URL.Path, "/ws/")
	// Check user token
	_, okUs := userData.TKx[token]
	if !okUs {
		log.Printf(`Didn't find the user with token: %s'`, token)
		return
	}

	conn, err := wsUp.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	clients[conn] = token

	for {
		_, msgByte, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			delete(clients, conn)
			return
		}
		msg := wsMessage{
			Text:  msgByte,
			Token: token,
		}
		broadcast <- msg
	}
}

func handleMessages() {
	for {
		msg := <-broadcast

		user := userData.TKx[msg.Token]
		newMsg := Message{
			Text:     string(msg.Text),
			UserId:   user.Id,
			UserName: user.Name,
			Created:  time.Now().Format(time.DateTime),
		}

		msgJson, err := json.Marshal(newMsg)
		if err != nil {
			log.Println(err)
		}

		// Write to datafile messages
		_, err = fmt.Fprintln(dbMessage, string(msgJson))
		if err != nil {
			log.Println(err)
		}

		for client := range clients {

			if msg.Token != clients[client] {
				err := client.WriteJSON(newMsg)
				if err != nil {
					fmt.Println(err)
					client.Close()
					delete(clients, client)
				}
			}
		}
	}
}
