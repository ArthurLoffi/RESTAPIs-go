package user

type User struct {
	ID string `json:"ID"`
	Name string `json:"Name"`
}

var Users = []User{}