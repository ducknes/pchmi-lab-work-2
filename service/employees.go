package service

import "fmt"

func (ui UI) EmployeeCommandPanel() {
	fmt.Println("[ 1 ] : Список сотрудников")
	fmt.Println("[ 2 ] : Добавить нового сотрудника")
	fmt.Println("[ 3 ] : Изменить данные сотрудника")
	fmt.Println("[ 4 ] : Удалить сотрудника")
	fmt.Println("[ 0 ] : Назад")

	switch ui.Input() {
	case "1":
	case "2":
	case "3":
	case "0":
		ui.CommandPanel()
	default:
		fmt.Println("такого раздела не существует или произошла ошибка при вводе")
	}
}
