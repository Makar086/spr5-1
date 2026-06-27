package trainings
import(
//	"fmt"
	"time"
	"errors"
	"strconv"
	"strings"
"github.com/Yandex-Practicum/tracker/internal/personaldata"
"github.com/Yandex-Practicum/tracker/internal/spentenergy"
		
)
type Training struct {
	// TODO: добавить поля
	personaldata.Personal
	Steps int
	TrainingType string
	Duration time.Duration
	
	
}

func (t *Training) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	slice := strings.Split(datastring, ",")
	if len(slice)==3{
	step, err := strconv.Atoi(slice[0])
	if err != nil {
			return err
			}
	if step<=0{
		return errors.New("ошибка 0 шагов")
		}
	t.Steps=step
	t.TrainingType=slice[1]
	time, err := time.ParseDuration(slice[2])
		if err != nil {
			
			return err
		}
	if time<=0{
		return errors.New("ошибка 0 время")
		}

	t.Duration =time
	return nil
	}
	return errors.New("ошибка")
}

func (t Training) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	hei:=t.Personal.Height
	distan:=spentenergy.Distance(t.Steps, hei)
	speed:=spentenergy.MeanSpeed(t.Steps, t.Personal.Height, t.Duration)




	switch t.TrainingType{
	case "Бег":
	calore, err:=spentenergy.RunningSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
if err != nil {
			
			
			return  "", err
		}
ti:=t.Duration.Hours()
str:="Тип тренировки: Бег\nДлительность: "+strconv.FormatFloat(ti, 'f', 2, 64)+" ч.\nДистанция: "+strconv.FormatFloat(distan, 'f', 2, 64)+" км.\nСкорость: "+strconv.FormatFloat(speed, 'f', 2, 64)+" км/ч\nСожгли калорий: "+strconv.FormatFloat(calore, 'f', 2, 64)+"\n"
return str, nil

	case "Ходьба":
	calore, err:=spentenergy.WalkingSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
if err != nil {
			
			
			return  "", err
		}
ti:=t.Duration.Hours()
str:="Тип тренировки: Ходьба\nДлительность: "+strconv.FormatFloat(ti, 'f', 2, 64)+" ч.\nДистанция: "+strconv.FormatFloat(distan, 'f', 2, 64)+" км.\nСкорость: "+strconv.FormatFloat(speed, 'f', 2, 64)+" км/ч\nСожгли калорий: "+strconv.FormatFloat(calore, 'f', 2, 64)+"\n"
return str, nil

	default:
return "", errors.New("неизвестный тип тренировки")
	}
}
