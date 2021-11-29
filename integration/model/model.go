package model

type Group struct {
	Name string
}

type User struct {
	Name string
	Email string
	Group *Group
}
