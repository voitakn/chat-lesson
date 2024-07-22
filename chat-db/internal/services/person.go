package services

import (
	"chat-db/internal/repository"
	"chat-db/internal/utils"
	"chat-db/pkg/pgdb"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func PersonCreate(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	if r.Method != "POST" {
		utils.ResponseJson(w, []byte(`{"success": false, "msg": "POST method is required"}`))
		return
	}

	var person repository.Person
	body, err := io.ReadAll(r.Body)
	if err != nil {
		utils.ResponseJson(w, []byte(
			fmt.Sprintf(`{"success": false,"msg": "%s"}`, err.Error()),
		))
		return
	}

	err = json.Unmarshal(body, &person)

	if err != nil {
		utils.ResponseJson(w, []byte(
			fmt.Sprintf(`{"success": false,"msg": "%s"}`, err.Error()),
		))
		return
	}

	if person.Username == nil {
		utils.ResponseJson(w, []byte(
			fmt.Sprintf(`{"success": false,"msg": "You have to input username"}`),
		))
		return
	}

	repo := repository.New(pgdb.DB.Conn())

	result, err := repo.CreateUser(ctx, person.Username)

	if err != nil {
		utils.ResponseJson(w, []byte(
			fmt.Sprintf(`{"success": false,"msg": "%s"}`, err.Error()),
		))
		return
	}

	resulByte, err := json.Marshal(result)
	utils.ResponseJson(w, resulByte)
}
