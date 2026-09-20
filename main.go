package main

import (
	"fmt"

	realty "rashid/reality"
)

func main() {
	fmt.Println("REAL ESTATE PROPERTY LISTING SYSTEM")

	for {
		fmt.Println("\nMENU:")
		fmt.Println("1. List a new property")
		fmt.Println("2. Search property by ID")
		fmt.Println("3. Update listing price")
		fmt.Println("4. Delist a property")
		fmt.Println("5. Exit")

		var choice int
		fmt.Print("Enter your choice (1-5): ")
		fmt.Scan(&choice)

		switch choice {

		case 1:
			fmt.Println("\nList a New Property:")

			var id string
			var address string
			var price float64
			var bedrooms int
			var sold string

			fmt.Print("Enter Listing ID   : ")
			fmt.Scan(&id)

			fmt.Print("Enter Address      : ")
			fmt.Scan(&address)

			fmt.Print("Enter Price ($)    : ")
			fmt.Scan(&price)

			fmt.Print("Enter Bedrooms     : ")
			fmt.Scan(&bedrooms)

			fmt.Print("Is it sold? (y/n)  : ")
			fmt.Scan(&sold)

			isSold := false
			if sold == "y" || sold == "Y" {
				isSold = true
			}

			p := realty.Property{
				ListingID: id,
				Address:   address,
				Price:     price,
				Bedrooms:  bedrooms,
				IsSold:    isSold,
			}

			if err := realty.ListProperty(p); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Property", id, "listed successfully.")
			}

		case 2:
			fmt.Println("\nSearch Property:")

			var id string

			fmt.Print("Enter Listing ID to search: ")
			fmt.Scan(&id)

			p, err := realty.GetProperty(id)

			if err != nil {
				fmt.Println("Error:", err)
				continue
			}

			status := "Available"

			if p.IsSold {
				status = "Sold"
			}

			fmt.Println("\nProperty Details:")
			fmt.Println("Listing ID :", p.ListingID)
			fmt.Println("Address    :", p.Address)
			fmt.Printf("Price      : $%.2f\n", p.Price)
			fmt.Println("Bedrooms   :", p.Bedrooms)
			fmt.Println("Status     :", status)

		case 3:
			fmt.Println("\nUpdate Listing Price:")

			var id string
			var newPrice float64

			fmt.Print("Enter Listing ID    : ")
			fmt.Scan(&id)

			fmt.Print("Enter New Price ($) : ")
			fmt.Scan(&newPrice)

			if err := realty.UpdateListingPrice(id, newPrice); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Printf("Price of %s updated to $%.2f\n", id, newPrice)
			}

		case 4:
			fmt.Println("\nDelist Property:")

			var id string

			fmt.Print("Enter Listing ID to delist: ")
			fmt.Scan(&id)

			if err := realty.DelistProperty(id); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Property", id, "delisted successfully.")
			}

		case 5:
			fmt.Println("Thank you. Goodbye!")
			return

		default:
			fmt.Println("Invalid choice. Please enter a number from 1 to 5.")
		}
	}
}
