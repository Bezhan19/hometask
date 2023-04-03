package funcs

import (
	"fmt"
	"log"
)

const InsertedDataErr = "Ошибка при вводе данных"

func DailyEaten() (eatenCalories float32, err error) {
	fmt.Println("Введите сколько съедено за сегодня: ")
	_, err = fmt.Scan(&eatenCalories)
	if err != nil {
		log.Println(InsertedDataErr)
		return
	}
	return
}
