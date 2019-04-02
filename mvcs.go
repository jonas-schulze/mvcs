package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

type Config struct {
	VCS     string
	Remotes []Remote
}

type Remote struct {
	BaseUrl      string
	Repositories []string
}

func main() {
	app := cli.NewApp()
	app.Name = "mvcs"
	app.Usage = "clone and update several repositories at once"

	app.Commands = []cli.Command{
		{
			Name:   "clone",
			Usage:  "Clone all repositories",
			Action: cloneRepositories,
		},
		{
			Name:   "pull",
			Usage:  "Pull changes for all repositories",
			Action: pullRepositories,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
