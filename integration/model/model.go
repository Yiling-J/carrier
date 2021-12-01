package model

type GroupCategory struct {
	Name string
}

type Group struct {
	Name     string
	Category GroupCategory
}

type User struct {
	Name  string
	Email string
	Title string
	Group *Group
}

type Foo struct {
	Name string
}
