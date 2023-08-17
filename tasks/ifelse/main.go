package main

import "fmt"

func main() {
	Ticket("AP", "F", 4, 75)
	Ticket("Delhi", "F", 8, 120)
	Ticket("UP", "M", 6, 130)
	Ticket("Karnataka", "M", 21, 172)
	Ticket("Kerla", "M", 4, 80)
}

func Ticket(state, gender string, age, height int) {

	if state == "UP" {
		if (gender == "F" && height < 120 && age < 6) || (gender == "M" && height < 120 && age < 6) {
			fmt.Println("No Ticket")
		} else {
			fmt.Println("Full Ticket")
		}
	} else if state == "Delhi" {
		if (gender == "F") || (gender == "M" && height < 130 && age < 7) {
			fmt.Println("No Ticket")
		} else {
			fmt.Println("Full Ticket")
		}
	} else if state == "AP" {
		if (gender == "F" && height < 110 && age < 5) || (gender == "M" && height < 110 && age < 5) {
			fmt.Println("No Ticket")
		} else {
			fmt.Println("Full Ticket")
		}

	} else if state == "Karnataka" {
		if (gender == "F") || (gender == "M" && height < 110 && age < 5) {
			fmt.Println("No Ticket")
		} else {
			fmt.Println("Full Ticket")
		}
	} else { // For other States
		fmt.Println("Full Ticket")
	}

}
