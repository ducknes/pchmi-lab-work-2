package main

import (
	"context"
	"fmt"
	"github.com/logrusorgru/aurora/v4"
	"lab-work-2/config"
	"lab-work-2/database"
	"lab-work-2/service"
	"os"
)

func main() {
	mainCtx := context.Background()

	cfg, err := config.New()
	if err != nil {
		fmt.Println(aurora.Red(err))
		os.Exit(1)
	}

	repos, err := database.NewLogisticRepository(mainCtx, cfg)
	if err != nil {
		fmt.Println(aurora.Red(err))
		os.Exit(1)
	}

	ui := service.NewUI(repos)
	ui.Start()
}
