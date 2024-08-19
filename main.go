package main

import (
	"errors"
	"fmt"
	"strings"
)

type Client struct {
	Name           string
	Revenue        int
	ContractSigned bool
}

func FilterClientsByRevenue(clients []Client, minRevenue int) ([]Client, error) {
	if minRevenue < 0 {
		return nil, errors.New("revenue cannot be negative")
	}

var filtered []Client
	for _, client := range clients {
		if client.Revenue >= minRevenue {
			filtered = append(filtered, client)
		}
	}
	return filtered, nil
}


func FilterClientsBySignedStatus(clients []Client, signed bool) []Client {
	var filtered []Client
	for _, client := range clients {
		if client.ContractSigned == signed {
			filtered = append(filtered, client)
		}
	}
	return filtered
}

func main() {

	clients := []Client{
			{"Client A", 10000, true},
			{"Client B", 5000, false},
			{"Client C", 15000, true},
		}

		
		var filterType string

		fmt.Println("How would you like to filter the clients? (revenue/signed):")

		fmt.Scanln(&filterType)

		filterType = strings.ToLower(filterType) 

		switch filterType {

		case "revenue":

			var minRevenue int

			fmt.Println("Enter the minimum revenue:")

			fmt.Scanln(&minRevenue)

			filtered, err := FilterClientsByRevenue(clients, minRevenue)

			if err != nil {
				fmt.Println("Error:", err)
				return
			}

			fmt.Println("Clients with revenue greater than or equal to", minRevenue, ":") 

			for _, client := range filtered {
				fmt.Println(client.Name)
			}

		case "signed":

			var signedStatus string

			fmt.Println("Filter by signed contract? (yes/no):")

			fmt.Scanln(&signedStatus)

			signed := strings.ToLower(signedStatus) == "yes"

			filtered := FilterClientsBySignedStatus(clients, signed)

			fmt.Println("Clients with contract signed status", signed, ":")

			for _, client := range filtered {
				fmt.Println(client.Name)
			}

		default:
			fmt.Println("Invalid filter type. Please enter 'revenue' or 'signed'.")
		}
	}
	
	

	
	
	
	
	
	
	
	

