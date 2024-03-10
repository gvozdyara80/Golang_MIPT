package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	slice := make([]string, 0)

	fmt.Print("Введите строку: ")
	for scanner.Scan() {
		input := scanner.Text()
		if input == "" {
			fmt.Print("Ошибка: ввод не может быть пустым. Попробуйте еще раз: ")
			continue
		}
		for _, char := range input {
			slice = append(slice, string(char))
		}
		break
	}

	fmt.Print(slice)
}
