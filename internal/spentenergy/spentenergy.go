package spentenergy

import (
	"errors"
	"time"
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
	if steps <= 0 {
		return 0, errors.New("error steps")
	}
	if height <= 0 {
		return 0, errors.New("error height")
	}
	if weight <= 0 {
		return 0, errors.New("error weight")
	}
	if duration <= 0 {
		return 0, errors.New("error time")
	}

	speed := MeanSpeed(steps, height, duration)
	return (weight * speed * duration.Minutes()) / minInH * walkingCaloriesCoefficient, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 {
		return 0, errors.New("error steps")
	}
	if height <= 0 {
		return 0, errors.New("error height")
	}
	if weight <= 0 {
		return 0, errors.New("error weight")
	}
	if duration <= 0 {
		return 0, errors.New("error time")
	}
	speed := MeanSpeed(steps, height, duration)
	return (weight * speed * duration.Minutes()) / minInH, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if steps <= 0 || height <= 0 || duration <= 0 {
		return 0
	}
	dist := Distance(steps, height)
	return dist / duration.Hours()
}

func Distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	if steps <= 0 || height <= 0 {
		return 0
	}
	lenstep := height * stepLengthCoefficient
	return float64(steps) * lenstep / mInKm
}
