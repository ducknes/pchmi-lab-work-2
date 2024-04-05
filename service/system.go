package service

import "fmt"

func (ui UI) SystemMonitor() {
	fmt.Println("[ 1 ] : Состояние сети")
	fmt.Println("[ 2 ] : Состояние сервера")
	fmt.Println("[ 0 ] : Назад")

	switch ui.Input() {
	case "1":
	case "2":
	case "0":
		ui.CommandPanel()
	default:
		fmt.Println("такого раздела не существует или произошла ошибка при вводе")
		ui.SystemMonitor()
	}
}
