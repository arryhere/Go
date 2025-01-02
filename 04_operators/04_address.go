package operators

import (
	"fmt"
)

func Address_Operators() {
	var v1 int = 5

	var v2 *int = &v1   // pointer
	var v3 **int = &v2  // pointer to a pointer
	var v4 ***int = &v3 // pointer to a pointer to a pointer

	v5 := ***v4 // Dereferencing

	fmt.Println("v1:", v1) // v1: 5

	fmt.Println("v2:", v2) // v2: 0x14000104020
	fmt.Println("v3:", v3) // v3: 0x140000a4028
	fmt.Println("v4:", v4) // v4: 0x14000110030

	fmt.Println("v5:", v5) // v5: 5

}
