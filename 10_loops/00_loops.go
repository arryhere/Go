package loops

import "fmt"

func myLoop() {
	for i := 0; i <= 10; i++ {
		fmt.Println("i:", i)
	}
}

func myForWhileLoop() {
	var data string = "start"
	var i int = 0

	for data == "start" {
		if i == 10 {
			data = "stop"
		}

		fmt.Printf("i: [%v], data: [%v]\n", i, data)
		i++
	}
}

func Loops() {
	myLoop()
	/*
		i: 0
		i: 1
		i: 2
		i: 3
		i: 4
		i: 5
		i: 6
		i: 7
		i: 8
		i: 9
		i: 10
	*/

	myForWhileLoop()
	/*
		i: [0], data: [start]
		i: [1], data: [start]
		i: [2], data: [start]
		i: [3], data: [start]
		i: [4], data: [start]
		i: [5], data: [start]
		i: [6], data: [start]
		i: [7], data: [start]
		i: [8], data: [start]
		i: [9], data: [start]
		i: [10], data: [stop]
	*/
}
