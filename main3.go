package main

import (
	"fmt"
	"strings"
)

type Car struct {
	ID        int
	Brand     string
	Model     string
	Year      int
	Price     float32
	Available bool
}

var c = 0

var cars = []Car{
	{ID: 1, Brand: "Toyota", Model: "Corolla", Year: 2020, Price: 250, Available: true},
	{ID: 2, Brand: "Honda", Model: "Civic", Year: 2019, Price: 700, Available: true},
	{ID: 3, Brand: "Ford", Model: "Focus", Year: 2018, Price: 600, Available: true},
	{ID: 4, Brand: "BMW", Model: "M5F90", Year: 2023, Price: 500, Available: true},
	{ID: 5, Brand: "Merc", Model: "G63", Year: 2022, Price: 300, Available: true},
	{ID: 6, Brand: "Nissan", Model: "GTR", Year: 2012, Price: 400, Available: true},
}

func (car *Car) Rent(days int) float32 {
	return car.Price * float32(days)
}

type car interface {
	Rent()
}

func home() {
	var userInput string
	fmt.Print("1 - Rent a Car\n2 - Return a Car\n3 - View Available Cars\n4 - Exit\n")
	fmt.Print("Choose an Option: ")
	fmt.Scanln(&userInput)
	switch userInput {
	case "1":
		rent()
	case "2":
		returnCar()
	case "3":
		availableCars()
	case "4":
		fmt.Println("exit..")
		return
	default:
		fmt.Println("Wrong bro")
		return
	}
}

func availableCars() {
	fmt.Println("Available Cars:")
	for _, car := range cars {
		if c == car.ID {
			continue
		} else {
			fmt.Printf("ID: %v, Brand: %s, Model: %s, Year: %d, Price: %v\n", car.ID, car.Brand, car.Model, car.Year, car.Price)

		}
	}
	fmt.Println("Do u want to go home 1 or 0 to exit")
	answer := ""
	fmt.Scanln(&answer)
	for {
		if answer == "1" {
			home()
		} else if answer == "0" {
			fmt.Println("exit...")
			return
		}
	}
}

func rent() {
	var userID int
	fmt.Println("Available Cars:")
	for _, car := range cars {
		if car.Available {
			fmt.Printf("ID: %v, Brand: %s, Model: %s, Year: %d, Price: %v\n", car.ID, car.Brand, car.Model, car.Year, car.Price)
		}
	}
	fmt.Print("Enter the ID of the car you want to rent: ")
	fmt.Scanln(&userID)

	for _, car := range cars {
		if car.ID == userID && car.Available {
			var days int
			fmt.Print("How many days would you like to rent the car? ")
			fmt.Scanln(&days)
			fmt.Printf("Total price for %d days: %v\n", days, car.Rent(days))
			fmt.Print("Do you want to  rent it? (yes/no) ")
			var deal string
			fmt.Scanln(&deal)
			deal = strings.ToLower(deal)
			if deal == "yes" {
				fmt.Println("Successfully rented the car.")
				car.Available = false
				c = userID

			} else {
				fmt.Println("Rental cancelled.")
				return
			}
		}
	}

	fmt.Println("Do u want to go home 1 or 0 to exit")
	answer := ""
	fmt.Scanln(&answer)
	for {
		if answer == "1" {
			home()
		} else if answer == "0" {

			return
		}
	}

}

func returnCar() {
	var idCard int
	fmt.Print("Enter the ID of the car you want to return: ")
	fmt.Scanln(&idCard)

	for i, car := range cars {
		if c == car.ID && car.Available == true {
			fmt.Println("Thank you for returning the car.")
			cars[i].Available = true

			break

		}
	}

	fmt.Println("Do u want to go home 1 or 0 to exit")
	answer := ""
	fmt.Scanln(&answer)
	for {
		if answer == "1" {
			home()
		} else if answer == "0" {
			return
		}
	}
}

func main() {
	home()
}
