package service

import (
	"fmt"
	"github.com/logrusorgru/aurora/v4"
	"math/rand"
)

var ethernetInfo = map[int]string{
	0: "Состояние сети: Ошибка, не удается подключиться к сети.",
	1: "Сеть работает в нормальном режиме, пакетная задержка минимальна.",
	2: "Интернет-соединение установлено, но скорость передачи данных ниже ожидаемой.",
	3: "Сеть работает, но сигнал слабый, возможны перебои в работе.",
}

var serverInfoLogs = map[int]string{
	0: "Сервер недоступен, проверьте соединение с сетью и настройки сетевых параметров.",
	1: "Сервер работает в нормальном режиме, загрузка ЦП и использование памяти на уровне нормы.",
	2: "Сервер активен, но одна или несколько служб не запущены. Проверьте настройки служб и журнал событий.",
	3: "Состояние сервера: Активен, все службы запущены.",
}

func (ui UI) SystemMonitor() {
	ui.ClearConsole()
	fmt.Println("[ 1 ] : Состояние сети")
	fmt.Println("[ 2 ] : Состояние сервера")
	fmt.Println("[ 0 ] : Назад")

	switch ui.Input() {
	case "1":
		ui.PrintSystemInfo(ethernetInfo)
	case "2":
		ui.PrintSystemInfo(serverInfoLogs)
	case "0":
		ui.CommandPanel()
	default:
		fmt.Println("такого раздела не существует или произошла ошибка при вводе")
		ui.SystemMonitor()
	}
}

func (ui UI) PrintSystemInfo(info map[int]string) {
	ui.ClearConsole()
	infoInt := rand.Intn(len(info))
	if infoInt == 0 {
		fmt.Println(aurora.BrightRed(info[infoInt]))
		ui.AnyKeyToBack(ui.SystemMonitor)
		return
	}
	fmt.Println(info[infoInt])
	ui.AnyKeyToBack(ui.SystemMonitor)
	return
}
