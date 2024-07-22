package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var dbUser *os.File
var dbMessage *os.File

func main() {
	port := os.Getenv("PORT")
	var err error
	dbUser, err = os.OpenFile("./data/users.txt",
		os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	defer dbUser.Close()

	if err != nil {
		log.Panic("Error open file users.txt: " + err.Error())
		return
	}

	dbMessage, err = os.OpenFile("./data/messages.txt",
		os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	defer dbMessage.Close()

	if err != nil {
		log.Panic("Error open file messages.txt: " + err.Error())
		return
	}

	readAllUsers()

	http.HandleFunc("/", checkService)
	http.HandleFunc("/api/sign-in", signIn)
	http.HandleFunc("/ws/", handleConnections)

	fmt.Println("Это я!")

	go handleMessages()

	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		log.Panic("Error starting server: " + err.Error())
	}
}

func checkService(w http.ResponseWriter, r *http.Request) {
	responseString(w, `{"success": true}`)
}

func responseString(w http.ResponseWriter, text string) {
	responseJson(w, []byte(text))
}

func responseJson(w http.ResponseWriter, v []byte) {
	w.Header().Set("Content-Type", "application/json;  charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(v); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`Error`))
	}
}
