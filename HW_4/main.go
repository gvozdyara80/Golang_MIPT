package main

import "fmt"

type Book struct {
	Title  string
	Author string
	Year   int
}

type Library struct {
	Books []Book
}

func (l *Library) AddBook(book Book) {
	l.Books = append(l.Books, book)
}

func (l *Library) RemoveBook(title string) {
	for i, book := range l.Books {
		if book.Title == title {
			l.Books = append(l.Books[:i], l.Books[i+1:]...)
			break
		}
	}
}

func (l *Library) DisplayBooks() {
	for _, book := range l.Books {
		fmt.Printf("Title: %s, Author: %s, Year: %d\n", book.Title, book.Author, book.Year)
	}
}

type Readable interface {
	Read()
}

func (b Book) Read() {
	fmt.Printf("Reading the book: %s by %s\n", b.Title, b.Author)
}

func ReadBook(r Readable) {
	r.Read()
}

func main() {
	library := Library{}

	library.AddBook(Book{Title: "The Go Programming Language", Author: "Alan A. A. Donovan and Brian W. Kernighan", Year: 2015})
	library.AddBook(Book{Title: "Clean Code", Author: "Robert C. Martin", Year: 2008})

	fmt.Println("Library Books:")
	library.DisplayBooks()

	ReadBook(library.Books[0])
}
