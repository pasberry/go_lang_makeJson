//Write a program which prompts the user to first enter a name, and then enter an address.
//Your program should create a map and add the name and address to the map using the keys “name” and “address”,
//respectively. Your program should use Marshal() to create a JSON object from the map, and then your program should
//print the JSON object.

//Submit your source code for the program, “makejson.go”.

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	
	fmt.Println("Please enter a name")

	inputScanner := bufio.NewScanner(os.Stdin)
	inputScanner.Scan()
	userInput := inputScanner.Text()

	jsonMap := make(map[string]string)

	jsonMap["name"] = userInput

	fmt.Println("Please enter the address")
	
	inputScanner.Scan()
	userInput = inputScanner.Text()
	
	jsonMap["address"] = userInput
	
	fmt.Println("Here is your json response")

	jsonData , _ := json.MarshalIndent(jsonMap, "","   ")

	fmt.Println(string(jsonData))
}
