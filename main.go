package main

import (
	"fmt"
	"os"

	"gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.Name = "Skeleton"
	app.Usage = "Scaffolding for a command line interface."
	app.Action = func(c *cli.Context) error {
		fmt.Println("Hello, world!")
		return nil
	}

	app.Run(os.Args)
}
