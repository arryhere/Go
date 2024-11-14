package interfaces

import "fmt"

// Interface
type Book interface {
	getBookName() string
}

// Go Book
type GoBook struct {
	name        string
	goExtension string
}

func (gb GoBook) getBookName() string {
	return gb.name
}

// Rsut Book
type RustBook struct {
	name          string
	rustExtension string
}

func (rb RustBook) getBookName() string {
	return rb.name
}

// Functions
func getBookInfo(b Book) []string {

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

	fmt.Println("Go Book Name:", getBookInfo(myGoBook))
	fmt.Println("Rust Book Name:", getBookInfo(myRustBook))
}
