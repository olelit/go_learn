package modes

import "golang.org/x/crypto/bcrypt"

type Users struct {
	Id       int
	Email    string `json:"email"`
	Password string `json:"password"`
}

var List = make([]Users, 0)

func Register(user Users) []Users {
	hashedPass, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPass)
	List = append(List, user)
	return List
}
