package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error)  {
	fmt.Print(prompt)

	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput("Create a new bill name: ", reader)

	b := newBill(name)
	fmt.Printf("Created bill: %s\n", b.name)

	return b
}

func promptOption(b bill) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Choose an option (a - add item, s - save bill, t - add tip): ", reader)

	switch opt {

	case "a":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item price: ", reader)

		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("price must be a number")
			promptOption(b)
		}

		b.addItem(name, p)

		fmt.Println("added - ", name, price)
			promptOption(b)

	case "t":
		tip, _ := getInput("Enter tip amount ($): ", reader)

		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("tip must be a number")
			promptOption(b)
		}

		b.updateTip(t)

		fmt.Println("added tip - ", tip)
			promptOption(b)

	case "s":
		b.save()
		fmt.Println("you saved the bill", b.name)
	default:
		fmt.Println("invalid option")
		promptOption(b)
	}
}

func main() {
	mybill := createBill()
	promptOption(mybill)
}