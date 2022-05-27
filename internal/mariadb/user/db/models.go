package db

type User struct {
	ID        int64
	Username  string
	Password  string
	Email     string
	LastName  string
	FirstName string
	Age       int32
}
