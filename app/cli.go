package app

import (
	"fmt"
	"regexp"
	"strings"
)

type command struct {
	pattern     string
	description string
	handler     func([]string)
}

var commands []command

func init() {
	commands = []command{
		command{"users list", "users list --- lists all users", listUsers},
		command{"users create .+", "users create [username] --- creates a new user", createUser},
		command{"users resetpw .+", "users resetpw [username] --- generates a new password for [username]", resetPassword},
		command{".*", "help --- shows this menu", help},
	}
}

func HandleCommandLine(args []string) bool {
	if len(args) > 1 {
		command := parseCommand(args)
		command.handler(args)
		return true
	}

	return false
}

func parseCommand(args []string) command {
	helpCommand := commands[len(commands)-1]
	cmdLine := strings.Join(args[1:], " ")

	for _, cmd := range commands {
		if match, _ := regexp.MatchString(cmd.pattern, cmdLine); match {
			return cmd
		}
	}

	return helpCommand
}

func help(args []string) {
	appName := args[0]

	fmt.Printf("USAGE: %s [COMMAND] [OPTIONS]\n\n", appName)
	fmt.Println("COMMANDS:")
	fmt.Printf("%s --- runs the main application\n", appName)

	for _, cmd := range commands {
		fmt.Printf("%s %s\n", appName, cmd.description)
	}
}

func listUsers(args []string) {
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

func createUser(args []string) {
	username := args[3]

	if _, err := GetUser(username); err == nil {
		fmt.Println("ERROR: User already exists!")
		fmt.Println(err.Error())
		return
	}

	user := NewUser(username)
	if err := user.Save(); err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
	} else {
		fmt.Printf("Created user: %s\n", user.Name)
		showPassword(*user)
	}
}

func resetPassword(args []string) {
	username := args[3]
	user, err := GetUser(username)
	if err != nil {
		fmt.Printf("User %s not found", username)
		return
	}

	if err = user.ResetPassword(); err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
	} else {
		fmt.Printf("Reset password for user: %s\n", user.Name)
		showPassword(user)
	}
}

func showPassword(user User) {
	fmt.Printf("Password: %s\n\n", user.Password)
	fmt.Println("The password is not stored in plaintext and cannot be recovered. Be sure to copy it now.")
}
