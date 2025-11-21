package spentcalories

import (
	"fmt"
	"log"
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
	if steps <= 0 {
		return 0, "", 0, fmt.Errorf("ошибка при преобразовании шагов: шагов должно быть больше нуля")
	}
	activityType := parts[1]
	// Преобразовать третий элемент слайса в time.Duration. В пакете time есть метод для парсинга строки в time.Duration.
	duration, err := time.ParseDuration(parts[2])
	// Обработать возможные ошибки. При их возникновении из функции вернуть 0 шагов, 0 продолжительность и ошибку.
	if err != nil {
		return 0, "", 0, fmt.Errorf("ошибка при преобразовании продолжительности: %v", err)
	}
	if duration <= 0 {
		return 0, "", 0, fmt.Errorf("ошибка при преобразовании продолжительности: продолжительность должна быть больше нуля")
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
	// Проверить, что продолжительность duration больше 0. Если это не так, вернуть 0.
	if duration <= 0 {
		return 0
	}
	// Вычислить дистанцию с помощью distance
	distance := distance(steps, height)
	// Вычислить и вернуть среднюю скорость.
	// Для этого разделите дистанцию на продолжительность в часах. 
	// Чтобы перевести продолжительность в часы, воспользуйтесь функцией из пакета time.
	speed := distance / duration.Hours()
	return speed
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// Получить значения из строки данных с помощью функции parseTraining().
	steps, activityType, duration, err := parseTraining(data)
	// Обработать возможные ошибки и вывести их в лог с помощью log.Println(err).
	if err != nil {
		log.Println(err)
	}
	// Проверить, какой вид тренировки был передан в строке, которую парсили (лучше использовать switch).
	// Для каждого из видов тренировки рассчитать дистанцию, среднюю скорость и калории.
	calories := float64(0)
	switch activityType {
		case "Бег": {
			calories, err = RunningSpentCalories(steps, weight, height, duration)
			if err != nil {
				return "", fmt.Errorf("ошибка при определении типа тренеровки: %v", err)
			}
		}
		case "Ходьба": {
			calories, err = WalkingSpentCalories(steps, weight, height, duration)
			if err != nil {
				return "", fmt.Errorf("ошибка при определении типа тренеровки: %v", err)
			}
		} 
		// Если был передан неизвестный тип тренировки, вернуть ошибку с текстом неизвестный тип тренировки.
		default: 
			return "", fmt.Errorf("ошибка при определении типа тренеровки: неизвестный тип тренировки")
	}
	distance := distance(steps, height)
	speed := meanSpeed(steps, height, duration)
	// Для каждого вида тренировки сформировать и вернуть строку, образец которой был представлен выше.
	message := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", activityType, duration.Hours(), distance, speed, calories)

	return message, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// Проверить входные параметры на корректность. 
	// Если параметры некорректны, вернуть 0 калорий и соответствующую ошибку.
	if steps <= 0 {
		return 0, fmt.Errorf("шагов должно быть больше нуля")
	}
	if weight <= 0 {
		return 0, fmt.Errorf("вес быть больше нуля")
	}
	if height <= 0 {
		return 0, fmt.Errorf("рост должен быть больше нуля")
	}
	if duration <= 0 {
		return 0, fmt.Errorf("длительность должна быть больше нуля")
	}
	// Рассчитать среднюю скорость с помощью meanSpeed().
	speed := meanSpeed(steps, height, duration)
	// Рассчитать и вернуть количество калорий.
	// Переведите продолжительность в минуты с помощью функции из пакета time.
	durationMinutes := duration.Minutes()
	// Умножьте вес пользователя на среднюю скорость и продолжительность в минутах.
	// Разделите результат на число минут в часе для получения количества потраченных калорий. (weight * meanSpeed * durationInMinutes) / minInH
	spentCalories := (weight * speed * durationMinutes) / minInH
	return spentCalories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// Проверить входные параметры на корректность. Если параметры некорректны, вернуть 0 калорий и соответствующую ошибку.
	if steps <= 0 {
		return 0, fmt.Errorf("шагов должно быть больше нуля")
	}
	if weight <= 0 {
		return 0, fmt.Errorf("вес быть больше нуля")
	}
	if height <= 0 {
		return 0, fmt.Errorf("рост должен быть больше нуля")
	}
	if duration <= 0 {
		return 0, fmt.Errorf("длительность должна быть больше нуля")
	}
	// Рассчитать среднюю скорость с помощью meanSpeed().
	speed := meanSpeed(steps, height, duration)
	// Рассчитать количество калорий.
	// Переведите продолжительность в минуты с помощью функции из пакета time.
	durationMinutes := duration.Minutes()
	// Умножьте вес пользователя на среднюю скорость и продолжительность в минутах.
	// Разделите результат на число минут в часе для получения количества потраченных калорий.
	// Умножить полученное число калорий на корректирующий коэффициент walkingCaloriesCoefficient. Соответствующая константа объявлена в пакете. Вернуть полученное значение.
	spentCalories := ((weight * speed * durationMinutes) / minInH) * walkingCaloriesCoefficient
	return spentCalories, nil
}
