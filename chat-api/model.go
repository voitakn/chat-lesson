package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"time"
)

type User struct {
	Id    int    `json:"id"`
	Token string `json:"token"`
	Name  string `json:"name"`
}

type Users struct {
	IDs   map[int]*User
	TKs   map[string]*User
	Items []*User
}

var userData = Users{
	IDs:   make(map[int]*User, 0),
	TKs:   make(map[string]*User, 0),
	Items: make([]*User, 0, 100),
}

func (u *Users) Add(user *User) {
	userKey := fmt.Sprintf(`%s_%v`, user.Name, time.Now())
	hasher := sha1.New()
	hasher.Write([]byte(userKey))
	userToken := hex.EncodeToString(hasher.Sum(nil))
	user.Token = userToken
	user.Id = len(u.Items) + 1
	u.Insert(user)
}

func (u *Users) Insert(user *User) {
	u.Items = append(u.Items, user)
	u.IDs[user.Id] = user
	u.TKs[user.Token] = user
}

type wsMessage struct {
	Text  []byte
	Token string
}

type Message struct {
	Created  string `json:"created"`
	Text     string `json:"text"`
	UserId   int    `json:"user_id"`
	UserName string `json:"name"`
}
