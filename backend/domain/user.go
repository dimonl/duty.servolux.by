package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserData struct {
	FirstName string `json:"firstname"`
	LastName  int    `json:"lastname"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	IsAdmin   int    `json:"isadmin"`
}

type IdUserData struct {
	Id string `json:"id"`
}

type PatchUserData struct {
	Id        string `json:"id"`
	FirstName string `json:"firstname"`
}

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	FirstName string
	LastName  int
	Phone     string
	Email     string
	Password  string
	IsAdmin   int
}

func NewUser() *User {
	return &User{}
}

func ConvertUser(data UserData) *User {
	newUser := NewUser()
	newUser.FirstName = data.FirstName
	newUser.LastName = data.LastName
	newUser.Phone = data.Phone
	newUser.Email = data.Email
	newUser.Password = data.Password
	newUser.IsAdmin = data.IsAdmin
	return newUser
}
