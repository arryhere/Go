package loops

import "fmt"

func continueLoop() {
	for i := 10; i > 0; i-- {
		if i == 5 {
			fmt.Println("skipping")
			continue
		}
		fmt.Println("i:", i)
	}
}

func Continue() {
	continueLoop()
	/*
		i: 10
		i: 9
		i: 8
		i: 7
		i: 6
		skipping
		i: 4
		i: 3
		i: 2
		i: 1
	*/
}
