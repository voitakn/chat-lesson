package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

	userData.Add(&user)

	userJson, _ := json.Marshal(user)

	_, err = fmt.Fprintln(dbUser, string(userJson))

	if err != nil {
		responseString(w, fmt.Sprintf(`{"success": false, "msg": "%v"}`, err.Error()))
		return
	}

	responseJson(w, userJson)
}

func readAllUsers() {
	scanner := bufio.NewScanner(dbUser)
	for scanner.Scan() {
		var user User
		if err := json.Unmarshal([]byte(scanner.Text()), &user); err == nil {
			userData.Insert(&user)
		}
	}
}
