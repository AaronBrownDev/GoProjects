package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/AaronBrownDev/ContactManagementCLI/repository"
)

func main() {

	// If no arguments, gives instructions then exits program
	if len(os.Args) < 2 {
		fmt.Println("go run main.go help")
		os.Exit(1)
	}

	// Gets the content from the contacts.json. If fails, prints error then exits program.
	content, err := os.ReadFile("contacts.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Gets the specified operation from the arguments
	operation := os.Args[1]

	// Gets the repository that has the JSON Contact functions
	contactRepo := repository.GetJsonContactRepository(content)

	// Matches the given operation to the corresponding function.
	switch operation {
	case "add":
		// If there are not enough arguments to create a valid Contact then prints instructions then exits program.
		if len(os.Args) < 5 {
			fmt.Print("not enough arguments. Usage: go run main.go add <name> <phoneNumber> <emailAddress>")
			os.Exit(1)
		}
		// Creates contact with given arguments. If fails then outputs error then exits program.
		err := contactRepo.Create(os.Args[2], os.Args[3], os.Args[4])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

	case "list":
		// Gets all the contacts and the lists them
		contacts, err := contactRepo.GetAll()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println(contacts)

	case "find":
		// If there are not enough arguments to find a valid Contact then prints instructions then exits program.
		if len(os.Args) < 3 {
			fmt.Println("not enough arguments: Usage: go run main.go find <name>")
			os.Exit(1)
		}

		// Gets all the contacts with the same name
		contacts, err := contactRepo.GetByName(os.Args[2])
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(contacts)

	case "show":
		// converts the argument into an integer
		contactID, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid contactID. contactID has to be a positive integer. Usage: go run main.go delete <contactID>")
			os.Exit(1)
		}

		contact, err := contactRepo.GetByID(contactID)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println(contact)

	case "update":
		// Updates the contact of the same contactID and replaces it with the new contact
		// TODO: Have to find out how I want to implement this
		// contactRepo.Update()

	case "delete":
		// converts the argument into an integer
		contactID, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid contactID. contactID has to be a positive integer. Usage: go run main.go delete <contactID>")
			os.Exit(1)
		}
		// Deletes the contact from the contacts.json
		err = contactRepo.Delete(contactID)
		if err != nil {
			fmt.Println(err)
		}

	case "help":
		fmt.Println("Operations:\ngo run main.go add <name> <phoneNumber> <emailAddress>\ngo run main.go list\ngo run main.go find <name>\ngo run main.go show <contactID>\ngo run main.go update\ngo run main.go delete <contactID>")

	default:
		fmt.Println("Invalid operation.")
	}

}
