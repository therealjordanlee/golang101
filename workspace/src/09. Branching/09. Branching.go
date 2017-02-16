// Go has 2 types of branching statements: IF and SWITCH

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Read from commandline using the bufio library.
	// We create a Reader object...
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("What's 1 + 1?")

	// and then call the ReadString method
	input, _ := reader.ReadString('\n')

	// TrimSpace returns a slice of the string s
	// with all leading and trailing white space removed
	input = strings.TrimSpace(input)

	// Convert the input into an integer.
	// Invalid text (non-numbers) will be converted to '0'
	var inputInt int
	inputInt, _ = strconv.Atoi(input)

	// If statements in go follow the usual syntax:
	// if [condition] {
	// } else if [condition] {
	// } else {
	// }
	if int(inputInt) == 2 {
		fmt.Println("Correct!")
	} else if inputInt > 0 && inputInt < 4 {
		fmt.Println("Close!")
	} else {
		fmt.Println("Wrong!")
	}
}
