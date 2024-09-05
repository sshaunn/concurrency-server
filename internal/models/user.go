package models

type User struct {
	Userid    string `json:"userid"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
}

type Queue struct {
	UUID    string `json:"uuid"`
	Maxsize uint   `json:"maxsize"`
	Users   []User `json:"users"`
}
