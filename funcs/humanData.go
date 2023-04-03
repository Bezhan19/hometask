package funcs

import (
	"fmt"
	"log"
	"sportsman/models"
)

func EnterHumanData() error {
	_, err := fmt.Scan(&models.CurrentHuman.Weight, &models.CurrentHuman.Height, &models.CurrentHuman.Age)
	if err != nil {
		log.Println(InsertedDataErr)
		return err
	}
	return nil
}
