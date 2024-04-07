package service

import (
	"bufio"
	"fmt"
	"github.com/briandowns/spinner"
	"github.com/logrusorgru/aurora/v4"
	"lab-work-2/database"
	"log"
	"os"
	"os/exec"
	"strings"
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
	ui.ClearConsole()
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
		ui.CommandPanel()
	}
}

func (ui UI) Input() (out string) {
	fmt.Print(aurora.BrightYellow("Выбор раздела: "))
	if _, err := fmt.Scanln(&out); err != nil {
		return ""
	}
	return
}

func (ui UI) InputToEnter() string {
	buf, _ := bufio.NewReader(os.Stdin).ReadBytes('\n')
	return strings.TrimSpace(string(buf))
}

func (ui UI) AnyKeyToBack(anyFunc func()) {
	fmt.Println("Нажмите любую кнопку для возвращения назад ...")
	bufio.NewReader(os.Stdin).ReadByte()
	anyFunc()
}

func (ui UI) ClearConsole() {
	cmd := exec.Command("clear")
	if err := cmd.Run(); err != nil {
		log.Fatalf("Failed to clear console: %v", err)
	}
}
