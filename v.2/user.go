package main

import (
	"bufio"
	"chat-lesson/internal/model"
	"chat-lesson/internal/repository/person"
	"chat-lesson/pkg/pgdb"
	"context"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func signIn(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		responseString(w, `{"success": false, "msg": "POST method is required"}`)
		return
	}
	var user User
	body, err := io.ReadAll(r.Body)
	if err != nil {
		responseString(w, fmt.Sprintf(`{"success": false,"msg": "%s"}`, err.Error()))
		return
	}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responseString(w, fmt.Sprintf(`{"success": false, "msg": "%s"}`, err.Error()))
		return
	}
	if len(user.Name) == 0 {
		responseString(w, `{"success": false, "msg": "Please enter your name"}`)
		return
	}

	userKey := fmt.Sprintf(`%s_%v`, user.Name, time.Now())
	hasher := sha1.New()
	hasher.Write([]byte(userKey))
	userToken := hex.EncodeToString(hasher.Sum(nil))

	user.Token = userToken
	user.Id = len(userData.Items) + 1

	userData.Items = append(userData.Items, user)
	last := &userData.Items[len(userData.Items)-1]
	userData.IDx[user.Id] = last
	userData.TKx[user.Token] = last

	userJson, _ := json.Marshal(user)

	// Write new user to data file
	_, err = fmt.Fprintln(dbUser, string(userJson))
	if err != nil {
		responseString(w, fmt.Sprintf(`{"success": false, "msg": "%v"}`, err.Error()))
		return
	}

	repoPerson := person.New(pgdb.DB.Conn())
	ctx := context.Background()
	result, err := repoPerson.PersonCreate(ctx, model.PersonNew{
		UserName: user.Name,
	})

	log.Println("signIn", result)

	responseJson(w, userJson)
}

func readAllUsers() {
	scanner := bufio.NewScanner(dbUser)
	for scanner.Scan() {
		var u User
		if err := json.Unmarshal([]byte(scanner.Text()), &u); err == nil {
			userData.Items = append(userData.Items, u)
			row := &userData.Items[len(userData.Items)-1]
			userData.IDx[row.Id] = row
			userData.TKx[row.Token] = row
		}
	}
	log.Println("readAllUsers", userData)
}
