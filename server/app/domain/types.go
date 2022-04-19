package domain

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func NewUser(id string, name string) User {
	return User{
		ID:   id,
		Name: name,
	}
}
