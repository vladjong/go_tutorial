package main

import (
	"log"
	"strings"
)

type StringManipulator func(string) string

func LogDecorator(fn StringManipulator) StringManipulator {
	return func(s string) string {
		log.Println("Начало исполнения со строкой:", s)

		result := fn(s)

		log.Println("Исполнение завершено с результатом:", result)

		return result
	}
}

// Реализация StringManipulator
func ToLower(s string) string {
	return strings.ToLower(s)
}

// Пример использования
func main() {
	s := "Привет, МИР!"
	fn := LogDecorator(ToLower)
	fn(s)
}

// Ожидаемый результат:
// 2024/01/05 17:38:04 Начало исполнения со строкой: Привет, МИР!
// 2024/01/05 17:38:04 Исполнение завершено с результатом: привет, мир!
