package main

import (
	"fmt"
	"os"

	"gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.Name = "sapper"
	app.Usage = "A client for WordPress."

	wordpress_endpoint := os.Getenv("WORDPRESS_ENDPOINT")
	wordpress_cookie := os.Getenv("WORDPRESS_COOKIE")

	app.Action = func(c *cli.Context) error {
		fmt.Println(wordpress_endpoint)
		fmt.Println(wordpress_cookie)
		return nil
	}

	app.Run(os.Args)
}
