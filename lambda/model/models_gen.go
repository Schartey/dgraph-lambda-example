package model

type User struct {
	Id       string `json:"id" dql:"uid"`
	Username string `json:"username" dql:"User.username"`
	Email    string `json:"email" dql:"User.email"`
}
