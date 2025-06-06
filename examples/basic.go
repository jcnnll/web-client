package main

import (
	"fmt"
	"io"
	"log"

	"guthub.com/jcnnll/web-client/httpx"
)

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

var githubHttpClient = getGithubClient()

func getGithubClient() httpx.HttpClient {
	client := httpx.New()

	return client
}

func main() {
	getUrl()
	user := User{
		FirstName: "John",
		LastName:  "Doe",
	}
	createUser(user)
}

func getUrl() {
	res, err := githubHttpClient.Get("https://api.github.com", nil)
	if err != nil {
		log.Fatalf("error getting response: %v", err)
	}

	fmt.Println(res.StatusCode)

	b, _ := io.ReadAll(res.Body)
	fmt.Println(string(b))
}

func createUser(user User) {
	res, err := githubHttpClient.Post("https://api.github.com", nil, user)
	if err != nil {
		log.Fatalf("error getting response: %v", err)
	}

	fmt.Println(res.StatusCode)

	b, _ := io.ReadAll(res.Body)
	fmt.Println(string(b))
}
