package users

import (
	"fmt"

	"github.com/MeLiger/backpack-bcgow6-german-torres/Arquitectura.GoWeb/pkg/store"
)

type User struct {
	Id     int
	Name   string `json:"name" binding:"required"`
	Email  string `binding:"required"`
	Age    int    `binding:"required"`
	Height int    `binding:"required"`
	Active bool   `binding:"required"`
	Date   string `binding:"required"`
}

var users []User

var lastID int

type Repository interface {
	GetAll() ([]User, error)
	Store(id int, name, email string, age int, height int, active bool, date string) (User, error)
	LastID() (int, error)
	Update(id int, name, email string, age int, height int, active bool, date string) (User, error)
	UpdateNameAge(id int, name string, age int) (User, error)
	Delete(id int) error
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Store(id int, name string, email string, age int, height int, active bool, date string) (User, error) {
	user := User{id, name, email, age, height, active, date}
	var us []User
	r.db.Read(&us)
	us = append(us, user)
	if err := r.db.Write(us); err != nil {
		return User{}, err
	}
	lastID = user.Id
	return user, nil
}

func (r *repository) GetAll() ([]User, error) {
	var us []User
	r.db.Read(&us)
	return us, nil
}

func (r *repository) LastID() (int, error) {
	var us []User
	if err := r.db.Read(&us); err != nil {
		return 0, err
	}
	if len(us) == 0 {
		return 0, nil
	}

	return us[len(us)-1].Id, nil
}

func (r *repository) Update(id int, name, email string, age int, height int, active bool, date string) (User, error) {
	user := User{Name: name, Email: email, Age: age, Height: height, Active: active, Date: date}
	updated := false
	for i := range users {
		if users[i].Id == id {
			user.Id = id
			users[i] = user
			updated = true
		}
	}
	if !updated {
		return User{}, fmt.Errorf("usuario %d no encontrado", id)
	}
	return user, nil
}

func (r *repository) Delete(id int) error {
	deleted := false
	var index int
	for i := range users {
		if users[i].Id == id {
			index = i
			deleted = true
		}
	}
	if !deleted {
		return fmt.Errorf("producto %d no encontrado", id)
	}
	users = append(users[:index], users[index+1:]...)
	return nil
}

func (r *repository) UpdateNameAge(id int, name string, age int) (User, error) {
	var u User
	updated := false
	for i := range users {
		if users[i].Id == id {
			users[i].Name = name
			updated = true
			u = users[i]
		}
	}
	for j := range users {
		if users[j].Id == id {
			users[j].Age = age
			updated = true
			u = users[j]
		}
	}
	if !updated {
		return User{}, fmt.Errorf("usuario %d no encontrado", id)
	}
	return u, nil

}
