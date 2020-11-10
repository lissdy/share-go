package basic

import "fmt"

type User struct {
	name string
	password string
}

func new() User {
	return User{
		"admin",
		"11",
	}
}

func (u *User) login() bool {
	if u.name=="admin" && u.password=="111"{
		return true
	}else {
		return false
	}
}

func main()  {
	a := new()
	b := &User{
		"Lisa",
		"222",
	}
	b.login()
	fmt.Println(a.login())
}
