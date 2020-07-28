package todo

import "time"

type Todo struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Text      string    `json:"text"`
	Completed bool      `json:"completed"`
	CreatedOn time.Time `json:"created_on"`
}
