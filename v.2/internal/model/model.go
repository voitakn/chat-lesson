package model

type MessageNew struct {
	Body     string `db:"body" json:"body"`
	PersonId string `db:"person_id" json:"person_id"`
}

type Message struct {
	Id       string `db:"id" json:"id"`
	Created  string `db:"created_at" json:"created_at"`
	Body     string `db:"body" json:"body"`
	PersonId string `db:"person_id" json:"person_id"`
}

type PersonNew struct {
	UserName string `db:"username" json:"username"`
}

type Person struct {
	Id       string `db:"id" json:"id"`
	UserName string `db:"username" json:"username"`
}
