package main

import (
	"fmt"
	"sportsman/funcs"
)

func main() {
	fmt.Println("Данные о: вес,рост,возраст")
	err := funcs.EnterHumanData()
	if err != nil {
		return
	}
	fmt.Println("Пробег:")
	var runData float32
	var yearCalorie float32
	fmt.Scan(&runData)
	burned := funcs.BurnedCalories(runData)
	initiallyBurned := funcs.CaloriesPerDay()
	eaten, err := funcs.DailyEaten()
	dif := eaten - burned - initiallyBurned
	if dif > 0 {
		fmt.Println("Мы толстеем с перееданием в %v калорий", dif)
		yearCalorie = dif * 365
		fmt.Print("количество колории прибавилоь за год :", yearCalorie)
	}
	if dif < 0 {
		fmt.Println("Мы худеем с недоеданием в %v калорий", dif)
		yearCalorie = -dif * 365

		fmt.Print("количество колории потеряно за год :", yearCalorie)

	}
	if dif == 0 {
		fmt.Printf("Всё отлично, так держать!")
	}
}

