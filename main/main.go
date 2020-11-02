package main

import (
	"bufio"
	"fmt"
	"os"
)

type Client struct {
	clientId string
	name     string
	surname  string
}

func (c *Client) addClient(clientId string, name string, surname string) {
	c.clientId = clientId
	c.name = name
	c.surname = surname
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter name: ")
	text, _ := reader.ReadString('\n')
	fmt.Print("Enter surname: ")
	text1, _ := reader.ReadString('\n')
	text2, _ := reader.ReadString('\n')

	fmt.Println("Hello")
	fmt.Println("Hi")

	text11, _ := reader.ReadString('\n')
	if text11 == "Hello" {
		fmt.Printf("%v, Hello")
	} else {
		fmt.Println("Hi")
	}

	client := Client{
		clientId: text,
		name:     text1,
		surname:  text2,
	}

	fmt.Printf("Hello %v %v %v", client.clientId, client.name, client.surname)
}
