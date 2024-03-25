package main

type Message struct {
	Created  string `json:"created"`
	Text     string `json:"text"`
	UserId   int    `json:"user_id"`
	UserName string `json:"name"`
}

type wsMessage struct {
	Text  []byte
	Token string
}

type User struct {
	Id    int    `json:"id"`
	Token string `json:"token"`
	Name  string `json:"name"`
}

type Users struct {
	IDx   map[int]*User
	TKx   map[string]*User
	Items []User
}

var userData = Users{
	IDx:   make(map[int]*User, 0),
	TKx:   make(map[string]*User, 0),
	Items: make([]User, 0, 100),
}
