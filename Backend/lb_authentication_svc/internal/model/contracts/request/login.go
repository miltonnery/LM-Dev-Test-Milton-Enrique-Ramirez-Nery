package request

// Create a struct that models the structure of a user, both in the request body, and in the DB
type Credentials struct {
	ID int
	Password string `contracts:"password"`
	Username string `contracts:"username"`
}
