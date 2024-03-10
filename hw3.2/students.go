package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Student struct {
	Name    string
	Surname string
	Age     int
	Grades  []float64
}

var students []Student

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Выберите действие:\n'1' - Добавить студента\n'2' - Вывести список студентов\n'3'- Завершить работу:")
		scanner.Scan()
		input := scanner.Text()

		switch input {
		case "1":
			addStudent(scanner)
		case "2":
			printStudents()
		case "3":
			fmt.Println("Программа завершена.")
			return
		default:
			fmt.Println("Некорректный ввод. Пожалуйста, выберите действие из предложенных.")
		}
	}
}

func addStudent(scanner *bufio.Scanner) {
	var student Student

	fmt.Print("Введите имя студента: ")
	for scanner.Scan() {
		input := scanner.Text()
		if input == "" {
			fmt.Print("Ошибка: имя не может быть пустым. Введите имя снова: ")
			continue
		}
		student.Name = input
		break
	}

	fmt.Print("Введите фамилию студента: ")
	for scanner.Scan() {
		input := scanner.Text()
		if input == "" {
			fmt.Print("Ошибка: фамилия не может быть пустой. Введите фамилию снова: ")
			continue
		}
		student.Surname = input
		break
	}

	fmt.Print("Введите возраст студента: ")
	for scanner.Scan() {
		input := scanner.Text()
		age, err := strconv.Atoi(input)
		if age < 0 || err != nil {
			fmt.Print("Ошибка: возраст должен быть неотрицательным числом. Введите возраст снова: ")
			continue
		}
		student.Age = age
		break
	}

	fmt.Print("Введите оценки студента через запятую (например, 4.5,3.7,5): ")
	for scanner.Scan() {
		input := scanner.Text()
		strGrades := strings.Split(input, ",")
		var grade float64
		var grades []float64
		var err error
		for _, strGrade := range strGrades {
			grade, err = strconv.ParseFloat(strGrade, 64)
			if grade < 0 || err != nil {
				break
			}
			grades = append(grades, grade)
		}
		if grade < 0 || err != nil {
			fmt.Print("Ошибка: оценки должны быть неотрицательными числами. Введите оценки снова: ")
			continue
		}
		student.Grades = grades
		break
	}

	students = append(students, student)
}

func printStudents() {
	if len(students) == 0 {
		fmt.Println("Список студентов пуст.")
		return
	}
	for i, student := range students {
		fmt.Printf("%d. Имя и фамилия: %s %s, Возраст: %d, Оценки: %v\n", i+1, student.Name, student.Surname, student.Age, student.Grades)
	}
}
