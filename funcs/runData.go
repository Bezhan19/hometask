package funcs

import (
	"fmt"
	"sportsman/models"
)

// BurnedCalories - возвращает колич. созжённых калорий в день
func BurnedCalories(distanceKm float32) float32 {
	calories := distanceKm * models.CurrentHuman.Weight
	fmt.Println("Сожжено за день:", calories)
	return calories
}
