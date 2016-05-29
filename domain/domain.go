package domain

type User struct {
	Uid       int64
	Firstname string
	Lastname  string
}

type UserList struct {
	Users []User
}
