package battle

import (
	"math/rand"

	"GoLang/4/models"

	"github.com/sirupsen/logrus"
)

const (
	MaxHeadsCut = 5
)

func Battle(hero *models.Hero, dragon *models.Dragon) {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	logrus.Info("Битва розпочата!")

	for i := 1; !dragon.IsWin() && !dragon.IsLoose(); i++ {
		headsCut := rand.Intn(MaxHeadsCut) + 1

		checkHeadsQuantity(dragon.Heads)

		dragon.Heads -= headsCut

		headsRegrown := 0
		for i := 0; i < headsCut; i++ {
			headsRegrown += growHead()
		}

		dragon.Heads += headsRegrown

		logrus.Infof("Богатир %s завдав %d удар, відрізав %d голів. Змій горинич відростив %d голів. Всього голів у Змія Горинича: %d.", hero.Name, i, headsCut, headsRegrown, dragon.Heads)

		if dragon.IsWin() {
			logrus.Fatal("У Змія Горинича більше 200 голів. Богатир лякається й панікує!")
			panic("Переміг Змій Горинич!")
		}
	}

	if dragon.IsLoose() {
		logrus.Info("Змій Горинич переможен! Вітаємо богатиря!")
	}
}

func growHead() int {
	chance := rand.Intn(100)
	switch {
	case chance < 35:
		return 0
	case chance < 70:
		return 1
	case chance < 90:
		return 2
	default:
		return 3
	}
}

func checkHeadsQuantity(h int) {
	switch h {
	case 111:
		logrus.Error("Юху! Змій Горинич дуже полюбляє число 111.")
		panic("Змій Горинич на радощах втік з 111 головами!")
	case 88:
		logrus.Error("Ой-ой! Змій Горинич боїться числа 88.")
		panic("Змій Горинич втік в паніці від числа 88!")
	}
}
