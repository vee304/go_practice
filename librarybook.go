// Library Management
// practicing State Management using a simple boolean (true/false) flag.
//

package main

import "fmt"

type Book struct {
	Title      string
	Author     string
	Ischeckout bool
}

func (b *Book) Borrow() {
	if b.Ischeckout {
		fmt.Println("book is already checked out")

	} else {
		b.Ischeckout = true
		fmt.Println("BOOK is successfully taken")
	}

}

func (b *Book) ReturnBook() {
	b.Ischeckout = false
	fmt.Println("Book has been given back")

}

func main() {
	mybook := Book{
		Title:  "Knight rider",
		Author: "Nolan",
	}

	fmt.Println("---Welcome to the Library---")

	mybook.Borrow()
	mybook.ReturnBook()

}
