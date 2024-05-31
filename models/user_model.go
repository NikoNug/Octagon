package models

type User struct {
	UserID    int    `json:"UserID"`
	Firstname string `json:"Firstname"`
	Lastname  string `json:"Lastname"`
	Username  string `json:"Username"`
	Email     string `json:"Email"`
	Password  string `json:"Password"`
	Address   string `json:"Address"`
	City      string `json:"City"`
}

type Users struct {
	Persons []User `json:"users"`
}
