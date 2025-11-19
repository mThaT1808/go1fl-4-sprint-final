package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	
	// Разделение приходящей строки на количество шагов и длительность
	parts := strings.Split(data, ",")
	// Проверить, чтобы длина слайса была равна 2, так как в строке данных у нас количество шагов и продолжительность.
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("ошибка при разделении строки. Ожидается: 'шаги,время'")
	}

	// Преобразование строки в число шагов
	// Преобразовать первый элемент слайса (количество шагов) в тип int.
	steps, err := strconv.Atoi(parts[0])
	// Обработать возможные ошибки. При их возникновении из функции вернуть 0 шагов, 0 продолжительность и ошибку.
	if err != nil {
		return 0, 0, fmt.Errorf("ошибка при определении шагов: %v", err)
	}
	// Проверить: количество шагов должно быть больше 0. Если это не так, вернуть нули и ошибку.
	if steps <= 0 {
		return 0, 0, fmt.Errorf("ошибка при определении шагов: Шагов должно быть больше нуля")
	}
	// Преобразование строки в time.Duration
	duration, err := time.ParseDuration(parts[1])
	// Обработать возможные ошибки. При их возникновении из функции вернуть 0 шагов, 0 продолжительность и ошибку.
	if err != nil {
		return 0, 0, fmt.Errorf("ошибка определении длительности: %v", err)
	}
	// Если всё прошло без ошибок, верните количество шагов, продолжительность и nil (для ошибки).
	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// Получить данные о количестве шагов и продолжительности прогулки с помощью функции parsePackage(). 
	steps, duration, err := parsePackage(data)
	// В случае возникновения ошибки вывести её на экран и вернуть пустую строку.
	if err != nil {
		fmt.Println(err)
		return ""
	}
	// Проверить, чтобы количество шагов было больше 0. В противном случае вернуть пустую строку.
	if steps <= 0 {
		return ""
	}
	// Вычислить дистанцию в метрах. Дистанция равна произведению количества шагов на длину шага. Константа stepLength (длина шага) уже определена в коде.
	distance := float64(steps) * stepLength
	// Перевести дистанцию в километры, разделив её на число метров в километре (константа mInKm, определена в пакете).
	distanceKm := float64(distance / mInKm)
	// Вычислить количество калорий, потраченных на прогулке. Функция для вычисления калорий WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error).
	calories, err := spentcalories.WalkingSpentCalories(steps, 0, 0, duration)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	message := fmt.Sprintf("Количество шагов: %d \nДистанция составила %f км. \nВы сожгли %f ккал.", steps, distanceKm, calories)
	return message
}
