package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const serverPort = ":8080"

func main() {
	args := os.Args
	InitializeDb()

	if len(args) > 1 {
		handleCommandLine(args)
	} else {
		routes := []Route{
			Route{"PerformRedirect", "GET", "/{token}", RedirectShow},
			Route{"CreateRedirect", "POST", "/redirects", RedirectCreate},
		}

		router := NewRouter(routes)
		log.Fatal(http.ListenAndServe(serverPort, router))
	}
}

func handleCommandLine(args []string) {
	defer func() {
		if r := recover(); r != nil {
			help(args[0])
		}
	}()

	switch string(args[1]) {
	case "users":
		if args[2] == "list" {
			listUsers()
		} else if args[2] == "create" {
			createUser(args[3])
		}
	default:
		help(args[0])
	}

	fmt.Println()
}

func createUser(name string) {
	user := NewUser(name)
	if err := user.Save(); err != nil {
		fmt.Printf("There was an error: %s\n", err.Error())
	} else {
		fmt.Printf("Created User: %s\n", name)
		fmt.Println("Here's the password. This will never be shown again, so be sure to copy it")
		fmt.Printf("Password: %s\n", user.Password)
	}
}

func listUsers() {
	users, err := GetUsers()
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
		return
	}

	fmt.Println("USERS:")
	for _, user := range users {
		fmt.Println(user.Name)
	}
}

func help(appName string) {
	fmt.Printf("USAGE: %s [COMMAND] [OPTIONS]\n\n", appName)
	fmt.Println("COMMANDS:")
	fmt.Printf("%s --- runs the main application\n", appName)
	fmt.Printf("%s users create [username] --- creates a new user\n", appName)
	fmt.Printf("%s users list --- list existing users\n", appName)
	fmt.Printf("%s help --- shows this menu\n", appName)
}
