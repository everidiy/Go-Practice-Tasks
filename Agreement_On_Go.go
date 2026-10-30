package main

import (
	"fmt"
	"time"
)

func main() {
	var name, surname, dad string
	var a, b, c float64
	date := "25.04.2005"

	fmt.Scan(&name)
	fmt.Scan(&surname)
	fmt.Scan(&dad)
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)

	parsed, _ := time.Parse("02.01.2006", date)

	newDate := time.Date(parsed.Year(), parsed.Month(), parsed.Day()+15, parsed.Hour(), parsed.Minute(), parsed.Second(), parsed.Nanosecond(), parsed.Location())

	formatted := newDate.Format("02.01.2006")

	salary := a + b + c
	rubles := int(salary)
	kopeks := int((salary - float64(rubles)) * 100)

	fmt.Printf(`Уважаемый, %s %s %s, доводим до вашего сведения, что бухгалтерия сформировала документы по факту выполненной вами работы.
Дата подписания договора: %s. Просим вас подойти в офис в любое удобное для вас время в этот день.
Общая сумма выплат составит %d руб. %d коп.

С уважением,
Гл. бух. Иванов А.Е.`, surname, name, dad, formatted, rubles, kopeks)
}
