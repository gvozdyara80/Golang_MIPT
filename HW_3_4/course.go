package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	courseStudents := make(map[string]int)
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\nВыберите действие: \n1. Добавить студентов \n2. Показать количество студентов по курсам \n3. Завершить работу")
		scanner.Scan()
		input := scanner.Text()

		switch input {
		case "1":
			addStudents(courseStudents, scanner)
		case "2":
			showStudentCounts(courseStudents)
		case "3":
			fmt.Println("Работа программы завершена.")
			return
		default:
			fmt.Println("Некорректный ввод. Пожалуйста, выберите действие из предложенных.")
		}
	}
}

func addStudents(courseStudents map[string]int, scanner *bufio.Scanner) {
	fmt.Print("Введите курс и количество студентов через запятую (например, Информатика,5): ")
	scanner.Scan()
	input := scanner.Text()
	parts := strings.Split(input, ",")
	if len(parts) != 2 {
		fmt.Println("Некорректный ввод. Попробуйте снова.")
		return
	}
	course := strings.TrimSpace(parts[0])
	students, err := strconv.Atoi(strings.TrimSpace(parts[1]))
	if err != nil || students < 0 {
		fmt.Println("Количество студентов должно быть положительным числом. Попробуйте снова.")
		return
	}

	courseStudents[course] += students
	fmt.Println("Студенты успешно добавлены.")
}

func showStudentCounts(courseStudents map[string]int) {
	if len(courseStudents) == 0 {
		fmt.Println("Пока что студенты не были добавлены.")
		return
	}
	fmt.Println("Количество студентов по курсам:")
	for course, count := range courseStudents {
		fmt.Printf("%s: %d\n", course, count)
	}
}
