package main

import "fmt"

type Users struct {
	Name     string
	Surname  string
	Age      int
	Email    string
	Password string
}

func (u *Users) changeName(name *string) {
	u.Name = *name

}
func (u *Users) changeAge(age *int) {
	u.Age = *age

}
func (u *Users) changeEmail(email *string) {
	u.Email = *email
}

func (u *Users) changePassword(password *string) {
	u.Password = *password

}

func main() {

	u1 := &Users{
		Name:     "Johnatan",
		Surname:  "Figueroa",
		Age:      22,
		Email:    "joni2002@yahoo.com",
		Password: "velezcapo",
	}
	var (
		name     string = "Jonatan"
		age      int    = 21
		email    string = "jonatanfigueroa@gmail.com"
		password string = "JNk3jds$)"
	)
	u1.changeName(&name)
	u1.changeAge(&age)
	u1.changeEmail(&email)
	u1.changePassword(&password)

	fmt.Println("Actualizado!", u1.Name, u1.Age, u1.Email, u1.Password)
}
