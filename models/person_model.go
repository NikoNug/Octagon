package models

type Person struct {
	PersonID  int    `json:"PersonID"`
	Firstname string `json:"Firstname"`
	Lastname  string `json:"Lastname"`
	Address   string `json:"Address"`
	City      string `json:"City"`
}

type Persons struct {
	Persons []Person `json:"persons"`
}
