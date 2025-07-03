package main

import (
	"fmt"
)

func printQuery(query [5]string) {
	for i := 0; i < len(query); i++ {
		if query[i] == "" {
			fmt.Println(
				fmt.Sprintf("%d. %s", i+1, "-"),
			)
		} else {
			fmt.Println(
				fmt.Sprintf("%d. %s", i+1, query[i]),
			)
		}
	}
}

func printAmount(query [5]string) {
	var occup int
	for i := 0; i < len(query); i++ {
		if query[i] != "" {
			occup++
		}
	}
	fmt.Println(
		fmt.Sprintf("Осталось свободных мест: %d", len(query)-occup),
	)
	fmt.Println(
		fmt.Sprintf("Всего человек в очереди: %d", occup),
	)
}

func isFull(query [5]string) (flag bool) {
	flag = true
	for i := 0; i < len(query); i++ {
		if query[i] == "" {
			flag = false
			return
		}
	}
	return flag
}

func isOccup(query [5]string, number int) (flag bool) {
	flag = true
	if query[number] == "" {
		flag = false
		return
	}
	return
}

func main() {
	var query [5]string

	for {
		var name string
		var number int
		fmt.Scan(&name)

		if name == "очередь" {
			printQuery(query)
		} else if name == "количество" {
			printAmount(query)
		} else if name == "конец" {
			printQuery(query)
			break
		} else {
			fmt.Scan(&number)
			if number <= 0 || number > 5 {
				fmt.Println(
					fmt.Sprintf("Запись на место номер %d невозможна: некорректный ввод", number),
				)
				continue
			} else if isFull(query) {
				fmt.Println(
					fmt.Sprintf("Запись на место номер %d невозможна: очередь переполнена", number),
				)
				continue
			} else if isOccup(query, number-1) {
				fmt.Println(
					fmt.Sprintf("Запись на место номер %d невозможна: место уже занято", number),
				)
				continue
			}
			query[int(number)-1] = name
		}
	}
}
