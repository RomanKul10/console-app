package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/eiannone/keyboard"
)

func main() {
	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = keyboard.Close()
	}()

	coffees := make(map[int]string)
	coffees[1] = "Cappucino"
	coffees[2] = "Latte"
	coffees[3] = "Americano"
	coffees[4] = "Mocha"
	coffees[5] = "Macchiato"
	coffees[6] = "Espresso"

	fmt.Println("MENU")
	fmt.Println("----")
	fmt.Println("1 - Cappuchino")
	fmt.Println("2 - Latte")
	fmt.Println("3 - Americano")
	fmt.Println("4 - Mocha")
	fmt.Println("5 - Macchiato")
	fmt.Println("6 - Espresso")
	fmt.Println("Q - Quit the program")

	/*for {
		char, _, err := keyboard.GetKey()
		if err != nil {
			log.Fatal(err)
		}
		if char == 'q' || char == 'Q' {
			break
		}

		i, _ := strconv.Atoi(string(char))

		/*_, ok := coffees[i]
		if ok {
			fmt.Println(fmt.Sprintf("You chose %s", coffees[i]))
		} // наступне це теж саме, що і це

		if _, ok := coffees[i]; ok {
			fmt.Println(fmt.Sprintf("You chose %s", coffees[i]))
		}

	} */
	char := ' '
	for ok := true; ok; ok = char != 'q' && char != 'Q' {
		char, _, err = keyboard.GetKey()
		if err != nil {
			log.Fatal(err)
		}
		i, _ := strconv.Atoi(string(char))

		if _, ok := coffees[i]; ok {
			fmt.Println(fmt.Sprintf("You chose %s", coffees[i]))
		}

	}

	fmt.Println("Program exiting...")
}
