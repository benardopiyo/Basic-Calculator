package calculator

import (
	"fmt"
	"os"
)

func Calc() {
	fmt.Println()
	fmt.Println("***********************")
	fmt.Println()
	fmt.Println("Welcome to Zone01 Calculator")
	fmt.Println()
	fmt.Println("****** MAIN MENU ******")
	fmt.Println()
	fmt.Println("1. Add")
	fmt.Println("2. Substract")
	fmt.Println("3. Multiply")
	fmt.Println("4. Divide")
	fmt.Println()
	fmt.Println("***********************")
	fmt.Println()
	fmt.Println("Enter choice:")
	fmt.Println()

	var choice int
	fmt.Scanln(&choice)
	fmt.Println()

	if choice <= 0 || choice >= 5 {
		fmt.Println("Error: enter correct number")
		fmt.Println()
		os.Exit(1)
		fmt.Println()
		fmt.Println("***********************")
	}

	var x, y int

	fmt.Println()
	fmt.Println("Enter the first number x:")
	fmt.Scanln(&x)
	fmt.Println()
	fmt.Println("Enter the second number y:")
	fmt.Scanln(&y)
	fmt.Println()

	if choice == 1 {
		ans := x + y
		fmt.Println("Answer =", ans)
		fmt.Println()
		fmt.Println("***********************")
	} else if choice == 2 {
		ans := x - y
		fmt.Println("Answer =", ans)
		fmt.Println()
		fmt.Println("***********************")
	} else if choice == 3 {
		ans := x * y
		fmt.Println("Answer =", ans)
		fmt.Println()
		fmt.Println("***********************")
	} else if choice == 4 {
		ans := x / y
		fmt.Println("Answer =", ans)
		fmt.Println()
		fmt.Println("***********************")
	} else {
		fmt.Println()
		fmt.Println("Enter the correct choice")
		fmt.Println()
		fmt.Println("***********************")
	}
}
