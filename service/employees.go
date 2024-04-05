package service

import (
	"fmt"
	"github.com/briandowns/spinner"
	"lab-work-2/domain"
	"time"

	"github.com/logrusorgru/aurora/v4"
)

func (ui UI) EmployeeCommandPanel() {
	fmt.Println("[ 1 ] : Список сотрудников")
	fmt.Println("[ 2 ] : Добавить нового сотрудника")
	fmt.Println("[ 3 ] : Изменить данные сотрудника")
	fmt.Println("[ 4 ] : Удалить сотрудника")
	fmt.Println("[ 0 ] : Назад")

	switch ui.Input() {
	case "1":
		ui.ViewAllEmployees()
	case "2":
	case "3":
	case "0":
		ui.CommandPanel()
	default:
		fmt.Println("такого раздела не существует или произошла ошибка при вводе")
		ui.EmployeeCommandPanel()
	}
}

func (ui UI) AddEmployee() {
	fmt.Println("[ 1 ] Добавить вручную")
	fmt.Println("[ 2 ] Подтянуть с госуслугами")
	fmt.Println("[ 0 ] Назад")

	switch ui.Input() {
	case "1":
		ui.handEmployeeInput()
	case "2":
		ui.getFromGosuslugi()
	case "0":
		ui.CommandPanel()
	default:
		fmt.Println("такого раздела не существует или произошла ошибка при вводе")
		ui.AddEmployee()
	}
}

func (ui UI) handEmployeeInput() {
	var employee domain.Employee
	fmt.Print("Фамилия: ")
	fmt.Scanln(&employee.Surname)
	fmt.Print("Имя: ")
	fmt.Scanln(&employee.Name)
	fmt.Print("Отчество: ")
	fmt.Scanln(&employee.Patronymic)
	fmt.Print("Номер паспорта: ")
	fmt.Scanln(&employee.DocumentNumber)
	fmt.Print("Дата рождения (в формате 2006-02-01): ")
	fmt.Scanln(&employee.Birthday)

	fmt.Println(aurora.BrightGreen("Сохраняем сотрудника"))

	err := ui.repos.AddEmployee(employee)

	spin := spinner.New(spinner.CharSets[33], 190000*time.Microsecond)
	spin.Start()
	time.Sleep(2 * time.Second)
	spin.Stop()

	if err != nil {
		fmt.Println(aurora.BrightRed(fmt.Sprintf("Не удалось сохранить сотрудника, ошибка: %v ", err)))
		fmt.Println("Нажмите любую кнопку для возвращения назад ...")
		var temp string
		fmt.Scan(&temp)
		ui.EmployeeCommandPanel()
		return
	}

	fmt.Println(aurora.BrightGreen("Сотрудник сохранен"))
	fmt.Println("[ 0 ] : Назад")
	userInput := ui.Input()
	for userInput != "0" {
		fmt.Println("вы ввели не 0")
		userInput = ui.Input()
	}
	ui.EmployeeCommandPanel()
}

func (ui UI) getFromGosuslugi() {
	var login string
	fmt.Print("Введите логин сотрудника: ")
	fmt.Scanln(&login)
	fmt.Println(aurora.BrightGreen(fmt.Sprintf("Подтягиваем сотрудника c логином %s", login)))

	err := ui.repos.AddEmployee(domain.GetRandomDefaultEmployee())

	spin := spinner.New(spinner.CharSets[33], 190000*time.Microsecond)
	spin.Start()
	time.Sleep(2 * time.Second)
	spin.Stop()

	if err != nil {
		fmt.Println(aurora.BrightRed(fmt.Sprintf("Не удалось сохранить сотрудника, ошибка: %v ", err)))
		fmt.Println("Нажмите любую кнопку для возвращения назад ...")
		var temp string
		fmt.Scan(&temp)
		ui.EmployeeCommandPanel()
		return
	}

	fmt.Println(aurora.BrightGreen("Сотрудник сохранен"))
	fmt.Println("[ 0 ] : Назад")
	userInput := ui.Input()
	for userInput != "0" {
		fmt.Println("вы ввели не 0")
		userInput = ui.Input()
	}
	ui.EmployeeCommandPanel()
}

func (ui UI) ViewAllEmployees() {
	employees, err := ui.repos.GetAllEmployees()
	if err != nil {
		fmt.Println(aurora.BrightRed(fmt.Sprintf("Не удалось получить список сотрудников. ошибка: %v", err)))
		fmt.Println("Нажмите любую кнопку для возвращения назад ...")
		var temp string
		fmt.Scan(&temp)
		ui.EmployeeCommandPanel()
		return
	}

	if len(employees) == 0 {
		fmt.Println(aurora.BrightYellow("Сотрудники не найдены"))
		fmt.Println("Нажмите любую кнопку для возвращения назад ...")
		var temp string
		fmt.Scan(&temp)
		ui.EmployeeCommandPanel()
		return
	}

	for _, employee := range employees {
		fmt.Println("----------------------------------------------------------")
		fmt.Printf("Id: %d\n", employee.Id)
		fmt.Printf("Фамилия: %s\n", employee.Surname)
		fmt.Printf("Имя: %s\n", employee.Name)
		if employee.Patronymic != "" {
			fmt.Printf("Отчество: %s\n", employee.Patronymic)
		}
		fmt.Printf("Номер паспорта: %s\n", employee.DocumentNumber)
		fmt.Printf("День рождения: %s\n", employee.Birthday.Format(time.DateOnly))
		fmt.Println("----------------------------------------------------------")
	}
	fmt.Println("[ 0 ] : Назад")
	userInput := ui.Input()
	for userInput != "0" {
		fmt.Println("вы ввели не 0")
		userInput = ui.Input()
	}
	ui.EmployeeCommandPanel()
}
