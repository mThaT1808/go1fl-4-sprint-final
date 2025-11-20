package spentcalories

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	// Разделить строку на слайс строк
	parts := strings.Split(data, ",")
	// Проверить, чтобы длина слайса была равна 3, так как в строке данных у нас количество шагов, вид активности и продолжительность.
	if len(parts) != 3 {
		return 0, "", 0, fmt.Errorf("ошибка при разделении строки. Ожидается формат: '3456,Ходьба,3h00m'")	
	}
	// Преобразовать первый элемент слайса (количество шагов) в тип int. 
	steps, err := strconv.Atoi(parts[0])
	// Обработать возможные ошибки. При их возникновении из функции вернуть 0 шагов, 0 продолжительность и ошибку.
	if err != nil {
		return 0, "", 0, fmt.Errorf("ошибка при преобразовании шагов: %v", err)
	}
	activityType := parts[1]
	// Преобразовать третий элемент слайса в time.Duration. В пакете time есть метод для парсинга строки в time.Duration.
	duration, err := time.ParseDuration(parts[2])
	// Обработать возможные ошибки. При их возникновении из функции вернуть 0 шагов, 0 продолжительность и ошибку.
	if err != nil {
		return 0, "", 0, fmt.Errorf("ошибка при преобразовании продолжительности: %v", err)
	}
	// Если всё прошло без ошибок, верните количество шагов, вид активности, продолжительность и nil (для ошибки).
	return steps, activityType, duration, nil
}

func distance(steps int, height float64) float64 {
	// рассчитайте длину шага. Для этого умножьте высоту пользователя на коэффициент длины шага stepLengthCoefficient.
	step := height * stepLengthCoefficient
	// умножьте пройденное количество шагов на длину шага. разделите полученное значение на число метров в километре
	distance := (float64(steps) * step) / float64(mInKm)
	return distance
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	return 0
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
	return "", nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	return 0, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	return 0, nil
}
