package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user Input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our pizza: ")


	// comma ok || comma err

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ",input)
	fmt.Printf("The type of this rating is %T", input)
}
