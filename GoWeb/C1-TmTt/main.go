package main

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/gin-gonic/gin"
)

type Users struct {
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	Email  string  `json:"email"`
	Height float64 `json:"height"`
	Active bool    `json:"active"`
	Date   string  `json:"date"`
}

var users = []Users{

	{ID: 1, Name: "Jerry Louis", Email: "jerryls@gmail.com", Height: 1.83, Active: true, Date: "19-10-2018"},
	{ID: 2, Name: "Moria Nurina", Email: "merinawitch99@gmail.com", Height: 1.71, Active: false, Date: "03-09-2019"},
	{ID: 3, Name: "Tadeo Namunkura", Email: "tanamunku@gmail.com", Height: 1.75, Active: false, Date: "03-09-2019"},
	{ID: 4, Name: "Ken Rawson", Email: "kencherawsonpota@gmail.com", Height: 1.80, Active: true, Date: "01-11-2017"},
	{ID: 5, Name: "Rafael Batalla", Email: "rafabatalla@gmail.com", Height: 1.67, Active: true, Date: "10-05-2022"},
	{ID: 6, Name: "Eliana Franco", Email: "eliluni12@gmail.com", Height: 1.65, Active: false, Date: "02-12-2014"},
	{ID: 7, Name: "Kendra Kamala", Email: "kamalaken@gmail.com", Height: 1.80, Active: true, Date: "31-10-2017"},
	{ID: 8, Name: "Marina Lagos", Email: "marinalagos@gmail.com", Height: 1.54, Active: false, Date: "21-03-2020"},
	{ID: 9, Name: "Justina Beleza", Email: "justibel.com", Height: 1.52, Active: true, Date: "10-11-2011"},
	{ID: 10, Name: "Donatello Henriquez", Email: "donatelloturtle@gmail.com", Height: 1.88, Active: false, Date: "15-09-2016"},
}

func GetUsers() ([]Users, error) {
	var users []Users
	raw, err := os.ReadFile("./users.json")
	if err != nil {
		return nil, errors.New("Error, no se pudo leer")
	}
	json.Unmarshal(raw, &users)
	return users, nil

}

func GetAll(ctx *gin.Context) {
	ctx.JSON(200, users)
}

func welcomeGer(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Hola Germán! Ánimo, vas mejorando!",
	})

}

func main() {
	router := gin.Default()
	router.GET("/", welcomeGer)
	router.GET("/users", GetAll)

	router.Run()
}