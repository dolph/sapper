package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strings"

	"gopkg.in/headzoo/surf.v1"
	"gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.Name = "sapper"
	app.Usage = "A client for WordPress."

	// Validate environment variables.
	if os.Getenv("WORDPRESS_ENDPOINT") == "" {
		log.Fatal("WordPress endpoint required (WORDPRESS_ENDPOINT)")
		return
	}
	if os.Getenv("WORDPRESS_COOKIE") == "" {
		log.Fatal("WordPress cookie required (WORDPRESS_COOKIE)")
		return
	}

	endpointUrl, err := url.Parse(os.Getenv("WORDPRESS_ENDPOINT"))
	if err != nil {
		log.Fatal("Unable to parse WORDPRESS_ENDPOINT: ", err)
		return
	}

	// The cookies are assumed to be a copy/paste of a full dump of cookies. We
	// need to parse it apart and put it back together.
	jar, _ := cookiejar.New(nil)
	var cookies []*http.Cookie
	for _, cookieStr := range strings.Split(os.Getenv("WORDPRESS_COOKIE"), "; ") {
		cookieParts := strings.SplitN(cookieStr, "=", 2)

		cookie := &http.Cookie{
			Name:   cookieParts[0],
			Value:  cookieParts[1],
			Path:   fmt.Sprintf("%s", endpointUrl.Path),
			Domain: fmt.Sprintf(".%s", endpointUrl.Host),
		}

		cookies = append(cookies, cookie)
	}
	jar.SetCookies(endpointUrl, cookies)

	basePath := endpointUrl.Path

	bow := surf.NewBrowser()
	bow.SetCookieJar(jar)

	app.Commands = []cli.Command{
		{
			Name:  "new-user",
			Usage: "add existing user",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "email",
					Usage: "email address",
				},
				cli.StringFlag{
					Name:  "role",
					Usage: "role to assign upon creation",
				},
			},
			Action: func(c *cli.Context) error {
				if c.String("email") == "" {
					return cli.NewExitError("email address required (see: new-user --help)", 1)
				}
				if c.String("role") == "" {
					return cli.NewExitError("role required (see: new-user --help)", 1)
				}

				endpointUrl.Path = basePath + "user-new.php"
				err := bow.Open(endpointUrl.String())

				if err != nil {
					panic(err)
				}
				log.Println("Opened page:", bow.Title())

				form := bow.Form("#adduser")
				form.Input("email", c.String("email"))
				form.Input("role", c.String("role"))
				err := form.Submit()
				if err != nil {
					panic(err)
				}

				log.Println("Submitted new user:", bow.Title())

				return nil
			},
		},
	}

	app.Run(os.Args)
}
