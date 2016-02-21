package main

import (
	"fmt"
	"log"

	"./client/riot"
)

func main() {
	client, err := riot.NAClient()
	if err != nil {
		log.Panic(err)
	}

	for _, c := range client.Champions {
		fmt.Printf("id: %d,  name: %s", c.ID, c.Name)
	}

	fmt.Println(len(client.Champions))
}
