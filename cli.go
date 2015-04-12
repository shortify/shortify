package main

import (
	"fmt"
	"regexp"
	"strings"
)

type CLICommand struct {
	Pattern     string
	Description string
	Handler     func([]string)
}

var commands []CLICommand

func init() {
	commands = []CLICommand{
		CLICommand{"users list", "users list --- lists all users", listUsers},
		CLICommand{"users create .+", "users create [username] --- creates a new user", createUser},
		CLICommand{"users resetpw .+", "users resetpw [username] --- generates a new password for [username]", resetPassword},
		CLICommand{".*", "help --- shows this menu", help},
	}
}

func GetCLICommand(args []string) CLICommand {
	helpCommand := commands[len(commands)-1]
	cmdLine := strings.Join(args[1:], " ")

	for _, cmd := range commands {
		if match, _ := regexp.MatchString(cmd.Pattern, cmdLine); match {
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
		fmt.Printf("%s %s\n", appName, cmd.Description)
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
