package service

import (
	"fmt"
	"github.com/briandowns/spinner"
	"github.com/logrusorgru/aurora/v4"
	"lab-work-2/database"
	"os"
	"time"
)

type UI struct {
	repos *database.LogisticRepository
}

func NewUI(repos *database.LogisticRepository) UI {
	return UI{repos: repos}
}

func (ui UI) Start() {
	fmt.Println(aurora.BrightGreen("Запуск приложения"))
	spin := spinner.New(spinner.CharSets[33], 190000*time.Microsecond)
	spin.Start()
	time.Sleep(2 * time.Second)
	spin.Stop()
	fmt.Println(aurora.BrightGreen("Приложение запущено"))
	ui.CommandPanel()
}

func (ui UI) CommandPanel() {
	fmt.Println("[ 1 ] : Сотрудники ")
	fmt.Println("[ 2 ] : Состояние системы ")
	fmt.Println("[ 3 ] : Логи ")
	fmt.Println(aurora.Red("[ 0 ] : Выход "))

	switch ui.Input() {
	case "1":
		ui.EmployeeCommandPanel()
	case "2":
		ui.SystemMonitor()
	case "3":
		ui.Logs()
	case "0":
		os.Exit(0)
	default:
		fmt.Println("такого раздела не существует или произошла ошибка при вводе")
	}
}

func (ui UI) Input() (out string) {
	fmt.Print(aurora.BrightYellow("Выбор раздела: "))
	if _, err := fmt.Scanln(&out); err != nil {
		return ""
	}
	return
}
