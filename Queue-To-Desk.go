package main

import "fmt"

func main() {
	queue := [5]string{"", "", "", "", ""}

	for {
		var input string
		fmt.Scan(&input)

		if input == "" {
			continue
		}

		if input == "queue" {
			for i := 0; i < 5; i++ {
				if queue[i] == "" {
					fmt.Printf("%d. -\n", i+1)
				} else {
					fmt.Printf("%d. %s\n", i+1, queue[i])
				}
			}
			continue
		}

		if input == "count" {
			free := 0
			occupied := 0
			for i := 0; i < 5; i++ {
				if queue[i] == "" {
					free++
				} else {
					occupied++
				}
			}
			fmt.Printf("Free spots left: %d\n", free)
			fmt.Printf("Total people in queue: %d\n", occupied)
			continue
		}

		if input == "end" {
			for i := 0; i < 5; i++ {
				if queue[i] == "" {
					fmt.Printf("%d. -\n", i+1)
				} else {
					fmt.Printf("%d. %s\n", i+1, queue[i])
				}
			}
			break
		}

		var num int
		_, err := fmt.Scan(&num)
		if err != nil {
			fmt.Printf("Recording to spot number %d is impossible: ", num)
			fmt.Printf("invalid input\n")
			continue
		}

		if num < 1 || num > 5 {
			fmt.Printf("Recording to spot number %d is impossible: ", num)
			fmt.Printf("invalid input\n")
			continue
		}

		freeSlots := 0
		for i := 0; i < 5; i++ {
			if queue[i] == "" {
				freeSlots++
			}
		}
		if freeSlots == 0 {
			fmt.Printf("Recording to spot number %d is impossible: ", num)
			fmt.Printf("queue is full\n")
			continue
		}

		if queue[num-1] != "" {
			fmt.Printf("Recording to spot number %d is impossible: ", num)
			fmt.Printf("spot is already taken\n")
			continue
		}

		queue[num-1] = input

	}
}
