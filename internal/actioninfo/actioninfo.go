package actioninfo

import (
	"log"
)

type DataParser interface {
	// TODO: добавить методы

	Parse(string) error

	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	// TODO: реализовать функцию
	for q := range dataset {
		er := dp.Parse(dataset[q])
		if er != nil {
			log.Println("Error parse", er)
			continue
		}
		dp.ActionInfo()
		if er != nil {
			log.Println("Error Action Info", er)
		}
		return
	}
}
