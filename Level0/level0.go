package main

import (
	"fmt"
	"math"
	"strings"
	"time"
)

func main() {
	var dateString, firstName, lastName, fatherName string

	var payout1, payout2, payout3 float64

	fmt.Scanln(&dateString)

	fmt.Scanln(&firstName)
	fmt.Scanln(&lastName)
	fmt.Scanln(&fatherName)

	fmt.Scanln(&payout1)
	fmt.Scanln(&payout2)
	fmt.Scanln(&payout3)

	dateSplit := strings.Split(dateString, ".")
	dateSplit[0], dateSplit[2] = dateSplit[2], dateSplit[0]
	dateString = strings.Join(dateSplit, "-")

	const DateTime = "2006-01-02"
	var date time.Time

	date, _ = time.Parse(DateTime, dateString)
	date = date.AddDate(0, 0, 15)

	sum := payout1 + payout2 + payout3
	puble := math.Trunc(sum)
	kopek := math.Round((sum - math.Trunc(sum)) * 100)

	fmt.Printf(`Уважаемый, %s %s %s, доводим до вашего сведения, что бухгалтерия сформировала документы по факту выполненной вами работы.
Дата подписания договора: %d.%02d.%d. Просим вас подойти в офис в любое удобное для вас время в этот день.
Общая сумма выплат составит %.0f руб. %0.f коп.

С уважением,
Гл. бух. Иванов А.Е. `,
		lastName,
		firstName,
		fatherName,
		date.Day(),
		date.Month(),
		date.Year(),
		puble,
		kopek)
}
