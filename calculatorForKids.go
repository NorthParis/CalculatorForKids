package main

import (
	"fmt"
)

func add(x, y float64) float64         { return x + y }
func subtraction(x, y float64) float64 { return x - y }
func multiply(x, y float64) float64    { return x * y }
func divide(x, y float64) (float64, error) {
	if y == 0 {
		return 0, fmt.Errorf("Нельзя делить на ноль")

	}
	return x / y, nil
}

func main() {
	fmt.Println("GOSHA IZOBRETAET CALCULATOR")
	for {
		var x, y float64
		fmt.Print("Введите x: ")
		if _, err := fmt.Scan(&x); err != nil {
			fmt.Println("Ошибка при вводе: ", err)
			continue
		}
		fmt.Print("Введите y: ")
		if _, err := fmt.Scan(&y); err != nil {
			fmt.Println("Ошибка при вводе: ", err)
			continue
		}
		var res float64
		var op int
		fmt.Println(`Введите операцию: 
		0 - Выход из программы
		1 - Сложение
		2 - Вычитание
		3 - Умножение
		4 - Деление`)
		if _, err := fmt.Scan(&op); err != nil {
			fmt.Println("Ошибка при вводе операции: ", err)
			continue
		}
		if op == 0 {
			fmt.Println("Выход. До встречи!")
			break
		}

		switch op {
		case 1:
			res = add(x, y)
		case 2:
			res = subtraction(x, y)
		case 3:
			res = multiply(x, y)
		case 4:
			var err error
			res, err = divide(x, y)
			if err != nil {
				fmt.Println("Ошибка: ", err)
				continue
			}
		default:
			fmt.Println("Неверная операция")
			continue
		}

		fmt.Printf("Результат: %.5f\n", res)
	}
}
