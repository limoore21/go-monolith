package repository

type User struct {
	Firstname string
	Surname   string
}

func (u User) GetFirstName() string {
	return u.Firstname
}

func (u User) GetSurName() string {
	return u.Surname
}
