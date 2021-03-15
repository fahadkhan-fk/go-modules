package main

import (
	cities "cities" // this `cities` is a seperate module who's functions are being called in this module ( you need to see mod.go file )
	"fmt"
)

func main() {
	fmt.Println("Showing list of cities of different countries")

	countries := []string{"pakistan", "bangladesh", "india"}

	// for loop just to iterate list of countries and get their cities
	for _, country := range countries {
		if country == "india" {
			fmt.Println("List of cities from India:")
			fmt.Print(cities.GetIndiaCitiesList()) // here we have called the funtion GetIndiaCitiesList() that's present in cities module
		}
		if country == "bangladesh" {
			fmt.Println("List of cities from Bangladesh:")
			fmt.Println(cities.GetBangladeshCitiesList()) // here we have called the funtion GetIndiaCitiesList() that's present in cities module
		}
		if country == "pakistan" {
			fmt.Println("List of cities from Pakistan:")
			fmt.Println(cities.GetPakistanCitiesList()) // here we have called the funtion GetIndiaCitiesList() that's present in cities module
		}
	}
}
