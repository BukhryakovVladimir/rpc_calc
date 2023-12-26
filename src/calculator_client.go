package main

import (
	"fmt"
	"net/rpc"
)

type Args struct {
	A, B float64
}

func main() {
	// Add
	additionClient, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		fmt.Println("Error:", err)
	}
	defer additionClient.Close()

	// Substract
	substractionClient, err := rpc.Dial("tcp", "localhost:1235")
	if err != nil {
		fmt.Println("Error:", err)
	}
	defer substractionClient.Close()

	// Multiply
	multiplicationClient, err := rpc.Dial("tcp", "localhost:1236")
	if err != nil {
		fmt.Println("Error:", err)
	}
	defer multiplicationClient.Close()

	// Divide
	divisionClient, err := rpc.Dial("tcp", "localhost:1237")
	if err != nil {
		fmt.Println("Error:", err)
	}
	defer divisionClient.Close()

	var choice uint8
	var A float64
	var B float64
	var result float64

	for choice != 5 {
		fmt.Println("\nВыберите действие:\n1. Сложение\n2. Вычитание\n3. Умножение\n4. Деление\n5. Выход")
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println(err)
			continue
		}

		if choice == 5 {
			return
		}

		if choice > 5 || choice < 1 {
			fmt.Println("\nТакого действия нет. Выберите существующее действие.")
			continue
		}

		for {
			fmt.Println("Введите 1-ое число: ")
			_, err := fmt.Scan(&A)
			if err != nil {
				fmt.Println(err)
				continue
			}
			break
		}

		for {
			fmt.Println("Введите 2-ое число: ")
			_, err := fmt.Scan(&B)
			if err != nil {
				fmt.Println(err)
				continue
			}
			break
		}

		args := Args{A, B}

		switch choice {
		case 1:
			err = additionClient.Call("MathService.Add", args, &result)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("Result: ", result)
		case 2:
			err = substractionClient.Call("MathService.Substract", args, &result)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("Result: ", result)
		case 3:
			err = multiplicationClient.Call("MathService.Multiply", args, &result)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("Result: ", result)
		case 4:
			if args.B != 0 {
				err = divisionClient.Call("MathService.Divide", args, &result)
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println("Result: ", result)
			} else {
				fmt.Println("Error: Can't divide by 0")
			}
		default:
			return
		}
	}
}
