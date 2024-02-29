// challenges/types-composite/begin/main.go
package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

// define an author type with a name
type author struct {
	name string
}

// define a book type with a title and author
type book struct {
	title string
	author author
}

// define a library type to track a list of books
type library map[string][]book

// define addBook to add a book to the library
func (l library) addBook (b book) {
	l[b.author.name] = append(l[b.author.name], b)
}

// define a lookupByAuthor function to find books by an author's name
func (l library) lookupByAuthor (authorName string) []book {
	a := l[authorName]
	return a
} 

func main() {
	// create a new library
	var lib = make(library)
	// lib := library{}

	// add 2 books to the library by the same author
	book1 := book {
		title: "The Ruins of Gorlan",
		author : author {
			name: "John Flanagan",
		},
	}

	book2 := book {
		title: "The Burning Bridge",
		author : author {
			name: "John Flanagan",
		},
	}

	lib.addBook(book1)
	lib.addBook(book2)

	// add 1 book to the library by a different author
	differentAuthor := book {
		title: "The Lightning Thief",
		author : author {
			name: "Rick Riordan",
		},
	}

	lib.addBook(differentAuthor)

	another := book {
		title: "Slathbog's Gold",
		author : author {
			name: "M.L. Forman",
		},
	}

	lib.addBook(another)

	// dump the library with spew
	fmt.Println("Using Spew")
	fmt.Println("=================")
	spew.Dump(lib)
	fmt.Println("=================")
	fmt.Println("=================")
	fmt.Println()
	// lookup books by known author in the library
	fmt.Println("LookUpByAuthor:")
	fmt.Println("=================")
	fmt.Println(lib.lookupByAuthor("John Flanagan"))
	fmt.Println("=================")
	fmt.Println(lib.lookupByAuthor("Rick Riordan"))
	fmt.Println("=================")
	fmt.Println(lib.lookupByAuthor("M.L. Forman"))
	fmt.Println("=================")
	fmt.Println()

	// print out the first book's title and its author's name
	fmt.Println("Looking up the first book")
	fmt.Println("=================")
	firstBook := lib.lookupByAuthor("John Flanagan")[0]
	fmt.Printf("Book Title: %s, Author Name: %s\n", firstBook.title, firstBook.author.name)

}
