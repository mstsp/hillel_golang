package main

import (
	"GoLang/4/battle"
	"GoLang/4/models"

	"github.com/sirupsen/logrus"
)

func main() {
	hero := models.NewHero()
	dragon := models.NewDragon()

	defer func() {
		if r := recover(); r != nil {
			logrus.Errorf("Гра завершилася панікою: %v", r)
		}
	}()

	battle.Battle(hero, dragon)
}
