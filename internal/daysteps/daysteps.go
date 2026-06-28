package daysteps

import (
	"errors"
	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
	"strconv"
	"strings"
	"time"
)

type DaySteps struct {
	// TODO: добавить поля
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	slice := strings.Split(datastring, ",")
	if len(slice) == 2 {
		step, err := strconv.Atoi(slice[0])
		if err != nil {
			return err
		}
		if step <= 0 {
			return errors.New("Error 0 steps")
		}
		ds.Steps = step

		time, err := time.ParseDuration(slice[1])
		if err != nil {

			return err
		}
		if time <= 0 {
			return errors.New("Error 0 Time")
		}

		ds.Duration = time
		return nil
	}
	return errors.New("Error")

}

func (ds DaySteps) ActionInfo() (string, error) {
	// TODO: реализовать функцию

	distan := spentenergy.Distance(ds.Steps, ds.Personal.Height)
	calore, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Personal.Weight, ds.Personal.Height, ds.Duration)
	if err != nil {
		return "", err
	}

	str := "Количество шагов: " + strconv.Itoa(ds.Steps) + ".\nДистанция составила " + strconv.FormatFloat(distan, 'f', 2, 64) + " км.\nВы сожгли " + strconv.FormatFloat(calore, 'f', 2, 64) + " ккал.\n"
	return str, nil

}
