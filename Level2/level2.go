package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
	"unicode/utf8"
)

var (
	ErrAnswer = errors.New("исправь свой ответ, а лучше ложись поспать")
)

func currentDayOfTheWeek() string {
	start := TimeNow()
	dayOfWeek := int(start.Weekday())
	switch dayOfWeek {
	case 1:
		return "Понедельник"
	case 2:
		return "Вторник"
	case 3:
		return "Среда"
	case 4:
		return "Четверг"
	case 5:
		return "Пятница"
	case 6:
		return "Суббота"
	case 7:
		return "Воскресенье"
	default:
		return "Ошибка"
	}
}

func dayOrNight() string {
	start := TimeNow()
	hourOfDay := int(start.Hour())
	if hourOfDay >= 10 && hourOfDay <= 22 {
		return "День"
	} else {
		return "Ночь"
	}
}

func nextFriday() int {
	start := TimeNow()
	dayOfWeek := int(start.Weekday())
	if dayOfWeek <= 5 {
		return 5 - dayOfWeek
	} else {
		return dayOfWeek - 5
	}
}

func CheckCurrentDayOfTheWeek(answer string) bool {
	return strings.EqualFold(answer, currentDayOfTheWeek())
}

func CheckNowDayOrNight(answer string) (bool, error) {
	if utf8.RuneCountInString(answer) == 4 {
		return strings.EqualFold(answer, dayOrNight()), nil
	} else {
		return false, ErrAnswer
	}
}

func TimeNow() time.Time {
	const DateTime = "2006-01-02 15:04:05"
	t, _ := time.Parse(DateTime, "2006-01-06 09:04:05")
	return t
}

func main() {
	fmt.Println(CheckCurrentDayOfTheWeek("вторник"))
	fmt.Println(CheckCurrentDayOfTheWeek("Вторник"))
	fmt.Println(CheckCurrentDayOfTheWeek("ВтОрНиК"))
	fmt.Println(CheckCurrentDayOfTheWeek("среда"))
	fmt.Println(CheckCurrentDayOfTheWeek("четверг"))
	fmt.Println(CheckCurrentDayOfTheWeek("пятница"))
	fmt.Println(CheckCurrentDayOfTheWeek("суббота"))
	fmt.Println(CheckCurrentDayOfTheWeek("воскресенье"))
	fmt.Println(CheckCurrentDayOfTheWeek("понедельник"))
	fmt.Println(CheckNowDayOrNight("день"))
	fmt.Println(CheckNowDayOrNight("День"))
	fmt.Println(CheckNowDayOrNight("ДеНь"))
	fmt.Println(CheckNowDayOrNight("ночь"))
	fmt.Println(CheckNowDayOrNight("ок"))
	fmt.Println(CheckNowDayOrNight("гидроэлектростанция"))
}

/*
utf8.RuneCountInString()
EqualFold
//return strings.ToLower(answer) == strings.ToLower(currentDayOfTheWeek())
*/
