package loops

import "fmt"

func breakLoop() {
	for i := 10; i > 0; i-- {
		if i == 5 {
			fmt.Println("i:", i, "- breaking loop")
			break
		}
		fmt.Println("i:", i)
	}
}

func Break() {
	breakLoop()
	/*
		i: 10
		i: 9
		i: 8
		i: 7
		i: 6
		i: 5 - breaking loop
	*/
}
