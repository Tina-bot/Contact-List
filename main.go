package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var reader *bufio.Reader //declare it global -
var id int

type User struct {
	id       int
	username string
	email    string
	age      int
}

var users map[int]User

func createUser() {
	clearConsole()
	fmt.Println("Username: ")
	username := readLine()
	fmt.Println("Email: ")
	email := readLine()
	fmt.Println("Age: ")
	age, err := strconv.Atoi(readLine()) //conver int to str
	if err != nil {
		panic("Could not convert str to int")
	}
	id++
	user := User{id, username, email, age}
	users[id] = user
	fmt.Println(" User created successfully! ")
	fmt.Println('\n')
}

func listUsers() {
	clearConsole()
	fmt.Println("User list :")
	for id, user := range users {
		fmt.Println(id, "Username: ", user.username, "Email: ", user.email, "Age: ", user.age)
	}
	fmt.Println('\n')
}

func updateUser() {
	clearConsole()

	fmt.Print("ID: ")
	if id, err := strconv.Atoi(readLine()); err != nil {
		panic("Could not convert ID to int")
	} else {
		if _, ok := users[id]; ok {
			fmt.Print("New Username: ")
			username := readLine()
			fmt.Print("New Email: ")
			email := readLine()
			fmt.Print("New Age: ")
			if age, err := strconv.Atoi(readLine()); err != nil {
				panic("Could not convert age to int")
			} else {
				user := User{id, username, email, age}
				users[id] = user
			}
		}
	}

	fmt.Println("User updated successfully")
	fmt.Println('\n')

}

func deleteUser() {
	clearConsole()
	id, err := strconv.Atoi(readLine())
	if err != nil {
		panic(" Could not convert str to int ")
	}
	if _, ok := users[id]; ok {
		delete(users, id)
	}

	fmt.Println("User deleted successfully")
	fmt.Println('\n')
}

func clearConsole() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func readLine() string {

	if option, err := reader.ReadString('\n'); err != nil {
		panic("Error reading option")
	} else {

		return strings.TrimSpace(option) //convert the writing into a new string eliminating the jump
	}
}

func main() {

	var option string
	users = make(map[int]User)
	reader = bufio.NewReader(os.Stdin)

	for {
		fmt.Println("a) Create")
		fmt.Println("B) List")
		fmt.Println("C) Update")
		fmt.Println("D) Delete")

		fmt.Println("enter an option: ")
		option = readLine()

		if option == "quit" || option == "q" {
			break
		}

		switch option {
		case "a", "create":
			createUser()
		case "b", "list":
			listUsers()
		case "c", "3", "update":
			updateUser()
		case "d", "4", "delete":
			deleteUser()

		default:
			fmt.Println("Unknown option")
		}
	}

	fmt.Println("Goodbye")
}
