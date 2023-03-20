package user_adapter

// CreateUser is used to create user
func CreateUser(name string, age int) *User {
	return &User{
		Name: name,
		Age:  age,
	}
}

func (u *User) GetUser() *User {
	return u
}
