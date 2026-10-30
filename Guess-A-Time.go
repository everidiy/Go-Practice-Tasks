package main

import (
	"errors"
	"strings"
	"time"
	"unicode/utf8"
)

var errorDayOrNight = errors.New("исправь свой ответ, а лучше ложись поспать")

func currentDayOfTheWeek() string {
	currentTime := TimeNow()
	return getRussianWeekday(currentTime.Weekday())
}

func dayOrNight() string {
	currentTime := TimeNow()
	hour := currentTime.Hour()

	if hour >= 10 && hour < 22 {
		return "День"
	}

	return "Ночь"
}

func nextFriday() int {
	currentTime := TimeNow()
	today := currentTime.Weekday()

	if today == time.Friday {
		return 0
	}

	nextFriday := (time.Friday - today + 7) % 7
	return int(nextFriday)

}

func CheckCurrentDayOfTheWeek(answer string) bool {
	currentDay := currentDayOfTheWeek()
	return strings.EqualFold(strings.TrimSpace(answer), currentDay)
}

func getRussianWeekday(weekday time.Weekday) string {
	switch weekday {
	case time.Sunday:
		return "Воскресенье"
	case time.Monday:
		return "Понедельник"
	case time.Tuesday:
		return "Вторник"
	case time.Wednesday:
		return "Среда"
	case time.Thursday:
		return "Четверг"
	case time.Friday:
		return "Пятница"
	case time.Saturday:
		return "Суббота"
	default:
		return "Неизвестный день"
	}
}

func CheckNowDayOrNight(answer string) (bool, error) {
	if utf8.RuneCountInString(answer) != 4 {
		return false, errorDayOrNight
	}

	currentTime := TimeNow()
	hour := currentTime.Hour()
	isDay := hour >= 10 && hour < 22

	lowerAnswer := strings.ToLower(strings.TrimSpace(answer))

	if (isDay && lowerAnswer == "день") || (!isDay && lowerAnswer == "ночь") {
		return true, nil
	}

	return false, nil
}

func TimeNow() time.Time {
	return time.Now()
}
