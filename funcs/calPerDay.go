package funcs

import (
	"fmt"
	"sportsman/models"
)

const (
	initialCal = 88.36
	weightKg   = 13.4
	heightCm   = 4.8
	ageY       = 5.7
)

// BMR = 88,36 + (13,4 × вес в кг) + (4,8 × рост в см) – (5,7 × возраст в годах).
func CaloriesPerDay() (result float32) {
	result = initialCal + (weightKg * models.CurrentHuman.Weight) + (heightCm * float32(models.CurrentHuman.Height)) + (ageY * float32(models.CurrentHuman.Age))
	fmt.Println("Сожжено по дефолту:", result)
	return
}
