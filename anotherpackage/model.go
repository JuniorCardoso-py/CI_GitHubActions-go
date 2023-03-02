package user_adapter

// User struct
type User struct {
	Name string
	Age  int
}

func CreateUser(name string, age int) *User {
	return &User{
		Name: name,
		Age:  age,
	}
}

func (u *User) GetUser() *User {
	return u
}
