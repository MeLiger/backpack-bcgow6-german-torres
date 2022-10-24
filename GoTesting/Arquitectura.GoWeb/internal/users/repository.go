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
	Patch(id int, name string, email string, age *int, height *int, active *bool, date string) (User, error)
	Delete(id int) (User, error)
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
	if err := r.db.Read(&us); err != nil {
		return User{}, err
	}
	us = append(us, user)
	if err := r.db.Write(us); err != nil {
		return User{}, err
	}
	lastID = user.Id
	return user, nil
}

func (r *repository) GetAll() ([]User, error) {
	var us []User
	if err := r.db.Read(&us); err != nil {
		return us, err
	}

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
	var u []User
	r.db.Read(&users)
	updated := User{}
	for i, user := range users {
		if user.Id == id {
			updated = User{id, name, email, age, height, active, date}
			u[i] = updated
			err := r.db.Write(u)
			if err != nil {
				return User{}, err
			}
			return updated, nil
		}
	}
	return updated, fmt.Errorf("usuario %d no encontrado", id)
}

func (r *repository) Delete(id int) (User, error) {
	var u []User
	r.db.Read(&users)

	deleted := User{}
	for i, user := range users {
		if user.Id == id {
			deleted = users[i]
			u = append(u[:i], u[i+1:]...)
			err := r.db.Write(u)
			if err != nil {
				return User{}, err
			}
			return deleted, nil
		}
	}
	return deleted, fmt.Errorf("usuario %d no existente", id)
}

func (r *repository) Patch(id int, name string, email string, age *int, height *int, active *bool, date string) (User, error) {
	var u []User
	r.db.Read(&users)

	for i, user := range users {
		if user.Id == id {
			if name != "" {
				u[i].Name = name
			}
			if email != "" {
				u[i].Email = email
			}
			if age != nil {
				u[i].Age = *age
			}
			if height != nil {
				u[i].Height = *height
			}
			if active != nil {
				u[i].Active = *active
			}
			if date != "" {
				u[i].Date = date
			}
			err := r.db.Write(users)
			if err != nil {
				return User{}, err
			}
			return u[i], nil
		}
	}

	return User{}, fmt.Errorf("usuario %d no registrado", id)

}
