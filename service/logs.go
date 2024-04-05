package service

import "fmt"

func (ui UI) Logs() {
	fmt.Println("[ 1 ] : Логи склада")
	fmt.Println("[ 2 ] : Логи ПВЗ")
	fmt.Println("[ 0 ] : Назад")

	switch ui.Input() {
	case "1":
	case "2":
	case "0":
		ui.CommandPanel()
	default:
		fmt.Println("такого раздела не существует или произошла ошибка при вводе")
		ui.Logs()
	}
}
