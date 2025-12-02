package main

type User struct {
	UID               string `json:"uid"`
	Username          string `json:"username"`
	DisplayName       string `json:"display_name"`
	PhotoURL          string `json:"photo_url"`
	encryptedPassword string
}

func (u User) VerifyPassword(s string)            {}
func (u User) GenerateJWT()                       {}
func (u *User) UpdatePassword(newPassword string) {}
func CreateNewUser(uname string, password string) User {
	return User{}
}
