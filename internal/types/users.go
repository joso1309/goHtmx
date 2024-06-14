package types

type User struct {
	ID        uint
	FirstName string `form:"firstname"`
	LastName  string `form:"lastname"`
	Email     string `form:"email"`
	Date      string `form:"date"`
	Time      string `form:"time"`
	Select    string `form:"select"`
}
