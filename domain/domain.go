package domain

type User struct {
	Uid      int64
	Username string
	Lastname string
}

type UserList struct {
	Users []User
}
