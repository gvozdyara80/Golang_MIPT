package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Структура для хранения данных о денежных потоках
type Budjet struct {
	Income float64
	Costs  map[string]float64
}

// Добавление дохода
func (b *Budjet) AddIncome(income float64) {
	b.Income += income
}

// Добавление расходов по категориям
func (b *Budjet) AddCost(category string, cost float64) {
	if b.Costs == nil {
		b.Costs = make(map[string]float64)
	}
	b.Costs[category] += cost
}

// Расчет чистого дохода
func (b *Budjet) CalculateBalance() float64 {
	totalCosts := 0.0
	for _, cost := range b.Costs {
		totalCosts += cost
	}
	return b.Income - totalCosts
}

// Анализ расходов
func (b *Budjet) AnalyzeCosts() {
	for category, cost := range b.Costs {
		ratio := (cost / b.Income) * 100
		if ratio > 30 {
			fmt.Printf("Внимание: расходы на %s составляют %.0f%% от вашего дохода.\nРекомендуется снизить траты в этой категории.\n", category, ratio)
		}
	}
}

func main() {
	budjet := Budjet{}
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Ввведите ваш месячный доход:") // Дополнить обработкой ошибки, когда пользователь вводит не число
	scanner.Scan()
	income, _ := strconv.ParseFloat(scanner.Text(), 64)
	budjet.AddIncome(income)

	fmt.Println("Введите ваши расходы в формате \"Категория расходов: размер расходов\" (например, Продукты: 7000). После завершения ввода всех расходов напишите \"готово\":")
	for scanner.Scan() {
		input := scanner.Text()
		if input == "готово" {
			break
		}
		parts := strings.Split(input, ": ")
		if len(parts) != 2 {
			fmt.Println("Вы ввели данные в неверном формате. Пожалуйста, попробуйте еще раз.")
			continue
		}
		cost, err := strconv.ParseFloat(parts[1], 64)
		if err != nil {
			fmt.Println("Вы ввели сумму расходов, которая не явлется числом. Пожалуйста, попробуйте еще раз.")
			continue
		}
		budjet.AddCost(parts[0], cost)
	}

	balance := budjet.CalculateBalance()
	fmt.Printf("Ваш чистый доход: %.0f\n", balance)

	budjet.AnalyzeCosts()
}
