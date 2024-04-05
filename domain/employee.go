package domain

import (
	"math/rand"
	"time"
)

var defaultEmployee = map[int]Employee{
	1: {
		Surname:        "Медведев",
		Name:           "Максим",
		Patronymic:     "Викторович",
		DocumentNumber: "4523 234233",
		Birthday:       time.Date(2003, 11, 02, 0, 0, 0, 0, time.UTC),
	},
	2: {
		Surname:        "Новоженов",
		Name:           "Иван",
		Patronymic:     "Алексеевич",
		DocumentNumber: "3252 343452",
		Birthday:       time.Date(2004, 1, 9, 0, 0, 0, 0, time.UTC),
	},
	3: {
		Surname:        "Антонов",
		Name:           "Илья",
		Patronymic:     "Игоревич",
		DocumentNumber: "9284 332633",
		Birthday:       time.Date(2003, 6, 03, 0, 0, 0, 0, time.UTC),
	},
}

type Employee struct {
	Id             int       `db:"id"`
	Surname        string    `db:"surname"`
	Name           string    `db:"name"`
	Patronymic     string    `db:"patronymic"`
	DocumentNumber string    `db:"document_number"`
	Birthday       time.Time `db:"birthday"`
}

func GetRandomDefaultEmployee() Employee {
	return defaultEmployee[rand.Intn(len(defaultEmployee))+1]
}
