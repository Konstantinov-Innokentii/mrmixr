package entities

type User struct {
	Id       int
	Email    string
	Username string
	password string
	Teams    *[]Team
}
