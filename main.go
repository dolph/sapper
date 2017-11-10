package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"gopkg.in/urfave/cli.v1"
)

type WordPressClient struct {
	endpoint string
	cookie   string
}

func (wp WordPressClient) ValidateOptions() error {
	if wp.endpoint == "" {
		return errors.New("WordPress endpoint required (WORDPRESS_ENDPOINT)")
	}

	if wp.cookie == "" {
		return errors.New("WordPress cookie required (WORDPRESS_COOKIE)")
	}

	return nil
}

func (wp WordPressClient) Get(path string) (string, error) {
	client := &http.Client{}

	endpointUrl, err := url.Parse(wp.endpoint)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	// Remove a trailing slash from the URL, if any, as we'll assume it's missing later.
	endpointUrl.Path = strings.TrimRight(endpointUrl.Path, "/")
	endpointUrl.Path = endpointUrl.Path + path

	// Build the HTTP request.
	fmt.Println("GET", endpointUrl.String())
	req, err := http.NewRequest("GET", endpointUrl.String(), nil)
	if err != nil {
		log.Fatal("Error building HTTP request for endpoint:", err)
		return "", err
	}
	req.Header.Set("Cookie", wp.cookie)

	// Send the HTTP request.
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error making HTTP request:", err)
		return "", err
	}

	// Read the HTTP response.
	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response body", err)
		return "", err
	}

	// Format the JSON response to be human-readable.
	var prettyJSON bytes.Buffer
	err = json.Indent(&prettyJSON, responseData, "", "  ")
	if err != nil {
		log.Fatal("Error parsing JSON: ", err)
		return "", err
	}

	return string(prettyJSON.Bytes()), nil
}

func main() {
	app := cli.NewApp()
	app.Name = "sapper"
	app.Usage = "A client for WordPress."

	wpClient := WordPressClient{
		endpoint: os.Getenv("WORDPRESS_ENDPOINT"),
		cookie:   os.Getenv("WORDPRESS_COOKIE"),
	}
	err := wpClient.ValidateOptions()
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(wpClient.Get("/wp/v2/users"))

	app.Action = func(c *cli.Context) error {
		return nil
	}

	// List users
	// https://developer.ibm.com/code/wp-json/wp/v2/users

	// Update a user
	// Docs: https://developer.wordpress.org/rest-api/reference/users/#update-a-user

	app.Run(os.Args)
}
