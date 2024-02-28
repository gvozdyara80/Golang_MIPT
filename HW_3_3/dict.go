package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	translations := make(map[string]string)
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\nВыберите действие: \n1. Добавить перевод \n2. Найти перевод \n3. Завершить работу")
		scanner.Scan()
		input := scanner.Text()

		switch input {
		case "1":
			addTranslation(translations, scanner)
		case "2":
			findTranslation(translations, scanner)
		case "3":
			fmt.Println("Работа программы завершена.")
			return
		default:
			fmt.Println("Некорректный ввод. Пожалуйста, выберите действие из предложенных.")
		}
	}
}

func addTranslation(translations map[string]string, scanner *bufio.Scanner) {
	fmt.Print("Введите слово и его перевод через запятую (например, cat,кот): ")
	scanner.Scan()
	input := scanner.Text()
	parts := strings.Split(input, ",")
	if len(parts) != 2 {
		fmt.Println("Некорректный ввод. Попробуйте снова.")
		return
	}
	word := strings.TrimSpace(parts[0])
	translation := strings.TrimSpace(parts[1])
	if word == "" || translation == "" {
		fmt.Println("Слово и перевод не могут быть пустыми. Попробуйте снова.")
		return
	}

	translations[word] = translation
	fmt.Println("Перевод добавлен успешно.")
}

func findTranslation(translations map[string]string, scanner *bufio.Scanner) {
	fmt.Print("Введите слово для поиска перевода: ")
	scanner.Scan()
	word := scanner.Text()
	translation, exists := translations[word]
	if !exists {
		fmt.Println("Перевод не найден.")
		return
	}
	fmt.Printf("Перевод слова \"%s\": %s\n", word, translation)
}
