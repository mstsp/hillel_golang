package main

import (
	"fmt"
	"strings"
)

func main() {
	// // 1 - Зробити конкатинацію рядків та помістити результат у змінну
	var userName string = "John"
	var userSername string = "Doe"
	var fullName = userName + " " + userSername

	fmt.Println(fullName)

	// // 2 - Зробити змінну результатом якої буде форматування рядків за допомогою fmt.Sprintf(...), використовуючи %s та %d, значення змінної вивести на екран
	newUser, userAge := "Joh Doe", 25
	infoMessage := fmt.Sprintf("New user is %s, %d", newUser, userAge)

	fmt.Println(infoMessage)

	// // 3 - Зробити математичні операції: +, -, *, / з будь якими числами результати помістити в змінні та вивести на екран
	numberOne, numberTwo := 10, 4

	var sum = numberOne + numberTwo
	fmt.Print(sum)

	var difference = numberTwo - numberOne
	fmt.Print(difference)

	var multiplication = numberOne * numberTwo
	fmt.Print(multiplication)

	var division = float32(numberOne) / float32(numberTwo)
	fmt.Println(division)

	// // 4 - Створити дві змінні типу float з різними значеннями і написати 4 умовні конструкції if використовуючи оператори:
	// >, >=, <, <=, ==, !=, &&, ||
	// також додати else до двух конструкцій, в тілі конструкцій виводити на екран ті операціїї які біли в умові

	var priceNet float64 = 3.10
	var priceGross float64 = 12.3

	if priceGross != priceNet && priceGross > priceNet {
		fmt.Println("priceGross != priceNet && priceGross > priceNet")
	}

	if priceGross == priceNet || priceGross < priceNet {
		fmt.Println("priceGross == priceNet || priceGross < priceNet")
	}

	if priceGross >= priceNet {
		fmt.Println("priceGross >= priceNet")
	} else {
		fmt.Println(" else priceGross >= priceNet")
	}

	if priceGross <= priceNet {
		fmt.Println("priceGross <= priceNet")
	} else {
		fmt.Println("else priceGross <= priceNet")
	}

	// 5 - Створити змінну типу string та за допомогою switch case і будь яких умов вивести ії значення на екран (має бути не менше трьох case)
	month := "March"

	switch month {
	case "December", "Jannuary", "February":
		fmt.Println(month, "is month of the winter")
	case "March", "April", "May":
		fmt.Println(month, "is month of the spring")
	case "June", "July", "August":
		fmt.Println(month, "is month of the summer")
	case "September", "October", "November":
		fmt.Println(month, "is month of the autumn")
	}

	// 6 - Створити ще один switch case, але без параметру switch
	// результатом має виводитись надписи з двох будь яких case стейтментів

	message := "Hello world"

	switch {

	case len(message) > 15:
		fmt.Println("Message has too long length. ")
		fallthrough
	case len(message) < 5:
		fallthrough
	case 5 >= len(message) || len(message) < 15:
		fmt.Println("Message has correct length. ")
		fallthrough
	case strings.Contains(message, "Hello"):
		fmt.Println("It is greeting message.")
	case strings.Contains(message, "Bye"):
		fmt.Println("It is goodbyes message.")
	}

	// 7 - Створити switch case з викроистанням дефолтного стейтменту, зробити вивід на екран будь чого з дефолтної частини блока

	switch {
	case strings.Contains(month, "April"):
		fmt.Println("It is April")
	case strings.Contains(month, "May"):
		fmt.Println("It is May")
	default:
		fmt.Println("It is default statement")
	}

	// 8 - Створити масив чисел з довжиною 5, інкрементувати елемент під індексом 3, декрементувати елемент під індиксом 4 - результат вивести на екран
	array := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array:", array)
	array[3]++
	fmt.Println("Incremented element with index 3: ", array)
	array[4]--
	fmt.Println("Decremented element with index 4: ", array)

	// 9 - Створити два слайси типа string двума способами (з використанням make та без використання make), додати в перший слайс рядок з пункту 1 та вивести на екран, також вивести кількість елементів з другого слайсу другим пареметром
	var sliceOne = make([]string, 0)
	var sliceTwo = []string{"one", "two", "three"}
	sliceOne = append(sliceOne, fullName)
	fmt.Println("First slice:", sliceOne, "length of the second slice:", len(sliceTwo))

	// У ході виконання домашнього завдання треба використовувати різні варіанти створення змінних (з указанням типу, без типа, з ключовим словом var, з := ), також використовувати вивід на екран fmt.Print() та fmt.Println()
	// Файл з кодом завантажити сюди, він має мати розширення .go. Або прикрепити лінк на гітхаб. Весь код писати в одному файлі.
}
