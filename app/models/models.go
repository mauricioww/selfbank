package models

// anotations?
type User struct {
	Id        uint
	FirstName string
	LastName  string
	Password  string
	Email     string
	Admin     bool
	Active    bool
}
