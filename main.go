package main

import (
	"errors"
	"fmt"
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
	filtered, _ := FilterClientsByRevenue(clients, 7000)
	for _, client := range filtered {
		fmt.Println(client.Name)
	}
	signedClients := FilterClientsBySignedStatus(clients, true)
	for _, client := range signedClients {
		fmt.Println(client.Name)
	}
}
