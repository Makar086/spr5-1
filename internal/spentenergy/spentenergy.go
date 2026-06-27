package spentenergy

import (
	"time"
	"errors"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps<=0||height<=0||weight<=0||duration<=0{return 0.00, errors.New("ошибка данных")}
	speed:=MeanSpeed(steps, height, duration)
	minut := float64(duration.Minutes())
	return (weight*speed*minut)/minInH*walkingCaloriesCoefficient, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps<=0||height<=0||weight<=0||duration<=0{return 0, errors.New("ошибка данных")}
	speed:=MeanSpeed(steps, height, duration)
	minut := float64(duration.Minutes())
	return (weight*speed*minut)/minInH, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if steps<=0||height<=0||duration<=0{return 0}
	dist:=Distance(steps, height)
	hours := float64(duration.Hours())
	return dist/hours
}

func Distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	if steps<=0||height<=0{return 0}
	lenstep:=height*stepLengthCoefficient
	return float64(steps)*lenstep/mInKm
}
