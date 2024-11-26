package maps

import "fmt"

func Maps() {
	map1 := map[string]int{}
	map2 := map[string]int{
		"key_1": 1,
		"key_2": 2,
		"key_3": 3,
	}

	fmt.Println("map1:", map1) // map1: map[]
	fmt.Println("map2:", map2) // map2: map[key_1:1 key_2:2 key_3:3]

	map2["key_1"] = 5 // key exist, it is being updated
	map2["key_4"] = 4 // key does not exist, it is being created

	v := map2["key_1"]

	fmt.Println("map2:", map2) // map2: map[key_1:5 key_2:2 key_3:3 key_4:4]
	fmt.Println("v:", v)       // v: 5

	map2length := len(map2)
	fmt.Println("map2length:", map2length) // map2length: 4

	/* make */
	map3 := make(map[string]int)

	fmt.Println("map3:", map3) // map3: map[]

	map3["key_1"] = 5
	map3["key_2"] = 4
	map3["key_3"] = 3

	fmt.Println("map3:", map3) // map3: map[key_1:5 key_2:4 key_3:3]

	/* delete */
	delete(map3, "key_1")

	fmt.Println("map3:", map3) // map3: map[key_2:4 key_3:3]

	/* check if key exist */
	e1, e1_ok := map3["key_1"]
	e2, e2_ok := map3["key_2"]

	fmt.Printf("e1: [%v], e1_ok: [%v]\n", e1, e1_ok) // e1: [0], e1_ok: [false]
	fmt.Printf("e2: [%v], e2_ok: [%v]\n", e2, e2_ok) // e2: [4], e2_ok: [true]
}
