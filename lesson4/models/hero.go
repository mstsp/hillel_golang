package models

import "math/rand"

type Hero struct {
	Name string
}

func NewHero() *Hero {
	return &Hero{Name: generateRandomName()}
}

func generateRandomName() string {
	firstNames := []string{"Скрипучий", "Чубатий", "Найзатхліший", "Невгамовний", "Заплутобородий", "Дзвінкоголосий"}
	secondNames := []string{"Сивий", "Румяний", "Пиріжкоголовий", "Випивший", "Щебетун", "Хвалько"}
	thirdNames := []string{"Пиріжок", "Кінь", "Дерево", "Пузан", "Ведмеже Вухо"}

	return firstNames[rand.Intn(len(firstNames))] + " " + secondNames[rand.Intn(len(secondNames))] + " " + thirdNames[rand.Intn(len(thirdNames))]
}
