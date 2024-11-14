package interfaces

import "fmt"

// Interface
type Book interface {
	GetBookName() string
}

// Go Book
type GoBook struct {
	name        string
	goExtension string
}

func (gb GoBook) GetBookName() string {
	return gb.name
}

// Rsut Book
type RustBook struct {
	name          string
	rustExtension string
}

func (rb RustBook) GetBookName() string {
	return rb.name
}

// Functions
func GetBookInfo(b Book) []string {
	gb, gb_ok := b.(GoBook)
	rb, rb_ok := b.(RustBook)

	if gb_ok {
		return []string{gb.name, gb.goExtension}
	}

	if rb_ok {
		return []string{rb.name, rb.rustExtension}
	}

	return nil
}

func Interface_Casting() {
	myGoBook := GoBook{name: "LetUsGo", goExtension: ".go"}
	myRustBook := GoBook{name: "TheRustLang", goExtension: ".rs"}

	fmt.Println("Go Book Name:", GetBookInfo(myGoBook))     // Go Book Name: [LetUsGo .go]
	fmt.Println("Rust Book Name:", GetBookInfo(myRustBook)) // Rust Book Name: [TheRustLang .rs]
}
